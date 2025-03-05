package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func loginForm(s *rweb.Server) {
	s.Get("/login", func(ctx rweb.Context) (err error) {
		loginMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strRegister, strDeleteUser)
		}
		return ctx.WriteHTML(PgLayout(loginMenu, LoginTitle, LoginPageBody))
	})

	/*	// TODO - Let's work on this one so it uses element.Component
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
	*/
}

// LoginTitle is an example of an Element Component
func LoginTitle(b *element.Builder, comps ...element.Component) {
	e, t := b.Ele, b.Text

	e("h1").R(
		t(`<h1><a href="/" style="text-decoration: none; color: inherit;">My Journal</a></h1>`),
	)
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
		func() (x any) {
			for _, comp := range comps {
				comp(b)
			}
			return
		}(),
	)
}
