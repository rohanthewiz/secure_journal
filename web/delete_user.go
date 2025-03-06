package web

import (
	"secure_journal/login"

	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func DeleteRouter(s *rweb.Server) {

	s.Get("/delete-user", func(ctx rweb.Context) (err error) {
		deleteUserMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strRegister, strLogin)
		}
		return ctx.WriteHTML(PgLayout(deleteUserMenu, DeleteUserForm))
	})

	s.Post("/delete-user", func(ctx rweb.Context) (err error) {
		successMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strMyJournal, strLogout)
		}
		errorMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strRegister, strDeleteUser)
		}

		username := ctx.Request().FormValue("username")
		password := ctx.Request().FormValue("password")

		err = login.Delete(username, password)
		if err != nil {
			return errorHandler(ctx, err.Error(), errorMenu)
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
			e("input", "type", "username", "id", "username").R(),
			e("br"),
			e("label", "for", "password").R(t("Password:")),
			e("br"),
			e("input", "type", "password", "id", "password").R(),
			e("br"),
			e("input", "type", "submit", "value", "Delete"),
		),
	)
}
