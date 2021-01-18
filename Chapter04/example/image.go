package main

import (
	"image"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/storage"

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

type itemLayout struct {
	bg, text, gradient fyne.CanvasObject
}

func (i *itemLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(160, 120)
}

func (i *itemLayout) Layout(objs []fyne.CanvasObject, size fyne.Size) {
	textHeight := float32(22)
	for _, o := range objs {
		if o == i.text {
			o.Move(fyne.NewPos(0, size.Height-textHeight))
			o.Resize(fyne.NewSize(size.Width, textHeight))
		} else if o == i.bg {
			o.Move(fyne.NewPos(0, size.Height-textHeight))
			o.Resize(fyne.NewSize(size.Width, textHeight))
		} else if o == i.gradient {
			o.Move(fyne.NewPos(0, size.Height-(textHeight*1.5)))
			o.Resize(fyne.NewSize(size.Width, textHeight/2))
		} else {
			o.Move(fyne.NewPos(0, 0))
			o.Resize(size)
		}
	}
}

func makeImageItem(u fyne.URI) fyne.CanvasObject {
	label := canvas.NewText(u.Name(), color.Gray{128})
	label.Alignment = fyne.TextAlignCenter

	bgColor := &color.NRGBA{R: 255, G: 255, B: 255, A: 224}
	bg := canvas.NewRectangle(bgColor)
	fade := canvas.NewLinearGradient(color.Transparent, bgColor, 0)
	return fyne.NewContainerWithLayout(&itemLayout{text: label, bg: bg, gradient: fade},
		loadImage(u), bg, fade, label)
}
