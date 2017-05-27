package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type TileType int

const (
	VOID        TileType = 1 + iota
	GRASS
	FOREST
	WATER
	RESIDENTIAL
	COMMERCIAL
	INDUSTRIAL
	ROAD
)

func IntToTileType(i int) TileType {
	switch i {
	case 1:
		return VOID
	case 2:
		return GRASS
	case 3:
		return FOREST
	case 4:
		return WATER
	case 5:
		return RESIDENTIAL
	case 6:
		return COMMERCIAL
	case 7:
		return INDUSTRIAL
	}
	return VOID
}

func TileTypeToStr(t TileType) string {
	switch t {
	case VOID:
		return "Void"
	case GRASS:
		return "Grass"
	case FOREST:
		return "Forest"
	case WATER:
		return "Water"
	case RESIDENTIAL:
		return "Residential Zone"
	case COMMERCIAL:
		return "Commercial Zone"
	case INDUSTRIAL:
		return "Industrial Zone"
	}
	return "Void"
}

type Tile struct {
	animHandler *AnimationHandler

	sprite *pixel.Sprite

	position pixel.Vec

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
	t.sprite.Draw(win, pixel.IM.Moved(t.position))
}

func (t *Tile) Update() {
	if (t.tileType == RESIDENTIAL || t.tileType == COMMERCIAL || t.tileType == INDUSTRIAL) && uint(t.population) == t.maxPopPerLevel*uint(t.tileVariant+1) && uint(t.tileVariant) < t.maxLevels {
		// if(rand() % int(1e4) < 1e2 / (this->tileVariant+1)) ++this->tileVariant;
		// Essentially the chance is 10% for tileVariant = 0, 5% for tileVariant = 1, 3.33% for tileVariant = 2, and so on.
		// @todo implement the above in golang
	}
}

func (t Tile) GetCost() uint {
	return t.cost
}

func NewTile(
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

	t.animHandler = NewAnimationHandler(pixel.R(0, 0, float64(tileSize*2), float64(tileSize*height)))
	t.sprite = pixel.NewSprite(texture, t.animHandler.frameSize)

	// @todo not implemented
	// because tiles are one half the height of the full tile we need to change what is regarded as (0,0) on the sprite
	// so that the correct thing is shown
	//this->sprite.setOrigin(sf::Vector2f(0.0f, tileSize*(height-1)));

	for _, animation := range (animations) {
		t.animHandler.AddAnimation(animation)
	}

	t.animHandler.Update(0.0)

	return &t
}
