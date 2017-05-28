package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"fmt"
)

type Map struct {
	width  uint
	height uint
	tiles  []*Tile

	// Resource Map
	resources []int

	tileSize    uint
	numSelected int
	numRegions  [1]int
}

func NewMap(fileName string, width uint, height uint, tileAtlas map[string]*Tile) *Map {
	m := Map{width: width, height: height, tileSize: tileSize}
	m.Generate(width, height, tileAtlas)
	return &m
}

func (m Map) depthFirstSearch(whitelist *[]TileType, pos pixel.Vec, label int, t int) {
	// ...
}

func (m *Map) Generate(width uint, height uint, tileAtlas map[string]*Tile) {
	m.width = width
	m.height = height
	m.tiles = make([]*Tile, width*height)

	fmt.Println("Width:", width, "Height:", height, "TileSize: ", m.tileSize)

	for y := uint(0); y < height; y++ {
		for x := uint(0); x < width; x++ {
			m.resources = append(m.resources, 255)
			var t *Tile
			t = &Tile{}
			if y%2 ==1 {
				*t = *tileAtlas["grass"]
			} else {
				*t = *tileAtlas["water"]
			}

			t.position = pixel.V(float64((x - y) * m.tileSize + m.width * m.tileSize), float64(x + y) * float64(m.tileSize) * 0.5)
			m.tiles[m.width * x + y] = t
		}
	}
}

// Load Map From Disk
func (m *Map) Load(fileName string, width uint, height uint, tileAtlas map[string]*Tile) error {
	return nil
}

// Save Map to Disk
func (m *Map) Save(fileName string) {
	// ...
}

func (m Map) Draw(dt float64, win *pixelgl.Window) {



	for y := uint(0); y < m.height; y++ {
		for x := uint(0); x < m.width; x++ {
			//pos := pixel.V(
			//	float64((x - y) * m.tileSize + m.width * m.tileSize),
			//	float64(x + y) * float64(m.tileSize) * 0.5)
			//m.tiles[m.width * x + y].position = pos

			m.tiles[m.width * x + y].Draw(dt, win)
		}
	}
}

func (m Map) FindConnectedRegions(whitelist []TileType, t int) {
	// ...
}

func (m Map) UpdateDirection(tileType TileType) {
	// ...
}
