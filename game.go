package main

import (
	"github.com/golang-collections/collections/stack"
	"time"
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel"
)

type Game struct {
	states *stack.Stack
}

func (g *Game) PushState(state GameState) {
	g.states.Push(state)
}

func (g *Game) PopState() {
	g.states.Pop()
}

func (g *Game) ChangeState(state GameState) {
	if (g.states.Len() > 0) {
		g.states.Pop()
	}
	g.PushState(state)
}

func (g *Game) PeekState() GameState {
	if (g.states.Len() == 0) {
		return nil
	}
	return g.states.Peek().(GameState)
}

func (g *Game) GameLoop() {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	//clock := time.Now()

	var (
		frames = 0
		second = time.Tick(time.Second)
	)

	for !win.Closed() {
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

	//for {
	//	//dt := time.Since(clock).Seconds()
	//	//clock = time.Now()
	//
	//	if (g.PeekState() == nil) {
	//		continue
	//	}
	//
	//	g.PeekState().handleInput()
	//	g.PeekState().update(1.1)
	//	g.PeekState().draw(1.1)
	//
	//	// FPS Counter
	//	frames++
	//	select {
	//	case <-second:
	//		fmt.Printf("FPS: %d\n", frames)
	//		frames = 0
	//	default:
	//	}
	//}
}