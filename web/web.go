package web

import (
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
		return ctx.WriteHTML(PgLayout())
	}

	s.Get("/", rootHandler)

	registerRouter(s)
	//get and post are within functions
	loginRouter(s)

	journalRouter(s)

	DeleteRouter(s)

	s.Get("/log-out", func(ctx rweb.Context) (err error) {
		return rootHandler(ctx)
	})

	// initweb return
	return
}
