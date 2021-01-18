package main

import (
	"testing"

	"fyne.io/fyne/v2/test"

	"github.com/stretchr/testify/assert"
)

func TestGreeter_UpdateGreeting(t *testing.T) {
	g := &greeter{}
	g.makeUI()
	assert.Equal(t, "Hello World!", g.greeting.Text)
	assert.Equal(t, "", g.name.Text)

	test.Type(g.name, "Joe")
	test.Tap(g.updateGreeting)
	assert.Equal(t, "Hello Joe!", g.greeting.Text)
}

func TestGreeter_Render(t *testing.T) {
	g := &greeter{}
	w := test.NewWindow(g.makeUI())
	test.AssertImageMatches(t, "default.png", w.Canvas().Capture())

	test.Type(g.name, "Joe")
	test.Tap(g.updateGreeting)
	test.AssertImageMatches(t, "typed_joe.png", w.Canvas().Capture())
}
