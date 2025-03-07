package web

import (
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

type MenuFunc func() string

func InitWeb() (s *rweb.Server) {
	s = rweb.NewServer(
		rweb.ServerOptions{
			Address: "localhost:8000",
			Verbose: true,
		},
	)
	// HANDLERS
	rootHandler := func(ctx rweb.Context) error {
		rootMenu := func(b *element.Builder, comps ...element.Component) {
			Menu(b, strRegister, strLogin, strDeleteUser)
		}
		return ctx.WriteHTML(PgLayout(rootMenu))
	}

	s.Get("/", rootHandler)

	registerRouter(s)

	loginRouter(s)

	journalRouter(s)

	DeleteRouter(s)

	s.Get("/log-out", func(ctx rweb.Context) (err error) {
		return rootHandler(ctx)
	})

	// initweb return
	return
}
