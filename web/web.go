package web

import (
	"database/sql"
	"github.com/rohanthewiz/rweb"
)

type MenuFunc func() string

func InitWeb(db *sql.DB) (s *rweb.Server) {
	s = rweb.NewServer(
		rweb.ServerOptions{
			Address: "localhost:8000",
			Verbose: true,
		},
	)
	// HANDLERS
	rootHandler := func(ctx rweb.Context) error {
		rootMenu := PageMenu{Items: []string{strRegister, strLogin, strDeleteUser, strTable}}
		return ctx.WriteHTML(PgLayout(rootMenu))
	}

	s.Get("/", rootHandler)

	registerHandler(s, db)

	loginHandler(s, db)

	journalHandler(s)

	tableHandler(s, db)

	DeleteHandler(s, db)

	s.Get("/logout", func(ctx rweb.Context) (err error) {
		return rootHandler(ctx)
	})

	// initweb return
	return
}
