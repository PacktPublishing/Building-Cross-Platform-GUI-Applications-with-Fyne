package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	a := app.New()

	w := a.NewWindow("Pixels")
	w.SetContent(canvas.NewRasterWithPixels(generate))
	w.SetPadded(false)
	w.Resize(fyne.NewSize(40, 40))
	w.ShowAndRun()
}

func generate(x, y, w, h int) color.Color {
	if (x/20)%2 == (y/20)%2 {
		return color.White
	}

	return color.Black
}
