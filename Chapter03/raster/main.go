package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
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
