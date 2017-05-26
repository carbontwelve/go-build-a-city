package gobuildacity

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"fmt"
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
	if win.JustPressed(pixelgl.KeyRight) {
		fmt.Printf("%+v\n", lS.g.States)
	}

	if win.JustPressed(pixelgl.KeyLeft) {
		lS.g.ChangeState(NewLoadingState(lS.g))
	}
}

func NewBlueState(g *Game) *BlueState {
	blueState := BlueState{}
	blueState.setGame(g)
	return &blueState
}