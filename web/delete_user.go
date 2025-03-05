package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func DeleteUser(ctx rweb.Context) (err error) {
	b := element.NewBuilder()
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

	return ctx.WriteHTML(PageLayout(MenuProvider(strLogin), b.String()))
}
