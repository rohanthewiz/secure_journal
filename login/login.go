package login

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// Pretend DB
var Key []byte
var HashedPassword string

func Register(password string) (err error) {
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return err
	}
	HashedPassword = hashPassword(key, []byte(password))
	Key = key
	return nil
}

func Login(password string) (err error) {
	fmt.Printf("You entered %q\n", password) // Super IMPORTANT!
	// Compare the hash of the password with the hash of the stored password
	// If they match, the password is correct
	// If they don't match, the password is incorrect
	currentHash := hashPassword(Key, []byte(password))
	if currentHash == HashedPassword {
		fmt.Println("You are logged in")
		fmt.Println("Show me my private journals!")
	} else {
		fmt.Println("Password is incorrect")
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

	strHash = fmt.Sprintf("%x", hashOutput)

	return
}
