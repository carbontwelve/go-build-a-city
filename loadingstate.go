package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type LoadingState struct{
	TextureManager *TextureManager
	BaseState
}

func (lS LoadingState) draw(dt float64, win *pixelgl.Window) {
	win.Clear(colornames.Whitesmoke)
	tree := pixel.NewSprite(lS.g.TextureManager.GetRef("trees"), pixel.R(0, 0, 32, 32))
	tree.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

}

func (lS LoadingState) update(dt float64, win *pixelgl.Window) {
// ...
}

func (lS LoadingState) handleInput(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		lS.g.ChangeState(NewEditorState(lS.g))
	}
}

func NewLoadingState(g *Game) *LoadingState {
	s := LoadingState{NewTextureManager(), BaseState{g}}
	return &s
}