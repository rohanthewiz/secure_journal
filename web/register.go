package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
	"secure_journal/login"
)

func registerRouter(s *rweb.Server) {

	s.Get("/register", func(ctx rweb.Context) (err error) {
		registerMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strLogin, strDeleteUser)
		}
		return ctx.WriteHTML(PgLayout(registerMenu, RegisterGETHandler))
	})

	s.Post("/register", func(ctx rweb.Context) (err error) {
		successMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strMyJournal, strLogout)
		}
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")
		confirm_password := ctx.Request().FormValue("confirm_password")

		if username == "" || password == "" || confirm_password == "" {
			return errorHandler(ctx, "You must fill out all boxes!")
		}
		if password != confirm_password {
			return errorHandler(ctx, "Registration failed: Passwords don't match!")
		}

		err = login.Register(username, password)
		if err != nil {
			return errorHandler(ctx, "Registration failed:"+err.Error())
		}

		return successHandler(ctx, "Registration Successful!", successMenu)
	})
}

// This function has been fixed
func RegisterGETHandler(b *element.Builder, comp ...element.Component) {
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

}
