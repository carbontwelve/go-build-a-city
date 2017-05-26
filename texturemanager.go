package main

import(
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
)

type TextureManager struct {
	textures map[string]pixel.Picture
}

func (tM *TextureManager) LoadTexture(name string, path string) (error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}
	tM.textures[name] = pixel.PictureDataFromImage(img)
	return nil
}

func (tM TextureManager) GetRef(texture string) pixel.Picture {
	return tM.textures[texture]
}

func NewTextureManager() *TextureManager {
	tM := TextureManager{make(map[string]pixel.Picture)}
	return &tM
}
