package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func firstJournal(s *rweb.Server) {

	s.Get("/my-journals", func(ctx rweb.Context) (err error) {
		b := element.NewBuilder()
		e := b.Ele
		t := b.Text

		e("html").R(
			e("p").R(
				t("I can do all things through christ who strengthens me!"),
			),
		)
		return ctx.WriteHTML(PageLayout(noMenu, b.String()))
	})
}
