package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/carbontwelve/go-build-a-city/gobuildacity"
)

func main() {
	game := Game{
		states: NewStack(),
		userQuits: false,
		TextureManager: NewTextureManager(),
	}

	game.LoadTextures()
	game.PushState(NewLoadingState(&game))
	pixelgl.Run(game.GameLoop)
}