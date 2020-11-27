package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func loadMessages() *fyne.Container {
	messages := []fyne.CanvasObject{
		newMessage("Hi there, how are you doing?", "Jim"),
		newMessage("Yeah good thanks, you?", "Me"),
		newMessage("Not bad thanks. Weekend!", "Jim"),
		newMessage("Want to visit the cinema?", "Jim"),
		newMessage("Great idea, what's showing?", "Me"),
	}

	return container.NewVBox(messages...)
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
	//	a.Settings().SetTheme(theme.ExtendDefaultTheme(&myTheme{}))
	w := a.NewWindow("Messages")

	w.SetContent(makeUI())
	w.Resize(fyne.NewSize(160, 280))
	w.ShowAndRun()
}
