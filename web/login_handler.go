package web

import (
	"secure_journal/login"

	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func loginRouter(s *rweb.Server) {
	s.Get("/login", func(ctx rweb.Context) (err error) {
		loginMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strRegister, strDeleteUser)
		}
		return ctx.WriteHTML(PgLayout(loginMenu, LoginTitle, LoginPageBody))
	})

	// TODO - Let's work on this one so it uses element.Component
	s.Post("/login", func(ctx rweb.Context) (err error) {
		password := ctx.Request().FormValue("password")
		username := ctx.Request().FormValue("username")
		var str string

		successMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strMyJournal, strLogout)
			successHandler(b, str)
		}
		errorBody := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strRegister, strLogin, strDeleteUser)
			errHandler(b, str)
		}

		if password == "" || username == "" {
			str = "You must have a username and password"
			return ctx.WriteHTML(PgLayout(LoginTitle, errorBody))
		}
		err = login.Login(username, password)
		if err != nil {
			str = "Login gailed" + err.Error()
			return ctx.WriteHTML(PgLayout(LoginTitle, errorBody))
		}
		str = "Login successful!"
		return ctx.WriteHTML(PgLayout(LoginTitle, successMenu))
	})

}

// LoginTitle is an example of an Element Component
func LoginTitle(b *element.Builder, _ ...element.Component) {
	e, t := b.Ele, b.Text

	e("h1").R(
		t(`<h1><a href="/" style="text-decoration: none; color: inherit;">My Journal</a></h1>`),
	)
	element.RenderComponents(b)
}

// LoginPageBody is an example of an Element Component
func LoginPageBody(b *element.Builder, comps ...element.Component) {
	e, t := b.Ele, b.Text

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

		// We _could_ render out other comps wherever, if ever, we want
		element.RenderComponents(b),
	)
}
