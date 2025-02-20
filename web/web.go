package web

import (
	_ "fmt"
	"secure_journal/login"

	"github.com/rohanthewiz/element"
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

	chooseMenu := func(menufunc MenuFunc) string {
		return menufunc()
	}

	headerMenu := func(menu MenuFunc) string {
		b := element.NewBuilder()
		e := b.Ele
		t := b.Text

		e("div").R(
			e("head").R(
				e("title").R(
					t("My Journal"),
				),
				e("style").R(
					t("body {background-color: lightblue;} h1 a {text-decoration: none; color: inherit;}"),
				),
			),
			e("body").R(
				e("h1").R(
					t(`<h1><a href="/" style="text-decoration: none; color: inherit;">My Journal</a></h1>`),
				),
				e("div").R(
					t(chooseMenu(menu)),
				),
			),
		)
		return b.String()
	}

	rootHandler := func(ctx rweb.Context) error {
		return ctx.WriteHTML(headerMenu(RegisterMenu))
	}

	s.Get("/", rootHandler)

	s.Get("/register", func(ctx rweb.Context) (err error) {
		b := element.NewBuilder()
		e := b.Ele
		t := b.Text

		e("html").R(
			t(headerMenu(RegisterMenu)),
			e("div").R(
				e("form", "action", "/register", "method", "POST").R(
					e("label", "for", "username").R(t("Username:")),
					e("br"),
					e("input", "type", "username", "id", "username").R(),
					e("br"),
					e("label", "for", "password").R(t("Password:")),
					e("br"),
					e("input", "type", "password", "id", "password").R(),
					e("br"),
					e("label", "for", "confirm_password").R(t("Confirm_Password:")),
					e("br"),
					e("input", "type", "password", "id", "confirm_password").R(),
					e("br"),
					e("input", "type", "submit", "value", "Register"),
				),
			),
		)
		return ctx.WriteHTML(b.String())
	})

	s.Post("/register", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")
		confirm_password := ctx.Request().FormValue("confirm_password")

		if password != confirm_password {
			return errorHandler(ctx, "Registration failed: Passwords don't match!")
		}

		err = login.Register(username, password)
		if err != nil {
			return errorHandler(ctx, "Registration failed:"+err.Error())
		}

		return successHandler(ctx, "Registration Successful!", LogMenu)
	})

	s.Get("/login", func(ctx rweb.Context) (err error) {
		b := element.NewBuilder()
		e := b.Ele
		t := b.Text

		e("html").R(
			t(headerMenu(RegisterMenu)),
			e("div").R(
				e("form", "action", "/login", "method", "POST").R(
					e("label", "for", "username").R(t("Username:")),
					e("br"),
					e("input", "type", "username", "id", "username").R(),
					e("br"),
					e("label", "for", "password").R(t("Password:")),
					e("br"),
					e("input", "type", "password", "id", "password").R(),
					e("br"),
					e("input", "type", "submit", "value", "Login"),
				),
			),
		)
		return ctx.WriteHTML(b.String())
	})

	s.Post("/login", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")
		if password == "" || username == "" {
			return errorHandler(ctx, "Login Failed: You must enter a password")
		}
		err = login.Login(username, password)
		if err != nil {
			return errorHandler(ctx, "Login Failed:"+err.Error())
		}
		return successHandler(ctx, "Welcome to your Journals!", JournalMenu)
	})

	s.Get("/my-journals", func(ctx rweb.Context) (err error) {
		b := element.NewBuilder()
		e := b.Ele
		t := b.Text

		e("html").R(
			t(headerMenu(noMenu)),
			e("p").R(
				t("I can do all things through christ who strengthens me!"),
			),
		)
		return ctx.WriteHTML(b.String())
	})

	s.Get("/log-out", func(ctx rweb.Context) (err error) {
		return rootHandler(ctx)
	})

	s.Get("/delete-user", func(ctx rweb.Context) (err error) {
		b := element.NewBuilder()
		e := b.Ele
		t := b.Text

		e("html").R(
			t(headerMenu(RegisterMenu)),
			e("div").R(
				e("form", "action", "/delete-user", "method", "POST").R(
					e("label", "for", "username").R(t("Username:")),
					e("br"),
					e("input", "type", "username", "id", "username").R(),
					e("br"),
					e("label", "for", "password").R(t("Password:")),
					e("br"),
					e("input", "type", "password", "id", "password").R(),
					e("br"),
					e("input", "type", "submit", "value", "Delete"),
				),
			),
		)
		return ctx.WriteHTML(b.String())
	})

	s.Post("/delete-user", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")

		err = login.Delete(username, password)
		if err != nil {
			return errorHandler(ctx, err.Error())
		}
		return successHandler(ctx, "Deletion Successful!", RegisterMenu)
	})

	//initweb return
	return
}
func successHandler(ctx rweb.Context, successMsg string, menufunc MenuFunc) error {
	b := element.NewBuilder()
	e := b.Ele
	t := b.Text

	e("html").R(
		e("head").R(
			e("title").R(
				t("My Journal"),
			),
			e("style").R(
				t("body {background-color: lightblue;} h1 a {text-decoration: none; color: inherit;}"),
			),
		),
		e("body").R(
			e("h1").R(
				t(`<h1><a href="/" style="text-decoration: none; color: inherit;">My Journal</a></h1>`),
			),
			e("div").R(
				t(menufunc()),
			),
			e("div").R(
				e("p", "style", "color: green").R(
					t(successMsg),
				),
			),
		),
	)
	return ctx.WriteHTML(b.String())
}
func errorHandler(ctx rweb.Context, errorMessage string) error {
	b := element.NewBuilder()
	e := b.Ele
	t := b.Text

	e("html").R(
		e("head").R(
			e("title").R(
				t("My Journal"),
			),
			e("style").R(
				t("body {background-color: lightblue;} h1 a {text-decoration: none; color: inherit;}"),
			),
		),
		e("body").R(
			e("h1").R(
				t(`<h1><a href="/" style="text-decoration: none; color: inherit;">My Journal</a></h1>`),
			),
			e("div").R(
				t(RegisterMenu()),
			),
			e("div").R(
				e("p", "style", "color: red").R(
					t(errorMessage),
				),
				e("a", "href", "/register").R(
					t("Try again"),
				),
			),
		),
	)
	return ctx.WriteHTML(b.String())
}
