package main

import (
	//"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type LoadingState struct{
	TextureManager *TextureManager

	BaseState

	tile *Tile
}

func (lS LoadingState) draw(dt float64, win *pixelgl.Window) {
	win.Clear(colornames.Whitesmoke)
	lS.tile.Draw(dt, win)
}

func (lS LoadingState) update(dt float64, win *pixelgl.Window) {
// ...
}

func (lS LoadingState) handleInput(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		lS.g.ChangeState(NewEditorState(lS.g))
	}

	if win.JustPressed(pixelgl.KeySpace) {
		lS.tile.tileVariant++
		if lS.tile.tileVariant > len(lS.tile.animHandler.animations) - 1 {
			lS.tile.tileVariant = 0
		}
	}
}

func NewLoadingState(g *Game) *LoadingState {
	tileAnim := make([]Animation,3)
	tileAnim[0] = *NewAnimation(0,3,0.5)
	tileAnim[1] = *NewAnimation(0,3,0.5)
	tileAnim[2] = *NewAnimation(0,3,0.5)

	s := LoadingState{
		NewTextureManager(),
		BaseState{g},
		NewTile(8, 1, g.TextureManager.GetRef("water"), tileAnim, WATER, 0,0,1)}
	return &s
}