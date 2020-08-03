package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func makeImageItem(u fyne.URI) fyne.CanvasObject {
	label := canvas.NewText(u.Name(), color.Gray{128})
	label.Alignment = fyne.TextAlignCenter
	img := canvas.NewRectangle(color.Black)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, label, nil, nil),
		label, img)
}
