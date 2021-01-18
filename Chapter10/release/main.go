package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

const (
	serverKeyDevelopment = "DEVELOPMENT_KEY"
	serverKeyProduction  = "PRODUCTION_KEY"
)

func connectToServer(a fyne.App, w fyne.Window) {
	key := serverKeyDevelopment
	if a.Settings().BuildType() == fyne.BuildRelease {
		key = serverKeyProduction
	}

	dialog.ShowInformation("Connect to server", "Using key: "+key, w)
}

func main() {
	a := app.New()
	w := a.NewWindow("Server key demo")

	w.SetContent(widget.NewLabel("Connecting..."))
	w.Resize(fyne.NewSize(300, 160))
	connectToServer(a, w)
	w.ShowAndRun()
}
