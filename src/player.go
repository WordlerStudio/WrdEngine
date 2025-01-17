package WrdEngine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	*BaseObject
}

func NewPlayer(renderer *sdl.Renderer, imagePath string, x, y int32) (*Player, error) {
	img, err := sdl.LoadBMP(imagePath)
	if err != nil {
		return nil, fmt.Errorf("Image load error: ", err)
	}
	defer img.Free()

	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, err
	}

	return &Player{
		&BaseObject{
			Texture:  texture,
			X:        x,
			Y:        y,
			Width:    img.W,
			Height:   img.H,
			Renderer: renderer,
		},
	}, nil
}

func (self *Player) Tick(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.KeyboardEvent:
		if e.Type == sdl.KEYDOWN {
			switch e.Keysym.Sym {
			case sdl.K_UP:
				self.ChangePos(self.X, self.Y-10)
			case sdl.K_DOWN:
				self.ChangePos(self.X, self.Y+10)
			case sdl.K_LEFT:
				self.ChangePos(self.X-10, self.Y)
			case sdl.K_RIGHT:
				self.ChangePos(self.X+10, self.Y)
			}
		}
	}
}
