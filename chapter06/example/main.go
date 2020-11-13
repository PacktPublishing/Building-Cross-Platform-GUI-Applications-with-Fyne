package main

import (
	"log"
	"strconv"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/data/binding"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func dateKey(t time.Time) string {
	return t.Format("2006-01-02") // ISO 8601 date, YYYY-MM-DD
}

func historyLabel() fyne.CanvasObject {
	num := widget.NewLabel("0ml")
	num.Alignment = fyne.TextAlignTrailing
	return num
}

func makeUI(p fyne.Preferences) fyne.CanvasObject {
	total := binding.BindPreferenceInt(dateKey(time.Now()), p)

	label := canvas.NewText("0ml", theme.PrimaryColor())
	label.TextSize = 42
	label.Alignment = fyne.TextAlignCenter
	totalStr := binding.IntToStringWithFormat(total, "%dml")
	totalStr.AddListener(binding.NewDataItemListener(
		func(_ binding.DataItem) {
			label.Text = totalStr.Get()
			label.Refresh()
		}))

	date := widget.NewLabel(time.Now().Format("Mon Jan 2 2006"))
	date.Alignment = fyne.TextAlignCenter

	amount := widget.NewEntry()
	amount.SetText("250")
	input := container.NewBorder(nil, nil, nil, widget.NewLabel("ml"), amount)
	add := widget.NewButton("Add", func() {
		inc, err := strconv.Atoi(amount.Text)
		if err != nil {
			log.Println("Failed to parse integer: " + amount.Text)
			return
		}

		total.Set(total.Get() + inc)
	})

	history := container.NewGridWithColumns(2,
		widget.NewLabel("Monday"), historyLabel(),
		widget.NewLabel("Tuesday"), historyLabel(),
		widget.NewLabel("Wednesday"), historyLabel(),
		widget.NewLabel("Thursday"), historyLabel(),
		widget.NewLabel("Friday"), historyLabel(),
		widget.NewLabel("Saturday"), historyLabel(),
		widget.NewLabel("Sunday"), historyLabel(),
	)

	return container.NewVBox(label, date,
		container.NewGridWithColumns(2, input, add),
		widget.NewCard("History", "Totals this week", history))
}

func main() {
	a := app.NewWithID("com.example.watertracker")
	w := a.NewWindow("Water Tracker")

	pref := a.Preferences()
	w.SetContent(makeUI(pref))
	w.ShowAndRun()
}
