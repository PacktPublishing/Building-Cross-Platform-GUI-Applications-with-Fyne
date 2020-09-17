package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func makeUI() fyne.CanvasObject {
	todos := widget.NewList(func() int {
		return 5
	},
		func() fyne.CanvasObject {
			return widget.NewCheck("TODO Item x", func(bool) {})
		},
		func(int, fyne.CanvasObject) {})

	details := widget.NewForm(
		widget.NewFormItem("Title", widget.NewEntry()),
		widget.NewFormItem("Description", widget.NewMultiLineEntry()),
		widget.NewFormItem("Category", widget.NewSelect([]string{"Home"}, func(string) {})),
		widget.NewFormItem("Priority", widget.NewRadio([]string{"Low", "Mid", "High"}, func(string) {})),
		widget.NewFormItem("Due", widget.NewEntry()),
		widget.NewFormItem("Completion", widget.NewSlider(0, 100)),
	)

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentRemoveIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.CheckButtonCheckedIcon(), func() {}),
	)
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, todos, nil),
		toolbar, todos, details)
}

func main() {
	a := app.New()
	w := a.NewWindow("Task List")

	w.SetContent(makeUI())
	w.ShowAndRun()
}
