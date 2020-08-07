package main

import (
	"image"
	"image/color"
	"log"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/storage"

	"github.com/nfnt/resize"
)

type bgImageLoad struct {
	uri fyne.URI
	img *canvas.Image
}

var loads = make(chan bgImageLoad, 1024)

func scaleImage(img image.Image) image.Image {
	return resize.Thumbnail(320, 240, img, resize.Lanczos3)
}

func doLoadImage(u fyne.URI, img *canvas.Image) {
	read, err := storage.OpenFileFromURI(u)
	if err != nil {
		log.Println("Error opening image", err)
		return
	}

	defer read.Close()
	raw, _, err := image.Decode(read)
	if err != nil {
		log.Println("Error decoding image", err)
		return
	}

	img.Image = scaleImage(raw)
	img.Refresh()
}

func doLoadImages() {
	for load := range loads {
		doLoadImage(load.uri, load.img)
	}
}

func loadImage(u fyne.URI) fyne.CanvasObject {
	img := canvas.NewImageFromResource(nil)
	img.FillMode = canvas.ImageFillContain

	loads <- bgImageLoad{u, img}
	return img
}

func makeImageItem(u fyne.URI) fyne.CanvasObject {
	label := canvas.NewText(u.Name(), color.Gray{128})
	label.Alignment = fyne.TextAlignCenter
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, label, nil, nil),
		label, loadImage(u))
}
