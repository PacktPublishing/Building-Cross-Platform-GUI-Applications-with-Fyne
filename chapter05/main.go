package main

import (
	"time"

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
	a.tasks.OnItemSelected = func(id int) {
		a.setTask(a.visible[id])
	}

	a.title = widget.NewEntry()
	a.title.OnChanged = func(text string) {
		if a.current == nil {
			return
		}

		a.current.title = text
		a.tasks.Refresh()
	}
	a.description = widget.NewMultiLineEntry()
	a.description.OnChanged = func(text string) {
		if a.current == nil {
			return
		}

		a.current.description = text
	}

	a.category = widget.NewSelect([]string{"Home"}, func(cat string) {
		if a.current == nil {
			return
		}

		a.current.category = cat
	})
	a.priority = widget.NewRadio([]string{"Low", "Mid", "High"}, func(pri string) {
		if a.current == nil {
			return
		}

		if pri == "Mid" {
			a.current.priority = midPriority
		} else if pri == "High" {
			a.current.priority = highPriority
		} else {
			a.current.priority = lowPriority
		}
	})
	a.due = widget.NewEntry()
	a.due.Validator = dateValidator
	a.due.OnChanged = func(str string) {
		if a.current == nil {
			return
		}

		if str == "" {
			a.current.due = nil
		} else {
			date, err := time.Parse(dateFormat, str)
			if err != nil {
				a.current.due = &date
			}
		}
	}

	a.completion = widget.NewSlider(0, 100)
	a.completion.OnChanged = func(done float64) {
		if a.current == nil {
			return
		}

		a.current.completion = done
	}

	details := widget.NewForm(
		widget.NewFormItem("Title", a.title),
		widget.NewFormItem("Description", a.description),
		widget.NewFormItem("Category", a.category),
		widget.NewFormItem("Priority", a.priority),
		widget.NewFormItem("Due", a.due),
		widget.NewFormItem("Completion", a.completion),
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

func (a *taskApp) setTask(t *task) {
	a.current = t

	a.title.SetText(t.title)
	a.description.SetText(t.description)
	a.category.SetSelected(t.category)
	if t.priority == midPriority {
		a.priority.SetSelected("Mid")
	} else if t.priority == highPriority {
		a.priority.SetSelected("High")
	} else {
		a.priority.SetSelected("Low")
	}
	a.due.SetText(formatDate(t.due))
	a.completion.Value = t.completion
	a.completion.Refresh()
}

func formatDate(date *time.Time) string {
	if date == nil {
		return ""
	}
	return date.Format(dateFormat)
}

func main() {
	a := app.New()
	w := a.NewWindow("Task List")

	data := dummyData()
	ui := &taskApp{data: data, visible: data.remaining()}
	w.SetContent(ui.makeUI())
	if len(data.remaining()) > 0 {
		ui.setTask(data.remaining()[0])
	}
	w.ShowAndRun()
}
