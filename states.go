package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
)

type GameState interface {
	draw(dt float64, win *pixelgl.Window, cam pixel.Matrix)
	update(dt float64, win *pixelgl.Window)
	handleInput(win *pixelgl.Window)
	setGame(g *Game)
}

type BaseState struct {
	g *Game
}

func (s *BaseState) setGame(g *Game) {
	s.g = g
}