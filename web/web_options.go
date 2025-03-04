package web

import (
	"github.com/rohanthewiz/element"
	"strings"
)

// List of strings for the menus
var strRegister = "Register"
var strLogin = "Login"
var strDeleteUser = "Delete-User"
var strMyJournal = "My-Journals"
var strLogout = "Logout"

func MenuProvider(items ...string) MenuFunc {
	return func() string {
		return Menu(items...)
	}
}

func Menu(str ...string) string {
	b := element.NewBuilder()
	e, t := b.Ele, b.Text

	e("div").R(
		e("style").R(
			t(`ul {
					padding: 0;
					margin: 0;
					list-style: none;
				}
				li {
					display: inline-block;
					margin-right: 15px;
				}`),
		),
		e("ul").R(
			func() []any {
				var items []any
				for _, ele := range str {
					items = append(items, t(listCreate(ele)))
				}
				return items
			}()...,
		),
	)
	return b.String()
}

func listCreate(str string) string {
	b := element.NewBuilder()
	e, t := b.Ele, b.Text

	lowStr := strings.ToLower(str)
	e("li").R(
		e("a", "href", "/"+lowStr).R(
			t(str),
		),
	)

	return b.String()
}
