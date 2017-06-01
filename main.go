package main

import (
	"github.com/faiface/pixel/pixelgl"
)

const tileSize uint  = 8

func main() {
	game := Game{
		States: NewStack(),
		UserQuits: false,
		TextureManager: NewTextureManager(),
		TileAtlas: make(map[string]*Tile),
	}

	game.LoadTextures()
	game.LoadTiles()
	//game.PushState(NewLoadingState(&game))
	game.PushState(NewEditorState(&game))
	pixelgl.Run(game.GameLoop)
}