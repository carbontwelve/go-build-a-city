package gobuildacity

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"fmt"
)

type RedState struct{
	BaseState
}

func (lS RedState) draw(dt float64, win *pixelgl.Window) {
	win.Clear(colornames.Firebrick)
}

func (lS RedState) update(dt float64, win *pixelgl.Window) {
	// ...
}

func (lS RedState) handleInput(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyRight) {
		fmt.Printf("%+v\n", lS.g.states)
	}

	if win.JustPressed(pixelgl.KeyLeft) {
		lS.g.ChangeState(NewBlueState(lS.g))
	}
}

func NewRedState(g *Game) *RedState {
	redState := RedState{}
	redState.setGame(g)
	return &redState
}