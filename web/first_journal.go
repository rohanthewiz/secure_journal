package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func journalHandler(s *rweb.Server) {

	s.Get("/my-journals", func(ctx rweb.Context) (err error) {
		journalMenu := PageMenu{Items: []string{strMyJournal, strLogout}}
		return ctx.WriteHTML(PgLayout(journalMenu, firstJournal{}))
	})
}

type firstJournal struct{}

func (f firstJournal) Render(b *element.Builder) (x any) {
	e := b.Ele
	t := b.Text

	e("div").R(
		e("p").R(
			t("I can do all things through christ who strengthens me!"),
		),
	)
	return
}
