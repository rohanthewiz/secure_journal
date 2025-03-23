package web

import (
	"secure_journal/login"

	"database/sql"
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func loginRouter(s *rweb.Server, db *sql.DB) {
	s.Get("/login", func(ctx rweb.Context) (err error) {
		loginMenu := PageMenu{Items: []string{strRegister, strDeleteUser}}
		return ctx.WriteHTML(PgLayout(loginMenu, LoginPageBody{}))
	})

	s.Post("/login", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")
		var str string

		if password == "" || username == "" {
			str = "You must have a username and password"
			return ctx.WriteHTML(PgLayout(ErrorComp{Msg: str}))
		}

		err = login.Login(db, username, password)
		if err != nil {
			str = "Login failed" + err.Error()
			return ctx.WriteHTML(PgLayout(ErrorComp{Msg: str}))
		}

		return ctx.WriteHTML(PgLayout(SuccessComp{Msg: "Login successful!"}))
	})
}

// LoginPageBody defines the component for the body of the login page
type LoginPageBody struct{}

func (l LoginPageBody) Render(b *element.Builder) (x any) {
	e, t := b.Funcs()

	e("div").R(
		e("form", "action", "/login", "method", "POST").R(
			e("label", "for", "username").R(t("Username:")),
			e("br"),
			e("input", "type", "username", "id", "username", "name", "username").R(),
			e("br"),
			e("label", "for", "password").R(t("Password:")),
			e("br"),
			e("input", "type", "password", "id", "password", "name", "password").R(),
			e("br"),
			e("input", "type", "submit", "value", "Login"),
		),
	)
	return
}
