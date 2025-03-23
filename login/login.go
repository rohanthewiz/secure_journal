package login

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"

	_ "github.com/marcboeker/go-duckdb"
	"golang.org/x/crypto/bcrypt"
)

// InitDB initializes the database connection
func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("duckdb", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Create users table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			username VARCHAR PRIMARY KEY,
			key VARCHAR,
			password VARCHAR
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	return db, nil
}

// Register a new user
func Register(db *sql.DB, username string, password string) error {
	// Check if username exists
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", username).Scan(&count)
	if err != nil {
		return fmt.Errorf("database error checking username: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("username already exists")
	}

	// Generate a random key
	key := make([]byte, 32)
	_, err = rand.Read(key)
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}

	// Hash password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	keyHex := hex.EncodeToString(key)

	// Insert user into database
	_, err = db.Exec(
		"INSERT INTO users (username, key, password) VALUES ($1, $2, $3)",
		username, keyHex, hashedPassword,
	)
	if err != nil {
		return fmt.Errorf("failed to register user: %w", err)
	}

	return nil
}

// Login a user
func Login(db *sql.DB, username string, password string) error {
	var storedHash string

	// Retrieve password hash from database
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("invalid username or password")
		}
		return fmt.Errorf("database error during login: %w", err)
	}

	// Check password
	if !CheckPasswordHash(password, storedHash) {
		return fmt.Errorf("invalid username or password")
	}

	fmt.Println("Login successful!")
	return nil
}

// Delete a user
func Delete(db *sql.DB, username string, password string) error {
	// Verify user credentials
	var storedHash string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("database error: %w", err)
	}

	if !CheckPasswordHash(password, storedHash) {
		return fmt.Errorf("invalid password")
	}

	// Delete user
	_, err = db.Exec("DELETE FROM users WHERE username = $1", username)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// Hash password using bcrypt
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// Compare password with stored hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
