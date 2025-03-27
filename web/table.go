package web

import (
	"database/sql"
	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
	"secure_journal/login" // Replace with actual import path
)

func tableHandler(s *rweb.Server, db *sql.DB) {
	s.Get("/table", func(ctx rweb.Context) error {
		tableMenu := PageMenu{Items: []string{strLogout, strRegister, strDeleteUser}}
		var str string
		// Fetch users from the database
		users, err := login.Table(db)
		if err != nil {
			// Handle error, maybe log it or return an error page
			str = "Failed to retrieve users"
			return ctx.WriteHTML(PgLayout(tableMenu, ErrorComp{Msg: str}))
		}

		return ctx.WriteHTML(PgLayout(tableMenu, Table{
			Borders: true,
			Users:   users,
		}))
	})
}

type Table struct {
	Borders bool
	Users   []string
	Error   string
}

func (tbl Table) Render(b *element.Builder) (x any) {
	e, t := b.Funcs()
	border := ""
	if tbl.Borders {
		border = " border=1"
	}

	// If there's an error, render an error message
	if tbl.Error != "" {
		return e("div.error").R(t(tbl.Error))
	}

	// Dynamically create table rows based on users
	tbody := e("tbody")

	e("table" + border).R(
		e("thead").R(
			e("tr").R(
				e("th").R(t("Username")),
				e("th").R(t("Status")),
			),
			func() any {
				for _, username := range tbl.Users {
					tbody.R(
						e("tr").R(
							e("td").R(t(username)),
							e("td").R(t("Active")), // You can add more columns as needed
						),
					)
				}
				return tbody
			}(),
		),
	)
	return
}
