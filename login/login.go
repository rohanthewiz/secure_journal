package login

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/marcboeker/go-duckdb"
	"golang.org/x/crypto/sha3"
)

// DB is the database connection
var DB *sql.DB

// UserData structure (for compatibility with existing code)
type UserData struct {
	Username string
	Key      string
	Password string
}

// InitDB initializes the database connection and creates the users table if it doesn't exist
func InitDB(dbPath string) error {
	var err error
	// Open the database connection
	DB, err = sql.Open("duckdb", dbPath)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	// Create users table if it doesn't exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			username VARCHAR PRIMARY KEY,
			key VARCHAR,
			password VARCHAR
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	return nil
}

// CloseDB closes the database connection
func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// checkIfExists checks if a username already exists in the database
func checkIfExists(username string) (bool, error) {
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("database error checking username: %w", err)
	}
	return count > 0, nil
}

// Register adds a new user to the database
func Register(username string, password string) error {
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
		return fmt.Errorf("failed to generate key: %w", err)
	}

	// Hash the password
	hashedPassword := hashPassword(key, []byte(password))
	keyHex := hex.EncodeToString(key)

	// Insert the new user into the database
	_, err = DB.Exec(
		"INSERT INTO users (username, key, password) VALUES (?, ?, ?)",
		username, keyHex, hashedPassword,
	)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}

// Login verifies a user's credentials
func Login(username string, password string) error {
	var keyHex, storedHash string

	// Query the database for the user
	err := DB.QueryRow(
		"SELECT key, password FROM users WHERE username = ?",
		username,
	).Scan(&keyHex, &storedHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid username or password")
		}
		return fmt.Errorf("database error during login: %w", err)
	}

	// Decode the key from hex
	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil {
		return fmt.Errorf("error decoding key: %w", err)
	}

	// Hash the provided password with the stored key
	currentHash := hashPassword(keyBytes, []byte(password))

	// Compare the hashes
	if currentHash == storedHash {
		fmt.Println("Login successful!")
		return nil
	}

	return fmt.Errorf("invalid username or password")
}

// Delete removes a user from the database after verifying credentials
func Delete(username string, password string) error {
	// First verify the credentials
	err := Login(username, password)
	if err != nil {
		return err
	}

	// If credentials are valid, delete the user
	_, err = DB.Exec("DELETE FROM users WHERE username = ?", username)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// hashPassword hashes a password with a key using SHAKE-256
func hashPassword(key []byte, dataToEncrypt []byte) string {
	hashOutput := make([]byte, 32)
	hasher := sha3.NewShake256()
	// Write the key and data into the hash.
	hasher.Write(key)
	hasher.Write(dataToEncrypt)
	// Read 32 bytes of output from the hash.
	// nolint:errorcheck
	hasher.Read(hashOutput)
	// Return the hexadecimal representation of the hash.
	return hex.EncodeToString(hashOutput)
}
