package main

import (
	"image/color"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/storage"
)

func loadImage(u fyne.URI) fyne.CanvasObject {
	read, err := storage.OpenFileFromURI(u)
	if err != nil {
		log.Println("Error opening image", err)
		return canvas.NewRectangle(color.Black)
	}
	res, err := fyne.LoadResourceFromURI(read)
	if err != nil {
		log.Println("Error reading image", err)
		return canvas.NewRectangle(color.Black)
	}

	img := canvas.NewImageFromResource(res)
	img.FillMode = canvas.ImageFillContain
	return img
}

func makeImageItem(u fyne.URI) fyne.CanvasObject {
	label := canvas.NewText(u.Name(), color.Gray{128})
	label.Alignment = fyne.TextAlignCenter
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, label, nil, nil),
		label, loadImage(u))
}
