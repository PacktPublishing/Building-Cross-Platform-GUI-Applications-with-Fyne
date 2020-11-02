package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	a := app.New()

	w := a.NewWindow("Snake")
	w.SetContent(setupGame())
	w.Resize(fyne.NewSize(200, 200))
	w.SetFixedSize(true)
	w.ShowAndRun()
}

func setupGame() fyne.CanvasObject {
	var segments []fyne.CanvasObject
	for i := 0; i < 10; i++ {
		r := canvas.NewRectangle(&color.RGBA{G: 0x66, A: 0xff})
		r.Resize(fyne.NewSize(10, 10))
		r.Move(fyne.NewPos(90, 50+(i*10)))
		segments = append(segments, r)
	}

	return fyne.NewContainerWithoutLayout(segments...)
}
