package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func makeUI() fyne.CanvasObject {
	f := binding.NewFloat()

	prog := widget.NewProgressBarWithData(f)
	slide := widget.NewSliderWithData(0, 1, f)
	slide.Step = 0.01
	btn := widget.NewButton("Set to 0.5", func() {
		_ = f.Set(0.5)
	})

	return container.NewVBox(prog, slide, btn)
}

func main() {
	a := app.New()
	w := a.NewWindow("Widget Binding")

	w.SetContent(makeUI())
	w.ShowAndRun()
}
