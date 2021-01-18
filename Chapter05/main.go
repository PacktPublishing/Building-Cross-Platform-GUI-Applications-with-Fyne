package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type taskApp struct {
	data    *taskList
	visible []*task
	current *task

	tasks                   *widget.List
	title, description, due *widget.Entry
	category                *widget.Select
	priority                *widget.RadioGroup
	completion              *widget.Slider
}

func (a *taskApp) makeUI() fyne.CanvasObject {
	a.tasks = widget.NewList(
		func() int {
			return len(a.visible)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewCheck("",
				func(bool) {}),
				widget.NewLabel("TODO Item x"))
		},
		func(i int, c fyne.CanvasObject) {
			task := a.visible[i]
			box := c.(*fyne.Container)
			check := box.Objects[0].(*widget.Check)
			check.Checked = task.done
			check.OnChanged = func(done bool) {
				task.done = done
				a.refreshData()
			}
			label := box.Objects[1].(*widget.Label)
			label.SetText(task.title)
		})
	a.tasks.OnSelected = func(id int) {
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
	a.priority = widget.NewRadioGroup([]string{"Low", "Mid", "High"}, func(pri string) {
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
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			task := &task{title: "New task"}
			a.data.add(task)
			a.setTask(task)
			a.refreshData()
		}),
	)
	return container.NewBorder(toolbar, nil, a.tasks, nil, details)
}

func (a *taskApp) refreshData() {
	// hide done
	a.visible = a.data.remaining()
	a.tasks.Refresh()
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
	tasks := &taskApp{data: data, visible: data.remaining()}
	w.SetContent(tasks.makeUI())
	if len(data.remaining()) > 0 {
		tasks.setTask(data.remaining()[0])
	}
	w.ShowAndRun()
}
