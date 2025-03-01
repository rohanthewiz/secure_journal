package web

import (
	"fmt"
	"secure_journal/login"

	"github.com/rohanthewiz/rweb"
)

type MenuFunc func() string

func InitWeb() (s *rweb.Server) {
	s = rweb.NewServer(
		rweb.ServerOptions{
			Address: "localhost:8000",
			Verbose: true,
		},
	)
	// HANDLERS

	rootHandler := func(ctx rweb.Context) error {
		return ctx.WriteHTML(PageLayout(RegisterMenu, ""))
	}

	s.Get("/", rootHandler)

	s.Get("/register", RegisterGETHandler)

	s.Post("/register", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")
		confirm_password := ctx.Request().FormValue("confirm_password")
		fmt.Printf("%s, %s, %s\n", username, password, confirm_password)
		if password != confirm_password {
			return errorHandler(ctx, "Registration failed: Passwords don't match!", RegisterMenu)
		}

		err = login.Register(username, password)
		if err != nil {
			return errorHandler(ctx, "Registration failed:"+err.Error(), RegisterMenu)
		}

		return successHandler(ctx, "Registration Successful!", LogMenu)
	})

	//get and post are within functions
	loginForm(s)

	firstJournal(s)

	s.Get("/log-out", func(ctx rweb.Context) (err error) {
		return rootHandler(ctx)
	})

	s.Get("/delete-user", DeleteUser)

	s.Post("/delete-user", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")

		err = login.Delete(username, password)
		if err != nil {
			return errorHandler(ctx, err.Error(), RegisterMenu)
		}
		return successHandler(ctx, "Deletion Successful!", RegisterMenu)
	})

	// initweb return
	return
}
