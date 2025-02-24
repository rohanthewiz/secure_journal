package web

import "github.com/rohanthewiz/element"

func PageLayout(menu MenuFunc, yourStuff string) string {
	b := element.NewBuilder()
	e := b.Ele
	t := b.Text

	e("html").R(
		e("head").R(
			e("title").R(
				t("My Journal"),
			),
			e("style").R(
				t("body {background-color: lightblue;} h1 a {text-decoration: none; color: inherit;}"),
			),
		),
		e("body").R(
			e("h1").R(
				t(`<h1><a href="/" style="text-decoration: none; color: inherit;">My Journal</a></h1>`),
			),
			e("div").R(
				t(menu()),
			),

			// Add your things
			t(yourStuff),
		),
	)

	return b.String()
}
