package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/data/binding"
	"fyne.io/fyne/widget"
)

func makeUI() fyne.CanvasObject {
	f := binding.NewFloat()

	prog := widget.NewProgressBarWithData(f)
	slide := widget.NewSliderWithData(0, 1, f)
	slide.Step = 0.01
	btn := widget.NewButton("Set to 0.5", func() {
		f.Set(0.5)
	})

	return container.NewVBox(prog, slide, btn)
}

func main() {
	a := app.New()
	w := a.NewWindow("Widget Binding")

	w.SetContent(makeUI())
	w.ShowAndRun()
}
