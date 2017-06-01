package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

type EditorState struct{
	BaseState
	sprite *pixel.Sprite
}

func (lS EditorState) draw(dt float64, win *pixelgl.Window, cam pixel.Matrix) {
	win.Clear(colornames.Darkslategray)
	lS.sprite.Draw(win, pixel.IM.Moved(cam.Unproject(win.Bounds().Center())))
}

func (lS EditorState) update(dt float64, win *pixelgl.Window) {
// ...
}

func (lS EditorState) handleInput(win *pixelgl.Window) {

	// ...

}

func NewEditorState(g *Game) *EditorState {

	spriteSheet := g.TextureManager.GetSpriteSheet()

	s := EditorState{
		sprite: pixel.NewSprite(spriteSheet, spriteSheet.Bounds()),
	}
	s.setGame(g)
	return &s
}