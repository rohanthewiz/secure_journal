package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
	"secure_journal/login"
)

func registerRouter(s *rweb.Server) {

	s.Get("/register", func(ctx rweb.Context) (err error) {
		registerMenu := PageMenu{Items: []string{strLogin, strDeleteUser}}
		return ctx.WriteHTML(PgLayout(registerMenu, RegisterForm{}))
	})

	s.Post("/register", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")
		confirm_password := ctx.Request().FormValue("confirm_password")

		if username == "" || password == "" || confirm_password == "" {
			return ctx.WriteHTML(PgLayout(ErrorComp{"You must have a username and password"}))
		}
		if password != confirm_password {
			return ctx.WriteHTML(PgLayout(ErrorComp{"Your passwords do not match!"}))
		}

		err = login.Register(username, password)
		if err != nil {
			return ctx.WriteHTML(PgLayout(ErrorComp{"Registration failed:" + err.Error()}))
		}
		return ctx.WriteHTML(PgLayout(SuccessComp{"Registration successful!"}))
	})
}

type RegisterForm struct{}

// This function has been fixed
func (r RegisterForm) Render(b *element.Builder) (x any) {
	e := b.Ele
	t := b.Text

	e("div").R(
		e("form", "action", "/register", "method", "POST").R(
			e("label", "for", "username").R(t("Username:")),
			e("br"),
			e("input", "type", "text", "id", "username", "name", "username").R(),
			e("br"),
			e("label", "for", "password").R(t("Password:")),
			e("br"),
			e("input", "type", "password", "id", "password", "name", "password").R(),
			e("br"),
			e("label", "for", "confirm_password").R(t("Confirm_Password:")),
			e("br"),
			e("input", "type", "password", "id", "confirm_password", "name", "confirm_password").R(),
			e("br"),
			e("input", "type", "submit", "value", "Register"),
		),
	)
	return
}
