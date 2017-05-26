package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/carbontwelve/go-build-a-city/gobuildacity"
)

func main() {
	game := gobuildacity.Game{
		States: gobuildacity.NewStack(),
		UserQuits: false,
		TextureManager: gobuildacity.NewTextureManager(),
	}

	game.LoadTextures()
	game.PushState(gobuildacity.NewLoadingState(&game))
	pixelgl.Run(game.GameLoop)
}