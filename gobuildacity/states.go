package gobuildacity

import "github.com/faiface/pixel/pixelgl"

type GameState interface {
	draw(dt float64, win *pixelgl.Window)
	update(dt float64, win *pixelgl.Window)
	handleInput(win *pixelgl.Window)
	setGame(g *Game)
}

type BaseState struct {
	g *Game
}

func (s *BaseState) setGame(g *Game) {
	s.g = g
}