package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func loginForm(s *rweb.Server) {

	s.Get("/login", func(ctx rweb.Context) (err error) {
		b := element.NewBuilder()
		e := b.Ele
		t := b.Text

		e("html").R(
			t("TODO - FIX"),
			t(PageLayout(RegisterMenu, "")),
			e("div").R(
				e("form", "action", "/login", "method", "POST").R(
					e("label", "for", "username").R(t("Username:")),
					e("br"),
					e("input", "type", "username", "id", "username").R(),
					e("br"),
					e("label", "for", "password").R(t("Password:")),
					e("br"),
					e("input", "type", "password", "id", "password").R(),
					e("br"),
					e("input", "type", "submit", "value", "Login"),
				),
			),
		)
		return ctx.WriteHTML(b.String())
	})
}
