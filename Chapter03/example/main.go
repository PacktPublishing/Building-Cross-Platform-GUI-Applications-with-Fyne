package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

type snakePart struct {
	x, y float32
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
	head       *canvas.Rectangle
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
	nextPart := snakePart{snakeParts[0].x, snakeParts[0].y - 1}
	for {
		oldPos := fyne.NewPos(snakeParts[0].x*10, snakeParts[0].y*10)
		newPos := fyne.NewPos(nextPart.x*10, nextPart.y*10)
		canvas.NewPositionAnimation(oldPos, newPos, time.Millisecond*250, func(p fyne.Position) {
			head.Move(p)
			canvas.Refresh(head)
		}).Start()

		end := len(snakeParts) - 1
		canvas.NewPositionAnimation(fyne.NewPos(snakeParts[end].x*10, snakeParts[end].y*10),
			fyne.NewPos(snakeParts[end-1].x*10, snakeParts[end-1].y*10), time.Millisecond*250,
			func(p fyne.Position) {
				tail := game.Objects[end]
				tail.Move(p)
				canvas.Refresh(tail)
			}).Start()

		time.Sleep(time.Millisecond * 250)
		for i := len(snakeParts) - 1; i >= 1; i-- {
			snakeParts[i] = snakeParts[i-1]
		}
		snakeParts[0] = nextPart
		refreshGame()
		switch move {
		case moveUp:
			nextPart = snakePart{nextPart.x, nextPart.y - 1}
		case moveDown:
			nextPart = snakePart{nextPart.x, nextPart.y + 1}
		case moveLeft:
			nextPart = snakePart{nextPart.x - 1, nextPart.y}
		case moveRight:
			nextPart = snakePart{nextPart.x + 1, nextPart.y}
		}
	}
}

func setupGame() *fyne.Container {
	var segments []fyne.CanvasObject
	for i := 0; i < 10; i++ {
		seg := snakePart{9, float32(5 + i)}
		snakeParts = append(snakeParts, seg)

		r := canvas.NewRectangle(&color.RGBA{G: 0x66, A: 0xff})
		r.Resize(fyne.NewSize(10, 10))
		r.Move(fyne.NewPos(90, float32(50+i*10)))
		segments = append(segments, r)
	}

	head = canvas.NewRectangle(&color.RGBA{G: 0x66, A: 0xff})
	head.Resize(fyne.NewSize(10, 10))
	head.Move(fyne.NewPos(snakeParts[0].x*10, snakeParts[0].y*10))
	segments = append(segments, head)
	return fyne.NewContainerWithoutLayout(segments...)
}
