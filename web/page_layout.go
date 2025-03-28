package web

import "github.com/rohanthewiz/element"

// PgLayout is the new **top level** page layout function
// It is the only rendering function that should return a string
// Element components should take in the builder from here
func PgLayout(comps ...element.Component) string {
	b := element.NewBuilder()
	e, t := b.Funcs()

	appTitle := AppTitle{Text: "My Journal", Link: "/"}

	e("html").R(
		e("head").R(
			e("title").R(
				t("My Journal - Tab title"),
			),
			e("style").R(
				t("body {background-color: slategrey;} h1 a {text-decoration: none; color: inherit;}"),
			),
		),
		e("body").R(
			appTitle.Render(b), // The app title will be the same for all pages
			element.RenderComponents(b, comps...),
		),
	)
	// Return the whole page
	return b.String()
}
