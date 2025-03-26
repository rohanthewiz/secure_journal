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
          display: flex;
          gap: 15px; /* Spacing between menu items */
          transition: max-height 0.3s ease-out, opacity 0.3s ease-out;
          max-height: 0;
          opacity: 0;
          overflow: hidden;
          flex-wrap: wrap; /* Ensures responsiveness */
        }

        .horizontal {
          max-height: 50px; /* Adjust height based on your design */
          opacity: 1;
        }

        li {
          margin: 0;
        }

        li a {
          text-decoration: none;
          color: black;
          padding: 8px 12px;
          border-radius: 5px;
          transition: background-color 0.2s ease-in-out;
        }

        li a:hover {
          background-color: lightgray;
        }

        .toggle-button {
          margin-bottom: 10px;
          padding: 5px 10px;
          cursor: pointer;
          transition: background-color 0.2s ease-in-out, transform 0.2s ease-in-out;
        }

        .toggle-button:active {
          transform: scale(0.95);
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
          const button = document.getElementById('toggle-button');

          if (menu.classList.contains('horizontal')) {
            menu.style.opacity = "0";
            menu.style.maxHeight = "0";
            setTimeout(() => {
              menu.classList.remove('horizontal');
              menu.classList.add('empty-menu');
              button.innerHTML = "+ Menu";
            }, 300);
          } else {
            menu.classList.remove('empty-menu');
            menu.classList.add('horizontal');
            menu.style.maxHeight = "50px"; // Adjust based on content
            menu.style.opacity = "1";
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
