package web

import (
	"github.com/rohanthewiz/element"
)

type SuccessComp struct {
	Msg string
}

func (s SuccessComp) Render(b *element.Builder) (x any) {
	e, t := b.Funcs()

	menu := PageMenu{Items: []string{strMyJournal, strLogout}}
	menu.Render(b)

	e("div").R(
		e("p", "style", "color: green").R(
			t(s.Msg),
		),
	)
	return
}
