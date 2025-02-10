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

	rootHandler := func(ctx rweb.Context) error { // in-line func or anonymous function
		head := "<head><title>My Journal</title>"
		head += "<style>body {background-color: lightblue;}</style></head>"
		body := "<body><h1>My Journal</h1>" + WebOptions() + "</body>"
		pageStart := "<html>"
		pageEnd := "</html>"

		page := pageStart + head + body + pageEnd
		fmt.Println(page)

		return ctx.WriteHTML(page)
	}

	s.Get("/", rootHandler)

	s.Get("/register", func(ctx rweb.Context) (err error) {
		// Do some action
		_ = login.Register("password")

		// Return a resp
		head := "<head><title>My Journal</title>"
		head += "<style>body {background-color: lightblue;}</style></head>"
		body := "<body><h1>My Journal</h1>" + WebOptions() +
			`<p style="color: maroon">Register</p>` +
			"</body>"
		pageStart := "<html>"
		pageEnd := "</html>"

		page := pageStart + head + body + pageEnd

		return ctx.WriteHTML(page)
	})

	return
}
