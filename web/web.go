package web

import (
	"fmt"
	"secure_journal/login"

	"github.com/rohanthewiz/rweb"
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
	body := "<body><h1>My Journal</h1>" + WebOptions() + "</body>"
	pageStart := "<html>"
	pageEnd := "</html>"

	rootHandler := func(ctx rweb.Context) error { // in-line func or anonymous function

		page := pageStart + head + body + pageEnd
		fmt.Println(page)

		return ctx.WriteHTML(page)
	}

	s.Get("/", rootHandler)

	s.Get("/register", func(ctx rweb.Context) (err error) {
		// Return a resp
		body = "<body><h1>My Journal</h1>" + WebOptions() +
			`<p style="color: navy">Register</p>` +
			`<form action="/register" method="POST">
        <label for="username">Username:</label><br>
        <input type="text"><br>
        <label for="password">Password:</label><br>
        <input type="text"><br>
        <input type="submit" value="Register">
      </form>` +
			"</body>"

		page := pageStart + head + body + pageEnd

		return ctx.WriteHTML(page)
	})

	s.Post("/register", func(ctx rweb.Context) (err error) {
		// username := ctx.Request().FormValue("username")
		password := ctx.Request().FormValue("password")

		err = login.Register(password)
		if err != nil {
			return err
		}
		//currently redirects to http://localhost:8000/register but its blank
		return ctx.Redirect(200, "/register")
	})

	return
}
