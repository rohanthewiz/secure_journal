package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func journalRouter(s *rweb.Server) {

	s.Get("/my-journals", func(ctx rweb.Context) (err error) {
		return ctx.WriteHTML(PgLayout(firstJournal))
	})
}

func firstJournal(b *element.Builder, comps ...element.Component) {
	e := b.Ele
	t := b.Text

	e("div").R(
		e("p").R(
			t("I can do all things through christ who strengthens me!"),
		),
	)
}
