package gobuildacity

import (
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"github.com/faiface/pixel"
)

type EditorState struct{
	BaseState
}

func (lS EditorState) draw(dt float64, win *pixelgl.Window) {
	win.Clear(colornames.Darkslategray)

	spritesheet := lS.g.TextureManager.GetRef("trees")

	var treesFrames []pixel.Rect
	for x := spritesheet.Bounds().Min.X; x < spritesheet.Bounds().Max.X; x += 32 {
		for y := spritesheet.Bounds().Min.Y; y < spritesheet.Bounds().Max.Y; y += 32 {
			treesFrames = append(treesFrames, pixel.R(x, y, x+32, y+32))
		}
	}

	tree := pixel.NewSprite(spritesheet, treesFrames[6])
	tree.Draw(win, pixel.IM.Scaled(pixel.ZV, 16).Moved(win.Bounds().Center()))

}

func (lS EditorState) update(dt float64, win *pixelgl.Window) {
// ...
}

func (lS EditorState) handleInput(win *pixelgl.Window) {

	// ...

}

func NewEditorState(g *Game) *EditorState {
	s := EditorState{}
	s.setGame(g)
	return &s
}