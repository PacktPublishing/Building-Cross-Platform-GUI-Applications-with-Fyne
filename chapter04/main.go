package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/storage"
)

func chooseDirectory(w fyne.Window) {
	dialog.ShowFileOpen(func(dir fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		w.SetContent(makeUI(dir))
	}, w)
}

func startDirectory() fyne.ListableURI {
	flag.Parse()
	if len(flag.Args()) < 1 {
		cwd, _ := os.Getwd()
		return storage.NewListableURI("file://" + cwd)
	}

	dir, err := filepath.Abs(flag.Arg(0))
	if err != nil {
		log.Println("Could not find directory", dir)
		cwd, _ := os.Getwd()
		return storage.NewListableURI("file://" + cwd)
	}

	return storage.NewListableURI("file://" + dir)
}

func main() {
	a := app.New()
	w := a.NewWindow("Image Browser")

	w.SetContent(makeUI(startDirectory()))
	w.Resize(fyne.NewSize(480, 360))

	w.SetMainMenu(fyne.NewMainMenu(fyne.NewMenu("File",
		fyne.NewMenuItem("Open Directory...", func() {
			chooseDirectory(w)
		}))))

	go doLoadImages()
	w.ShowAndRun()
}
