package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/storage"
)

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
	w.ShowAndRun()
}
