package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type myTheme struct {
}

func (t *myTheme) Color(n fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	if n == theme.Colors.Text {
		return &color.NRGBA{0xff, 0xc1, 0x07, 0xff}
	}

	return nil
}

func (t *myTheme) Size(fyne.ThemeSizeName) int {
	return 0
}

func (t *myTheme) Font(fyne.TextStyle) fyne.Resource {
	return theme.DefaultTextMonospaceFont()
}

func makeUI() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("Hello there"),
		widget.NewEntry(),
		widget.NewButton("Tap me", func() {}),
	)
}

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.ExtendDefaultTheme(&myTheme{}))
	w := a.NewWindow("Theme")

	w.SetContent(makeUI())
	w.ShowAndRun()
}
