package login

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/crypto/sha3"
)

// Pretend DB
var Key []byte
var HashedPassword string

type UserData struct {
	Username string
	Key      string
	Password string
}

func checkIfExists(username string) (bool, error) {
	// Read the existing users from the file
	data, err := os.ReadFile("userdata.json")
	if err != nil {
		// If the file doesn't exist, assume no users exist yet
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	// Unmarshal the JSON data into a slice of UserData
	var users []UserData
	err = json.Unmarshal(data, &users)
	if err != nil {
		return false, err
	}

	// Check if the username already exists
	for _, user := range users {
		if user.Username == username {
			return true, nil // Username exists
		}
	}

	return false, nil // Username does not exist
}
func Register(username string, password string) (err error) {
	// Check if the username already exists
	exists, err := checkIfExists(username)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("username already exists")
	}

	// Generate a new key for the user
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return err
	}

	// Hash the password
	hashedPassword := hashPassword(key, []byte(password))

	// Read existing users
	var users []UserData
	data, err := os.ReadFile("userdata.json")
	if err == nil {
		json.Unmarshal(data, &users)
	}

	// Add new user
	newUser := UserData{
		Username: username,
		Key:      hex.EncodeToString(key),
		Password: hashedPassword,
	}
	users = append(users, newUser)

	// Write back to file
	jsonData, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile("userdata.json", jsonData, 0644)
}
func Login(username string, password string) (err error) {
	data, err := os.ReadFile("userdata.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	var users []UserData
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return err
	}

	for _, user := range users {
		if user.Username == username {
			keyBytes, err := hex.DecodeString(user.Key)
			if err != nil {
				fmt.Println("Error decoding key:", err)
				return err
			}
			currentHash := hashPassword(keyBytes, []byte(password))
			if currentHash == user.Password {
				fmt.Println("Login successful!")
				return nil
			}
		}
	}
	return fmt.Errorf("invalid username or password")
}

func hashPassword(key []byte, dataToEncrypt []byte) string {
	hashOutput := make([]byte, 32)
	hasher := sha3.NewShake256()

	// Write the key and data into the hash.
	hasher.Write(key)
	hasher.Write(dataToEncrypt)

	// Read 32 bytes of output from the hash.
	hasher.Read(hashOutput)

	// Return the hexadecimal representation of the hash.
	return hex.EncodeToString(hashOutput)
}
