package main

import (
	//"github.com/golang-collections/collections/stack"
	"github.com/faiface/pixel/pixelgl"
	"fmt"
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

	game.PushState(NewRedState(&game))

	fmt.Printf("%+v\n", game.states)

	pixelgl.Run(game.GameLoop)
}