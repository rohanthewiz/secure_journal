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

	return ctx.WriteHTML(PageLayout(MenuProvider(strRegister, strLogin, strDeleteUser), b.String()))
}
