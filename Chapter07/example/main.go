package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func loadMessages() *fyne.Container {
	return container.NewVBox(
		newMessage("Hi there, how are you doing?", "Jim"),
		newMessage("Yeah good thanks, you?", myName),
		newMessage("Not bad thanks. Weekend!", "Jim"),
		newMessage("Want to visit the cinema?", "Jim"),
		newMessage("Great idea, what's showing?", myName),
	)
}

func makeUI() fyne.CanvasObject {
	list := loadMessages()
	msg := widget.NewEntry()
	send := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		list.Add(newMessage(msg.Text, myName))
		msg.SetText("")
	})
	input := container.NewBorder(nil, nil, nil, send, msg)
	return container.NewBorder(nil, input, nil, nil,
		container.NewVScroll(list))
}

func main() {
	a := app.New()
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("Messages")

	w.SetContent(makeUI())
	w.Resize(fyne.NewSize(160, 280))
	w.ShowAndRun()
}
