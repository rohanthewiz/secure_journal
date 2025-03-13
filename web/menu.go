package web

import (
	"strings"

	"github.com/rohanthewiz/element"
)

// List of strings for the menus
var strRegister = "Register"
var strLogin = "Login"
var strDeleteUser = "Delete-User"
var strMyJournal = "My-Journals"
var strLogout = "Logout"

type PageMenu struct {
	Items []string
}

func (m PageMenu) Render(b *element.Builder) (x any) {
	e, t := b.Funcs()

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
		func() any {
			if len(m.Items) > 1 {
				return e("button", "class", "toggle-button", "onclick", "toggleMenuLayout()").R(
					t("Toggle Layout"),
				)
			}
			return nil
		}(),
		e("ul", "id", "menu", "class", "horizontal").R(
			func() []any {
				var items []any
				for _, ele := range m.Items {
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
	return
}

/*func Menu(b *element.Builder, str ...string) {
}
*/

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
