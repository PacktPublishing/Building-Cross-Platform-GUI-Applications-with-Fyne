package main

import (
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type submitEntry struct {
	widget.Entry
}

func newSubmitEntry() *submitEntry {
	e := &submitEntry{}
	e.ExtendBaseWidget(e)
	return e
}

func (s *submitEntry) TypedKey(k *fyne.KeyEvent) {
	if k.Name == fyne.KeyReturn {
		log.Println("Submit data", s.Text)
		s.SetText("")
		return
	}

	s.Entry.TypedKey(k)
}

func main() {
	a := app.New()
	w := a.NewWindow("Submit Entry")

	w.SetContent(newSubmitEntry())
	w.ShowAndRun()
}
