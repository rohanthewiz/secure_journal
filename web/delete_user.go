package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
	"secure_journal/login"
)

func DeleteRouter(s *rweb.Server) {
	s.Get("/delete-user", func(ctx rweb.Context) (err error) {
		deleteUserMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strRegister, strLogin)
		}
		return ctx.WriteHTML(PgLayout(deleteUserMenu, DeleteUserForm))
	})

	s.Post("/delete-user", func(ctx rweb.Context) (err error) {
		username := ctx.Request().FormValue("username")
		password := ctx.Request().FormValue("password")
		var str string

		successMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strMyJournal, strLogout)
		}
		errorBody := func(b *element.Builder, comps ...element.Component) {
			errHandler(b, str)
		}

		if username == "" || password == "" {
			str = "You must type username && password"
			return ctx.WriteHTML(PgLayout(LoginTitle, errorBody))
		}

		err = login.Delete(username, password)

		if err != nil {
			str = "username does not exist!" + err.Error()
			return ctx.WriteHTML(PgLayout(LoginTitle, errorBody))
		}

		return successHandler(ctx, "Deletion Successful!", successMenu)
	})
}

func DeleteUserForm(b *element.Builder, comps ...element.Component) {
	e := b.Ele
	t := b.Text
	e("div").R(
		e("form", "action", "/delete-user", "method", "POST").R(
			e("label", "for", "username").R(t("Username:")),
			e("br"),
			e("input", "type", "text", "id", "username", "name", "username").R(),
			e("br"),
			e("label", "for", "password").R(t("Password:")),
			e("br"),
			e("input", "type", "password", "id", "password", "name", "password").R(),
			e("br"),
			e("input", "type", "submit", "value", "Delete"),
		),
	)
}
