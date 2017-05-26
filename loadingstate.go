package main

import (
	//"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"github.com/faiface/pixel/text"
	//"github.com/faiface/pixel"
	"github.com/golang/freetype/truetype"
	"unicode"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"github.com/faiface/pixel"
	"fmt"
)

func ttfFromBytesMust(b []byte, size float64) font.Face {
	ttf, err := truetype.Parse(b)
	if err != nil {
		panic(err)
	}
	return truetype.NewFace(ttf, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	})
}

type LoadingState struct{
	TextureManager *TextureManager

	debugText *text.Text

	BaseState

	tile *Tile
}

func (lS LoadingState) draw(dt float64, win *pixelgl.Window) {
	win.Clear(colornames.Whitesmoke)
	lS.tile.Draw(dt, win)

	lS.debugText.Draw(win, pixel.IM.Moved(pixel.V(20, win.Bounds().H() - lS.debugText.LineHeight - 20)))

}

func (lS LoadingState) update(dt float64, win *pixelgl.Window) {

	lS.debugText.Clear()
	lS.debugText.Dot = lS.debugText.Orig
	lS.debugText.Color = colornames.Darkslategray

	lS.debugText.WriteString(fmt.Sprintf("Tile Variant: %d\nBounds W/H: %f/%f\nMin: (%f,%f) Max: (%f,%f)", lS.tile.tileVariant, lS.tile.animHandler.bounds.W(), lS.tile.animHandler.bounds.H(), lS.tile.animHandler.bounds.Min.X, lS.tile.animHandler.bounds.Min.Y, lS.tile.animHandler.bounds.Max.X, lS.tile.animHandler.bounds.Max.Y))

}

func (lS LoadingState) handleInput(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyEnter) {
		lS.g.ChangeState(NewEditorState(lS.g))
	}

	if win.JustPressed(pixelgl.KeySpace) {
		lS.tile.tileVariant++
		if lS.tile.tileVariant > len(lS.tile.animHandler.animations) - 1 {
			lS.tile.tileVariant = 0
		}
	}
}

func NewLoadingState(g *Game) *LoadingState {
	tileAnim := make([]Animation,3)
	tileAnim[0] = *NewAnimation(0,3,0.5)
	tileAnim[1] = *NewAnimation(0,3,0.5)
	tileAnim[2] = *NewAnimation(0,3,0.5)

	regular := text.NewAtlas(ttfFromBytesMust(goregular.TTF, 12), text.ASCII, text.RangeTable(unicode.Latin))
	dTxt := text.New(pixel.ZV, regular)
	dTxt.Color = colornames.Darkslategray

	s := LoadingState{
		NewTextureManager(),
		dTxt,
		BaseState{g},
		NewTile(8, 1, g.TextureManager.GetRef("water"), tileAnim, WATER, 0,0,1)}
	return &s
}