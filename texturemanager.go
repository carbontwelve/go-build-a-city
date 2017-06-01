package main

import (
	"image"
	_ "image/png"
	"image/color"
	"os"

	"github.com/faiface/pixel"
	"image/draw"
	"fmt"
)

// Create a struct to deal with pixel
type Pixel struct {
	Point image.Point
	Color color.Color
}

type TextureManager struct {
	textures    map[string]pixel.Picture
	spriteSheet image.Image
}

func (tM TextureManager) decodePixelsFromImage(img image.Image, offsetX, offsetY int) []*Pixel {
	pixels := []*Pixel{}
	for y := 0; y <= img.Bounds().Max.Y; y++ {
		for x := 0; x <= img.Bounds().Max.X; x++ {
			p := &Pixel{
				Point: image.Point{x + offsetX, y + offsetY},
				Color: img.At(x, y),
			}
			pixels = append(pixels, p)
		}
	}
	return pixels
}

func (tM *TextureManager) LoadTexture(name string, relativePath string) (error) {
	file, err := os.Open(relativePath)
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	if tM.spriteSheet == nil {
		tM.bootSpriteSheet(img)
	} else {
		tM.appendToSpriteSheet(img, name)
	}

	// @todo append spritesheet with texture

	tM.textures[name] = pixel.PictureDataFromImage(img)
	return nil
}

func (tM *TextureManager) bootSpriteSheet(img image.Image) {
	tM.spriteSheet = img
}

// https://stackoverflow.com/questions/35964656/golang-how-to-concatenate-append-images-to-one-another
func (tM *TextureManager) appendToSpriteSheet(img image.Image, name string) {
	// collect pixel data from each image
	pixels1 := tM.decodePixelsFromImage(tM.spriteSheet, 0, 0)
	// the second image has a Y-offset of tM.spriteSheet's max Y (appended at bottom)
	pixels2 := tM.decodePixelsFromImage(img, 0, tM.spriteSheet.Bounds().Max.Y)

	pixelSum := append(pixels1, pixels2...)

	// Set a new size for the new image equal to the max width
	// of bigger image and max height of two images combined

	// Identify bigger width
	w := tM.spriteSheet.Bounds().Dx()
	h := tM.spriteSheet.Bounds().Dy() + img.Bounds().Dy()

	if img.Bounds().Dx() > w {
		w = img.Bounds().Dx()
	}

	newRect := image.Rectangle{
		Min: image.Point{0,0},
		Max: image.Point{w,h},
	}

	//newRect := image.Rectangle{
	//	Min: tM.spriteSheet.Bounds().Min,
	//	Max: image.Point{
	//		X: img.Bounds().Max.X,
	//		Y: img.Bounds().Max.Y + tM.spriteSheet.Bounds().Max.Y,
	//	},
	//}
	finImage := image.NewRGBA(newRect)
	// This is the cool part, all you have to do is loop through
	// each Pixel and set the image's color on the go
	for _, px := range pixelSum {
		finImage.Set(
			px.Point.X,
			px.Point.Y,
			px.Color,
		)
	}

	draw.Draw(finImage, finImage.Bounds(), finImage, image.Point{0,0}, draw.Src)

	fmt.Println("Added [", name ,"] new spritesheet dimensions W:", finImage.Bounds().Dx(), "H:", finImage.Bounds().Dy())

	tM.spriteSheet = finImage
}

func (tM TextureManager) GetRef(texture string) pixel.Picture {
	return tM.textures[texture]
}

func (tM TextureManager) GetSpriteSheet() pixel.Picture {
	return pixel.PictureDataFromImage(tM.spriteSheet)
}

func NewTextureManager() *TextureManager {
	tM := TextureManager{make(map[string]pixel.Picture), nil}
	return &tM
}
