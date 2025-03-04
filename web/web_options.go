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

	e("div", "id", "menu-container").R(
		e("style").R(
			t(`
				ul {
					padding: 0;
					margin: 0;
					list-style: none;
				}
				li {
					margin-bottom: 8px;
				}
				.horizontal li {
					display: inline-block;
					margin-right: 15px;
					margin-bottom: 0;
				}
				.vertical li {
					display: block;
				}
				.toggle-button {
					margin-bottom: 10px;
					padding: 5px 10px;
					cursor: pointer;
				}
			`),
		),
		e("button", "class", "toggle-button", "onclick", "toggleMenuLayout()").R(
			t("Toggle Layout"),
		),
		e("ul", "id", "menu", "class", "horizontal").R(
			func() []any {
				var items []any
				for _, ele := range str {
					items = append(items, t(listCreate(ele)))
				}
				return items
			}()...,
		),
		e("script").R(
			t(`
				function toggleMenuLayout() {
					const menu = document.getElementById('menu');
					if (menu.classList.contains('horizontal')) {
						menu.classList.remove('horizontal');
						menu.classList.add('vertical');
					} else {
						menu.classList.remove('vertical');
						menu.classList.add('horizontal');
					}
				}
			`),
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
