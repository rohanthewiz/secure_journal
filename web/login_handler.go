package web

import (
	"secure_journal/login"

	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func loginForm(s *rweb.Server) {

	s.Get("/login", func(ctx rweb.Context) (err error) {
		b := element.NewBuilder()
		e := b.Ele
		t := b.Text

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
		return ctx.WriteHTML(PageLayout(MenuProvider(strRegister, strDeleteUser), b.String()))
	})

	s.Post("/login", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")
		if password == "" || username == "" {
			return errorHandler(ctx, "Login Failed: You must enter a password", MenuProvider(strRegister, strDeleteUser))
		}
		err = login.Login(username, password)
		if err != nil {
			return errorHandler(ctx, "Login Failed:"+err.Error(), MenuProvider(strRegister, strDeleteUser))
		}
		return successHandler(ctx, "Welcome to your Journals!", MenuProvider(strMyJournal, strLogout))
	})
}
