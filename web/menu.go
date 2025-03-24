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
				.empty-menu li {
					display: none;
				}
				.toggle-button {
					margin-bottom: 10px;
					padding: 2px 6px;
					cursor: pointer;
				}
			`),
		),
		func() any {
			if len(m.Items) > 1 {
				return e("button", "id", "toggle-button", "class", "toggle-button", "style", "font-size:12px;", "onclick", "toggleMenu()").R(
					t("+ Menu"),
				)
			}
			return nil
		}(),
		e("ul", "id", "menu", "class", "empty-menu").R(
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
				function toggleMenu() {
					const menu = document.getElementById('menu');
          const button = document.getElementById('toggle-button')
					if (menu.classList.contains('horizontal')) {
						menu.classList.remove('horizontal');
						menu.classList.add('empty-menu');
            button.innerHTML = "+ Menu";
					} else {
						menu.classList.remove('empty-menu');
						menu.classList.add('horizontal');
            button.innerHTML = "Menu &#215;";
					}
          console.log('Button text is now:', button.innerHTML);
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
