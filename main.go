package main

import (
	"log"
	"secure_journal/web"
)

func main() {

	s := web.InitWeb()
	log.Println(s.Run())

	/*	// Middleware
		s.Use(func(ctx rweb.Context) error {
			start := time.Now()

			defer func() {
				fmt.Println(ctx.Request().Method(), ctx.Request().Path(), time.Since(start))
			}()
			return ctx.Next()
		})


	*/
	/*	for {
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
	*/
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

/*  // SCRATCH
//
// func MenuOptions() (menu string) {
// 	menu = fmt.Sprintln("===== Menu =====")
// 	menu += fmt.Sprintln("1. Register")
// 	menu += fmt.Sprintln("2. Login")
// 	menu += fmt.Sprintln("3. Delete me")
// 	menu += fmt.Sprintln("4. Exit")
// 	menu += fmt.Sprint("Enter your choice: ")
// 	return
// }

// Two kinds of hashes

// Login to the app

// Encrypt the journal entry

// Display when needed
*/

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
