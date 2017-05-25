package main

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type LoadingState struct{
	BaseState
}

func (lS LoadingState) draw(dt float64, win *pixelgl.Window) {
	win.Clear(colornames.Whitesmoke)
}

func (lS LoadingState) update(dt float64, win *pixelgl.Window) {
// ...
}

func (lS LoadingState) handleInput(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyLeft) {
		lS.g.PopState()
	}
}