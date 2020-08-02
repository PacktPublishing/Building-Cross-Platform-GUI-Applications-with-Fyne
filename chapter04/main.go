package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Image Browser")

	w.SetContent(makeUI())
	w.Resize(fyne.NewSize(480, 360))
	w.ShowAndRun()
}
