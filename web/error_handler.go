package web

import (
	"github.com/rohanthewiz/element"
)

func errHandler(b *element.Builder, errorMessage string, _ ...element.Component) {
	e := b.Ele
	t := b.Text

	e("div").R(
		element.RenderComponents(b),
		e("div").R(
			e("p", "style", "color: red").R(
				t(errorMessage),
			),
			e("a", "href", "/delete-user").R(
				t("Try again"),
			),
		),
	)
}

// --- Error Component ---

type ErrorComp struct {
	Msg string
}

func (ec ErrorComp) Render(b *element.Builder) (x any) {
	e, t := b.Funcs()

	PageMenu{Items: []string{strRegister, strLogin, strDeleteUser}}.Render(b)

	e("div").R(
		e("p", "style", "color: red").R(
			t(ec.Msg),
		),
		e("a", "href", "/delete-user").R(
			t("Try again"),
		),
	)
	return
}
