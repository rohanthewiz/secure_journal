package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// Pretend DB
var Key []byte
var HashedPassword string

// key + password = hash

func main() {
	for {
		fmt.Print(MenuOptions())

		// Take in the user's choice
		var choice int
		_, err := fmt.Scanln(&choice) // read from keyboard
		if err != nil {
			fmt.Println("Error on input", err)
			return
		}

		switch choice {
		case 1:
			// Register
			fmt.Println("Register")
			fmt.Println("Enter a password")
			var password string
			_, _ = fmt.Scanln(&password) // read from keyboard
			_ = Register(password)

		case 2:
			// Login
			var password string
			fmt.Println("Enter your password")
			_, _ = fmt.Scanln(&password) // read from keyboard
			fmt.Println("Logging in...")
			Login(password)
		case 3:
			// Delete me
			fmt.Println("Delete me")
		case 4:
			// Exit
			return
		}
	}
	/*	// This would be stored in database of some sort
		// password + key (salt) = hash
		passHash1 := "314043fe3a87549076364c6c96d2dd793dc21fa5c767c4c0204e9500dda93d94"
		randomKey := "some-random-key" // I would get this from the database
		// This is the password we "registered" with  -> "$ecure$p@ssword"

		// Pretend that this is the password the user is trying to login with
		bytPassword := []byte("$ecure$p@ssword")

		key := []byte(randomKey)

		hash := hashPassword(key, bytPassword)

		// We were using %s here for string, but %q will also put double quotes around the value.
		// If we had done this earlier we would have seen that one of the values also included a
		// newline character (\n).
		fmt.Printf("Password: %q, Hash: %q\n", bytPassword, hash)

		fmt.Printf("passHash1: %s\n", passHash1)

		if passHash1 == hash {
			fmt.Println("Password is correct")

			// Show me my private journals!
			// - Decrypt my private journal
		}
	*/
}

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

func MenuOptions() (menu string) {
	menu = fmt.Sprintln("===== Menu =====")
	menu += fmt.Sprintln("1. Register")
	menu += fmt.Sprintln("2. Login")
	menu += fmt.Sprintln("3. Delete me")
	menu += fmt.Sprintln("4. Exit")
	menu += fmt.Sprint("Enter your choice: ")
	return
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

// SCRATCH

// Two kinds of hashes

// Login to the app

// Encrypt the journal entry

// Display when needed

// Generate a random salt
/*salt := make([]byte, 32)
if _, err := rand.Read(salt); err != nil {
fmt.Printf("Error generating salt: %v", err)
return ""
}
*/

// "some-text" --> [ hash function] -> gibberish
// - One way hash - (you can't reverse it -- not by any known means)

/*	str := "this is a secret key; you should generate a strong random key that's at least 32 bytes long"
	fmt.Println(str)

	k := []byte(str) // cast to []byte
	fmt.Println(k)

	fmt.Printf("%x\n", k)

	fmt.Printf("%T\n", k)
*/
