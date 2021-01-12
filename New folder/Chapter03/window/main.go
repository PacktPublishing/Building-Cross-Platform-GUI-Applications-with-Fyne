package main

import "fyne.io/fyne/app"

func main() {
	a := app.New()

	w := a.NewWindow("Hello")
	w.ShowAndRun()
}
