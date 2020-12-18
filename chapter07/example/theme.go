package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/theme"
)

type myTheme struct {
}

func (m *myTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	switch n {
	case theme.ColorNameBackground:
		if v == theme.VariantNameLight {
			return &color.NRGBA{0xcf, 0xd8, 0xdc, 0xff}
		}
		return &color.NRGBA{0x45, 0x5A, 0x64, 0xff}
	case theme.ColorNameFocus:
		return &color.NRGBA{0xff, 0xc1, 0x07, 0xff}
	}

	return theme.DefaultTheme().Color(n, v)
}

func (m *myTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}

func (m *myTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (m *myTheme) Icon(i fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(i)
}
