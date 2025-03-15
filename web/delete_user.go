package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
	"secure_journal/login"
)

func DeleteRouter(s *rweb.Server) {
	s.Get("/delete-user", func(ctx rweb.Context) (err error) {
		deleteUserMenu := PageMenu{Items: []string{strLogin, strRegister}}
		return ctx.WriteHTML(PgLayout(deleteUserMenu, DeleteUserForm{}))
	})

	s.Post("/delete-user", func(ctx rweb.Context) (err error) {
		username := ctx.Request().FormValue("username")
		password := ctx.Request().FormValue("password")

		if username == "" || password == "" {
			return ctx.WriteHTML(PgLayout(ErrorComp{"You must type username && password"}))
		}

		err = login.Delete(username, password)

		if err != nil {
			return ctx.WriteHTML(PgLayout(ErrorComp{"username does not exist!"}))
		}

		return ctx.WriteHTML(PgLayout(SuccessComp{"User Deleted!"}))
	})
}

type DeleteUserForm struct{}

func (d DeleteUserForm) Render(b *element.Builder) (x any) {
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
	return
}
