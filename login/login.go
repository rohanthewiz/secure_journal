package login

import (
	"crypto/rand"
	"fmt"
	"os"

	"golang.org/x/crypto/sha3"
)

// Pretend DB
var Key []byte
var HashedPassword string

type User struct {
	username string
	password string
}

func Register(username string, password string) (err error) {
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return err
	}
	HashedPassword = hashPassword(key, []byte(password))
	Key = key

	f, err := os.Create("userdata.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = fmt.Fprintf(f, "username: %s\nkey: %x\nhashed_password: %s", username, Key, HashedPassword)
	if err != nil {
		return err
	}
	return
}

func Login(username string, password string) (err error) {
	fmt.Printf("You entered %q\n", password) // Super IMPORTANT!
	// Compare the hash of the password with the hash of the stored password
	// If they match, the password is correct
	// If they don't match, the password is incorrect
	// user := User{username: username, password: password}
	f, err := os.Open("userdata.txt")
	if err != nil {
		return err
	}
	fileInfo, err := f.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = f.Read(buffer)
	if err != nil {
		return err
	}

	// Convert the file content to a string
	fileContent := string(buffer)

	// Check if the username and password match the content in the file
	expectedContent := fmt.Sprintf("username: %s, password: %s", username, password)
	if fileContent == expectedContent {
		currentHash := hashPassword(Key, []byte(password))
		if currentHash == HashedPassword {
			fmt.Println("You are logged in")
			fmt.Println("Show me my private journals!")
		} else {
			fmt.Println("Password is incorrect")
		}
	} else {
		fmt.Println("Username or password is incorrect")
	}

	return nil
}

func hashPassword(key []byte, dataToEncrypt []byte) (strHash string) {
	// A MAC with 32 bytes of output has 256-bit security strength -- if you use at least a 32-byte-long key.
	// A byte is 8 bits
	hashOutput := make([]byte, 32)
	hasher := sha3.NewShake256()

	// Write the key into the hash.
	_, err := hasher.Write(key)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Now write the data.
	_, err = hasher.Write(dataToEncrypt)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read 32 bytes of output from the hash into h.
	_, err = hasher.Read(hashOutput)
	if err != nil {
		fmt.Println(err)
		return
	}

	// strHash = fmt.Sprintf("%x", hashOutput)
	strHash = string(hashOutput)
	fmt.Printf("%s", strHash)

	return
}
