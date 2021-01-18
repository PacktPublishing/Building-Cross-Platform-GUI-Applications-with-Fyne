package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func makeTabs() *container.AppTabs {
	return container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Tab 1")),
		container.NewTabItem("JustText", widget.NewLabel("Tab 2")),
	)
}

func main() {
	a := app.New()
	w := a.NewWindow("Tabs hints")

	w.SetContent(makeTabs())
	w.ShowAndRun()
}
