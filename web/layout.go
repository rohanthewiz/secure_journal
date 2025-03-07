package web

import "github.com/rohanthewiz/element"

// PgLayout is the new **top level** page layout function
// It is the only rendering function that should return a string
// Element components should take in the builder from here
func PgLayout(comps ...element.Component) string {
	b := element.NewBuilder()
	e, t := b.Ele, b.Text

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
			element.RenderComponents(b, comps...),
		),
	)
	// Return the whole page
	return b.String()
}
