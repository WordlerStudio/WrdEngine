package WrdEngine

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func ImagePathToTexture(renderer *sdl.Renderer, imgPath string) (*sdl.Texture, error) {
	// Wczytaj obraz z pliku
	image, err := img.Load(imgPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load image: %w", err)
	}
	defer image.Free()

	// Konwertuj obraz na teksturę
	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		return nil, fmt.Errorf("failed to create texture: %w", err)
	}

	return texture, nil
}

func ImageToTexture(renderer *sdl.Renderer, imgData []byte) (*sdl.Texture, error) {
	tempRWops, err := sdl.RWFromMem(imgData)
	if err != nil {
		return nil, ImageLoadError(err)
	}

	surface, err := img.LoadRW(tempRWops, false)
	if err != nil {
		return nil, ImageLoadError(err)
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, ImageLoadError(err)
	}

	return texture, nil
}
