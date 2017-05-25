package main

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type BlueState struct{
	BaseState
}

func (lS BlueState) draw(dt float64, win *pixelgl.Window) {
	win.Clear(colornames.Blueviolet)
}

func (lS BlueState) update(dt float64, win *pixelgl.Window) {
	// ...
}

func (lS BlueState) handleInput(win *pixelgl.Window) {
	if win.Pressed(pixelgl.KeyLeft) {
		lS.g.PopState()
	}
}
