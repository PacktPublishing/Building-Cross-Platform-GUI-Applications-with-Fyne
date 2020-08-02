package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func makeImageGrid() fyne.CanvasObject {
	items := []fyne.CanvasObject{}
	for _ = range []int{1, 2, 3} {
		img := makeImageItem()
		items = append(items, img)
	}

	cellSize := fyne.NewSize(160, 120)
	return fyne.NewContainerWithLayout(layout.NewGridWrapLayout(cellSize), items...)
}

func makeStatus() fyne.CanvasObject {
	return canvas.NewText("status", color.Gray{128})
}

func makeUI() fyne.CanvasObject {
	status := makeStatus()
	content := makeImageGrid()
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, status, nil, nil),
		status, content)
}
