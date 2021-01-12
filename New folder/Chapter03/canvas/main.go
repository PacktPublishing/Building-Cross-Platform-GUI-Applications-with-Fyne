package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	a := app.New()

	w := a.NewWindow("Sign")
	w.SetContent(makeSign())
	w.SetPadded(false)
	w.Resize(fyne.NewSize(120, 120))
	w.ShowAndRun()
}

func makeSign() fyne.CanvasObject {
	bg := canvas.NewCircle(color.NRGBA{255, 0, 0, 255})
	bg.StrokeColor = color.White
	bg.StrokeWidth = 5
	bar := canvas.NewRectangle(color.White)

	c := fyne.NewContainerWithoutLayout(bg, bar)
	bg.Resize(fyne.NewSize(100, 100))
	bg.Move(fyne.NewPos(10, 10))
	bar.Resize(fyne.NewSize(80, 20))
	bar.Move(fyne.NewPos(20, 50))
	return c
}
