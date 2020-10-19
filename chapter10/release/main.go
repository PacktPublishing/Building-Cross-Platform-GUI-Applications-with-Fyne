package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
)

const (
	serverKeyDevelopment = "DEVELOPMENT_KEY"
	serverKeyProduction  = "PRODUCTION_KEY"
)

func connectToServer(a fyne.App, w fyne.Window) {
	key := serverKeyProduction
	if a.BuildType() == fyne.ReleaseBuild {
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
