package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myTheme struct {
}

func (t *myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	if n == theme.ColorNameForeground {
		return &color.NRGBA{0xff, 0xc1, 0x07, 0xff}
	}

	return theme.DefaultTheme().Color(n, v)
}

func (t *myTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}

func (t *myTheme) Font(fyne.TextStyle) fyne.Resource {
	return theme.DefaultTextMonospaceFont()
}

func (t *myTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
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
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Theme")

	w.SetContent(makeUI())
	w.ShowAndRun()
}
