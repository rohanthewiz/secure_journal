package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

// This function has been fixed
func RegisterGETHandler(ctx rweb.Context) (err error) {
	b := element.NewBuilder()
	e := b.Ele
	t := b.Text

	e("div").R(
		e("form", "action", "/register", "method", "POST").R(
			e("label", "for", "username").R(t("Username:")),
			e("br"),
			e("input", "type", "username", "id", "username").R(),
			e("br"),
			e("label", "for", "password").R(t("Password:")),
			e("br"),
			e("input", "type", "password", "id", "password").R(),
			e("br"),
			e("label", "for", "confirm_password").R(t("Confirm_Password:")),
			e("br"),
			e("input", "type", "password", "id", "confirm_password").R(),
			e("br"),
			e("input", "type", "submit", "value", "Register"),
		),
	)

	return ctx.WriteHTML(PageLayout(RegisterMenu, b.String()))
}
