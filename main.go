package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	game := Game{
		States: NewStack(),
		UserQuits: false,
		TextureManager: NewTextureManager(),
	}

	game.LoadTextures()
	game.PushState(NewLoadingState(&game))
	pixelgl.Run(game.GameLoop)
}