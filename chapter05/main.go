package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type taskApp struct {
	data    *taskList
	visible []*task
	current *task

	tasks                   *widget.List
	title, description, due *widget.Entry
	category                *widget.Select
	priority                *widget.Radio
	completion              *widget.Slider
}

func (a *taskApp) makeUI() fyne.CanvasObject {
	a.tasks = widget.NewList(
		func() int {
			return len(a.visible)
		},
		func() fyne.CanvasObject {
			return widget.NewCheck("TODO Item x", func(bool) {})
		},
		func(i int, c fyne.CanvasObject) {
			check := c.(*widget.Check)
			check.Text = a.visible[i].title
			check.Refresh()
		})

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
	return fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, a.tasks, nil),
		toolbar, a.tasks, details)
}

func main() {
	a := app.New()
	w := a.NewWindow("Task List")

	data := dummyData()
	ui := &taskApp{data: data, visible: data.remaining()}
	w.SetContent(ui.makeUI())
	w.ShowAndRun()
}
