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
	case theme.Colors.Background:
		if v == theme.Variants.Light {
			return &color.NRGBA{0xcf, 0xd8, 0xdc, 0xff}
		}
		return &color.NRGBA{0x45, 0x5A, 0x64, 0xff}
	case theme.Colors.Focus:
		return &color.NRGBA{0xff, 0xc1, 0x07, 0xff}
	}

	return nil
}

func (m *myTheme) Size(fyne.ThemeSizeName) int {
	return 0
}

func (m *myTheme) Font(fyne.TextStyle) fyne.Resource {
	return nil
}
