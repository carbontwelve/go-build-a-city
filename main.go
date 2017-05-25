package main

import (
	"github.com/faiface/pixel/pixelgl"
)


type GameState interface {
	draw(dt float64, win *pixelgl.Window)
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

func main() {
	game := Game{
		states: NewStack(),
		userQuits: false,
	}

	game.PushState(NewLoadingState(&game))
	pixelgl.Run(game.GameLoop)
}