package main

import (
	//"github.com/golang-collections/collections/stack"
	"time"
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
)

type Game struct {
	States *stack
	TextureManager *TextureManager
	UserQuits bool
	TileAtlas map[string]*Tile
}

func (g *Game) PushState(state GameState) {
	g.States.Push(state)
}

func (g *Game) PopState() {
	g.States.Pop()
	fmt.Printf("%+v\n", g.States)
}

func (g *Game) ChangeState(state GameState) {
	if g.States.Len() > 0 {
		g.States.Pop()
	}
	g.PushState(state)
}

func (g *Game) PeekState() GameState {
	if g.States.Len() == 0 {
		return nil
	}
	return g.States.Peek().(GameState)
}

func (g *Game) LoadTextures() {

	var r error

	r = g.TextureManager.LoadTexture("trees", "assets/trees.png")
	if r != nil {
		panic(r)
	}

	r = g.TextureManager.LoadTexture("grass", "assets/grass.png")
	if r != nil {
		panic(r)
	}

	r = g.TextureManager.LoadTexture("forest", "assets/forest.png")
	if r != nil {
		panic(r)
	}

	r = g.TextureManager.LoadTexture("water", "assets/water.png")
	if r != nil {
		panic(r)
	}

	r = g.TextureManager.LoadTexture("residential", "assets/residential.png")
	if r != nil {
		panic(r)
	}

	r = g.TextureManager.LoadTexture("commercial", "assets/commercial.png")
	if r != nil {
		panic(r)
	}

	r = g.TextureManager.LoadTexture("industrial", "assets/industrial.png")
	if r != nil {
		panic(r)
	}

	r = g.TextureManager.LoadTexture("road", "assets/road.png")
	if r != nil {
		panic(r)
	}
}

func (g *Game) LoadTiles() {
	staticAnim := Animation{0, 0, 1.0}

	g.TileAtlas["grass"] = NewTile(tileSize, 1, g.TextureManager.GetRef("grass"), []Animation{staticAnim}, GRASS, 50, 0 ,1)
	g.TileAtlas["forest"] = NewTile(tileSize, 1, g.TextureManager.GetRef("forest"), []Animation{staticAnim}, FOREST, 100, 0 ,1)

	g.TileAtlas["water"] = NewTile(tileSize, 1, g.TextureManager.GetRef("water"), []Animation{
		*NewAnimation(0,3,0.5),
		*NewAnimation(0,3,0.5),
		*NewAnimation(0,3,0.5),
	}, WATER, 100, 0 ,1)

	g.TileAtlas["residential"] = NewTile(tileSize, 2, g.TextureManager.GetRef("residential"), []Animation{
		staticAnim, staticAnim, staticAnim,
		staticAnim, staticAnim, staticAnim,
	}, RESIDENTIAL, 300, 50, 6)

	g.TileAtlas["commercial"] = NewTile(tileSize, 2, g.TextureManager.GetRef("commercial"), []Animation{
		staticAnim, staticAnim, staticAnim, staticAnim,
	}, COMMERCIAL, 300, 50, 4)

	g.TileAtlas["industrial"] = NewTile(tileSize, 2, g.TextureManager.GetRef("industrial"), []Animation{
		staticAnim, staticAnim, staticAnim, staticAnim,
	}, INDUSTRIAL, 300, 50, 4)

	g.TileAtlas["road"] = NewTile(tileSize, 1, g.TextureManager.GetRef("road"), []Animation{
		staticAnim, staticAnim, staticAnim,
		staticAnim, staticAnim, staticAnim,
		staticAnim, staticAnim, staticAnim,
		staticAnim, staticAnim,
	}, ROAD, 100, 0, 1)
}

func (g *Game) GameLoop() {

	cfg := pixelgl.WindowConfig{
		Title:  "Go Build a City",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	clock := time.Now()

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() && g.UserQuits == false {
		dt := time.Since(clock).Seconds()
		clock = time.Now()

		if g.PeekState() == nil {
			break
		}

		if win.JustPressed(pixelgl.KeyEscape) {
			g.UserQuits = true
		}

		g.PeekState().handleInput(win)
		g.PeekState().update(dt, win)
		win.Clear(colornames.Black)
		g.PeekState().draw(dt, win)
		win.Update()

		// FPS Counter
		frames++
		select {
		case <-second:
			fmt.Printf("FPS: %d\n", frames)
			frames = 0
		default:
		}
	}

	// Do some shutdown stuff
	fmt.Println("Shutdown")
}