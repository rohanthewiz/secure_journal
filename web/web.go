package web

import (
	"fmt"
	"github.com/rohanthewiz/rweb"
	"secure_journal/login"
)

func InitWeb() (s *rweb.Server) {
	s = rweb.NewServer(
		rweb.ServerOptions{
			Address: "localhost:8000",
			Verbose: true,
		},
	)

	head := "<head><title>My Journal</title>"
	head += "<style>body {background-color: lightblue;}</style></head>"
	pageStart := "<html>"
	pageEnd := "</html>"

	rootHandler := func(ctx rweb.Context) error {
		body := "<body><h1>My Journal</h1>" + WebOptions() + "</body>"
		page := pageStart + head + body + pageEnd
		fmt.Println(page)
		return ctx.WriteHTML(page)
	}

	s.Get("/", rootHandler)

	s.Get("/register", func(ctx rweb.Context) (err error) {
		body := "<body><h1>My Journal</h1>" + WebOptions() +
			`<p style="color: navy">Register</p>` +
			`<form action="/register" method="POST">
                <label for="username">Username:</label><br>
                <input type="text" name="username" id="username"><br>
                <label for="password">Password:</label><br>
                <input type="password" name="password" id="password"><br>
                <input type="submit" value="Register">
            </form>` +
			"</body>"
		page := pageStart + head + body + pageEnd
		return ctx.WriteHTML(page)
	})

	s.Post("/register", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")

		err = login.Register(password)
		if err != nil {
			// Return an error page instead of just the error
			errorBody := "<body><h1>My Journal</h1>" + WebOptions() +
				`<p style="color: red">Registration failed: ` + err.Error() + `</p>` +
				`<a href="/register">Try again</a>` +
				"</body>"
			page := pageStart + head + errorBody + pageEnd
			return ctx.WriteHTML(page)
		}

		err = login.Login(password)
		if err != nil {
			// Handle login error
			errorBody := "<body><h1>My Journal</h1>" + WebOptions() +
				`<p style="color: red">Login failed after registration: ` + err.Error() + `</p>` +
				`<a href="/">Return to home</a>` +
				"</body>"
			page := pageStart + head + errorBody + pageEnd
			return ctx.WriteHTML(page)
		}
		successMsg := `<div style="margin: 20px;"><p style="color: green">Registration successful!</p></div>`
		body := "<body><h1>My Journal</h1>" + successMsg + WebOptions() + "</body>"
		page := pageStart + head + body + pageEnd
		return ctx.WriteHTML(page)
	})

	return
}
