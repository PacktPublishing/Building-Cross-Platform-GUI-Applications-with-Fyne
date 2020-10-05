package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

type greeter struct {
	greeting       *widget.Label
	name           *widget.Entry
	updateGreeting *widget.Button
}

func (g *greeter) setGreeting() {
	text := fmt.Sprintf("Hello %s!", g.name.Text)
	g.greeting.SetText(text)
}

func (g *greeter) makeUI() fyne.CanvasObject {
	g.greeting = widget.NewLabel("Hello World!")
	g.name = widget.NewEntry()
	g.name.PlaceHolder = "Enter name"
	g.updateGreeting = widget.NewButton("Welcome", g.setGreeting)
	return container.NewVBox(g.greeting, g.name, g.updateGreeting)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello!")

	g := newGreeter()
	w.SetContent(g.makeUI())
	w.ShowAndRun()
}
