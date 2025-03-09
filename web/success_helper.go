package web

import (
	"github.com/rohanthewiz/element"
)

func successHandler(b *element.Builder, successMsg string, _ ...element.Component) {
	e := b.Ele
	t := b.Text

	e("div").R(
		element.RenderComponents(b),
		e("div").R(
			e("p", "style", "color: green").R(
				t(successMsg),
			),
		),
	)
}
