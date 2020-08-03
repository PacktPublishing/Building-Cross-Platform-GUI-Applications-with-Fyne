package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
)

func isImage(file fyne.URI) bool {
	ext := strings.ToLower(file.Extension())

	return ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".gif"
}

func filterImages(files []fyne.URI) []fyne.URI {
	images := []fyne.URI{}

	for _, file := range files {
		if isImage(file) {
			images = append(images, file)
		}
	}

	return images
}

func makeImageGrid(images []fyne.URI) fyne.CanvasObject {
	items := []fyne.CanvasObject{}

	for _, item := range images {
		img := makeImageItem(item)
		items = append(items, img)
	}

	cellSize := fyne.NewSize(160, 120)
	return fyne.NewContainerWithLayout(layout.NewGridWrapLayout(cellSize), items...)
}

func makeStatus(dir fyne.ListableURI, images []fyne.URI) fyne.CanvasObject {
	status := fmt.Sprintf("Directory %s, %d items", dir.Name(), len(images))
	return canvas.NewText(status, color.Gray{128})
}

func makeUI(dir fyne.ListableURI) fyne.CanvasObject {
	images := filterImages(dir.List())
	status := makeStatus(dir, images)
	content := makeImageGrid(images)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, status, nil, nil),
		status, content)
}
