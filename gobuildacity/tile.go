package gobuildacity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type TileType int

const (
	VOID TileType = 1 + iota
	GRASS
	FOREST
	WATER
	RESIDENTIAL
	COMMERCIAL
	INDUSTRIAL
	ROAD
)

type Tile struct {
	animHandler *AnimationHandler

	sprite *pixel.Sprite

	// Tile Type
	tileType TileType

	// Tile variant, allowing for different looking versions of the same tile
	tileVariant int

	// Region IDs of the tile, tiles in the same region are connected. First is for transport
	regions []uint

	// Placement cost of the tile
	cost uint

	// Current residents / employees
	population float64

	// Maximum population per growth stage / tile variant
	maxPopPerLevel uint

	// Maximum number of building levels
	maxLevels uint

	// Production output per customer/worker per day, either monetary or goods
	production float64

	// Goods stored
	storedGoods float64
}

func (t *Tile) Draw(dt float64, win *pixelgl.Window) {
	// Change the sprite to reflect the tile variant
	t.animHandler.ChangeAnimation(t.tileVariant)

	// Update the animation
	t.animHandler.Update(dt)

	// Update the sprite
	t.sprite.Set(t.sprite.Picture(), t.animHandler.bounds)

	// Draw the tile
	t.sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 16).Moved(win.Bounds().Center()))
}

func NewTile (
	tileSize uint,
	height uint,
	texture pixel.Picture,
	animations []Animation,
	tileType TileType,
	cost uint,
	maxPopPerLevel uint,
	maxlevels uint) *Tile {

	t := Tile{}

	t.tileType = tileType
	t.tileVariant = 0

	t.regions = make([]uint, 1)
	t.regions[0] = 0

	t.cost = cost
	t.population = 0
	t.maxPopPerLevel = maxPopPerLevel
	t.maxLevels = maxlevels
	t.production = 0
	t.storedGoods = 0

	t.animHandler = NewAnimationHandler()
	t.animHandler.frameSize = pixel.R(0,0, float64(tileSize*2), float64(tileSize*height))

	t.sprite = pixel.NewSprite(texture, t.animHandler.frameSize)

	// @todo not implemented, not sure what this does atm
	//this->sprite.setOrigin(sf::Vector2f(0.0f, tileSize*(height-1)));

	for _, animation := range(animations) {
		t.animHandler.AddAnimation(animation)
	}

	t.animHandler.Update(0.0)

	return &t
}