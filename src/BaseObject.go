package WrdEngine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type BaseObject struct {
	Texture       *sdl.Texture
	X, Y          int32
	Width, Height int32
	Renderer      *sdl.Renderer
}

func NewObj(renderer *sdl.Renderer, imagePath string, x, y int32) (*BaseObject, error) {
	img, err := sdl.LoadBMP(imagePath)
	if err != nil {
		return nil, fmt.Errorf("Image load error: ", err)
	}
	defer img.Free()

	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, err
	}

	return &BaseObject{
		Texture:  texture,
		X:        x,
		Y:        y,
		Width:    img.W,
		Height:   img.H,
		Renderer: renderer,
	}, nil
}

func (self *BaseObject) ChangePos(x, y int32) {
	self.X = x
	self.Y = y
}

func (self *BaseObject) Render() {
	rect := &sdl.Rect{X: self.X, Y: self.Y, W: self.Width, H: self.Height}
	self.Renderer.Copy(self.Texture, nil, rect)
}

func (self *BaseObject) Tick(_ sdl.Event) {
	// ... it's only template!
}
