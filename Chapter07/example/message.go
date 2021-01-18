package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	myName        = "Me"
	messageIndent = 20
)

type message struct {
	widget.BaseWidget

	text, from string
}

func newMessage(text, name string) *message {
	m := &message{text: text, from: name}
	m.ExtendBaseWidget(m)
	return m
}

func (m *message) CreateRenderer() fyne.WidgetRenderer {
	text := widget.NewLabel(m.text)
	text.Wrapping = fyne.TextWrapWord
	return &messageRender{msg: m, bg: &canvas.Rectangle{},
		txt: text}
}

type messageRender struct {
	msg *message

	bg  *canvas.Rectangle
	txt *widget.Label
}

func (r *messageRender) messageMinSize(s fyne.Size) fyne.Size {
	fitSize := s.Subtract(fyne.NewSize(messageIndent, 0))
	r.txt.Resize(fitSize.Max(r.txt.MinSize())) // have the wrap code run
	return r.txt.MinSize()
}

func (r *messageRender) Layout(s fyne.Size) {
	itemSize := r.messageMinSize(s)
	itemSize = itemSize.Max(fyne.NewSize(s.Width-messageIndent, s.Height))

	bgPos := fyne.NewPos(0, 0)
	if r.msg.from == myName {
		r.txt.Alignment = fyne.TextAlignTrailing
		r.bg.FillColor = theme.PrimaryColorNamed(theme.ColorBlue)
		bgPos = fyne.NewPos(s.Width-itemSize.Width, 0)
	} else {
		r.txt.Alignment = fyne.TextAlignLeading
		r.bg.FillColor = theme.PrimaryColorNamed(theme.ColorGreen)
	}

	r.txt.Move(bgPos)
	r.bg.Resize(itemSize)
	r.bg.Move(bgPos)
}

func (r *messageRender) MinSize() fyne.Size {
	itemSize := r.messageMinSize(r.msg.Size())
	return itemSize.Add(fyne.NewSize(messageIndent, 0))
}

func (r *messageRender) Refresh() {
}

func (r *messageRender) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *messageRender) Destroy() {
}

func (r *messageRender) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.bg, r.txt}
}
