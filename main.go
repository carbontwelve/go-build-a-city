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
	}

	redState := RedState{}
	redState.setGame(&game)
	game.PushState(&redState)

	blueState := BlueState{}
	blueState.setGame(&game)
	game.PushState(&blueState)

	//game.PushState(new(BlueState))
	// game.PushState(new(LoadingState))

	fmt.Printf("%+v\n", game.states)

	pixelgl.Run(game.GameLoop)
}