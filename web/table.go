package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func tableRouter(s *rweb.Server) {
	s.Get("/table", func(ctx rweb.Context) error {
		tableMenu := PageMenu{Items: []string{strLogout}}
		return ctx.WriteHTML(PgLayout(tableMenu, Table{Borders: true}))
	})
}

type Table struct {
	Borders bool
}

func (tbl Table) Render(b *element.Builder) (x any) {
	e, t := b.Funcs()

	border := ""
	if tbl.Borders {
		border = " border=1"
	}

	e("table"+border).R(
		e("thead").R(
			e("tr").R(
				e("th").R(t("Col1")),
				e("th").R(t("Col2")),
			),
		),
		e("tbody").R(
			e("tr").R(
				e("td").R(t("One")),
				e("td").R(t("Two")),
			),
		),
	)
	return
}
