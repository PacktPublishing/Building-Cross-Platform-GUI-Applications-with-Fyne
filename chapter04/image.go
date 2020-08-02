package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func makeImageItem() fyne.CanvasObject {
	label := canvas.NewText("label", color.Gray{128})
	label.Alignment = fyne.TextAlignCenter
	img := canvas.NewRectangle(color.Black)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, label, nil, nil),
		label, img)
}
