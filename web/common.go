package web

import "github.com/rohanthewiz/element"

type AppTitle struct {
	Text string // the inner text
	Link string // where to go when clicked
}

func (tl *AppTitle) Render(b *element.Builder) (x any) {
	e, t := b.Funcs()

	e("h1").R(
		e("a", "href", tl.Link, "style", "text-decoration: none; color: inherit").R(
			t(tl.Text),
		),
	)
	return
}
