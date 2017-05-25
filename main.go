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
		TextureManager: NewTextureManager(),
	}

	r := game.TextureManager.LoadTexture("trees", "trees.png")
	if r != nil {
		panic(r)
	}

	r = game.TextureManager.LoadTexture("water", "water.png")
	if r != nil {
		panic(r)
	}

	game.PushState(NewLoadingState(&game))
	pixelgl.Run(game.GameLoop)
}