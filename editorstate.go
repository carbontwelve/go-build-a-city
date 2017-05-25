package main

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type EditorState struct{
	BaseState
}

func (lS EditorState) draw(dt float64, win *pixelgl.Window) {
	win.Clear(colornames.Darkslategray)
}

func (lS EditorState) update(dt float64, win *pixelgl.Window) {
// ...
}

func (lS EditorState) handleInput(win *pixelgl.Window) {

	// ...

}

func NewEditorState(g *Game) *EditorState {
	s := EditorState{}
	s.setGame(g)
	return &s
}