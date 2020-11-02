package main

import (
	"image/color"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

type snakePart struct {
	x, y int
}

type moveType int

const (
	moveUp moveType = iota
	moveDown
	moveLeft
	moveRight
)

var (
	snakeParts []snakePart
	game       *fyne.Container
	move       = moveUp
)

func keyTyped(e *fyne.KeyEvent) {
	switch e.Name {
	case fyne.KeyUp:
		move = moveUp
	case fyne.KeyDown:
		move = moveDown
	case fyne.KeyLeft:
		move = moveLeft
	case fyne.KeyRight:
		move = moveRight
	}
}

func main() {
	a := app.New()

	w := a.NewWindow("Snake")
	w.Resize(fyne.NewSize(200, 200))
	w.SetFixedSize(true)

	game = setupGame()
	w.SetContent(game)
	w.Canvas().SetOnTypedKey(keyTyped)

	go runGame()
	w.ShowAndRun()
}

func refreshGame() {
	for i, seg := range snakeParts {
		rect := game.Objects[i]
		rect.Move(fyne.NewPos(seg.x*10, seg.y*10))
	}

	game.Refresh()
}

func runGame() {
	for {
		time.Sleep(time.Millisecond * 250)
		for i := len(snakeParts) - 1; i >= 1; i-- {
			snakeParts[i] = snakeParts[i-1]
		}
		switch move {
		case moveUp:
			snakeParts[0].y--
		case moveDown:
			snakeParts[0].y++
		case moveLeft:
			snakeParts[0].x--
		case moveRight:
			snakeParts[0].x++
		}
		refreshGame()
	}
}

func setupGame() *fyne.Container {
	var segments []fyne.CanvasObject
	for i := 0; i < 10; i++ {
		seg := snakePart{9, 5 + i}
		snakeParts = append(snakeParts, seg)

		r := canvas.NewRectangle(&color.RGBA{G: 0x66, A: 0xff})
		r.Resize(fyne.NewSize(10, 10))
		r.Move(fyne.NewPos(90, 50+(i*10)))
		segments = append(segments, r)
	}

	return fyne.NewContainerWithoutLayout(segments...)
}
