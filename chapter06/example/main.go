package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func historyLabel() fyne.CanvasObject {
	num := widget.NewLabel("0ml")
	num.Alignment = fyne.TextAlignTrailing
	return num
}

func makeUI() fyne.CanvasObject {
	label := canvas.NewText("0ml", theme.PrimaryColor())
	label.TextSize = 42
	label.Alignment = fyne.TextAlignCenter

	date := widget.NewLabel("Mon 9 Nov 2020")
	date.Alignment = fyne.TextAlignCenter

	amount := widget.NewEntry()
	amount.SetText("250")
	input := container.NewBorder(nil, nil, nil, widget.NewLabel("ml"), amount)
	add := widget.NewButton("Add", func() {})

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
	a := app.New()
	w := a.NewWindow("Water Tracker")

	w.SetContent(makeUI())
	w.ShowAndRun()
}
