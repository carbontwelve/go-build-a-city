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
	states *stack
	userQuits bool
}

func (g *Game) PushState(state GameState) {
	g.states.Push(state)
}

func (g *Game) PopState() {
	g.states.Pop()
	fmt.Printf("%+v\n", g.states)
}

func (g *Game) ChangeState(state GameState) {
	if g.states.Len() > 0 {
		g.states.Pop()
	}
	g.PushState(state)
}

func (g *Game) PeekState() GameState {
	if g.states.Len() == 0 {
		return nil
	}
	return g.states.Peek()
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

	for !win.Closed() && g.userQuits == false {
		dt := time.Since(clock).Seconds()
		clock = time.Now()

		if g.PeekState() == nil {
			break
		}

		if win.JustPressed(pixelgl.KeyEscape) {
			g.userQuits = true
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