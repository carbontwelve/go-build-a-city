package main

import (
	"github.com/golang-collections/collections/stack"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type GameState interface {
	draw(dt float64)
	update(dt float64)
	handleInput()
}

type LoadingState struct{}

func (lS LoadingState) draw(dt float64) {
	// ...
}

func (lS LoadingState) update(dt float64) {
	// ...
}

func (lS LoadingState) handleInput() {
	// ...
}

func main() {
	game := Game{
		states: stack.New(),
	}

	game.PushState(new(LoadingState))
	game.PushState(new(LoadingState))
	game.PushState(new(LoadingState))

	pixelgl.Run(game.GameLoop)
}