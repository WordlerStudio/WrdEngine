package WrdEngine

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	Name       string
	Background *sdl.Texture // Tekstura tła
	Objects    []BaseObject
}

func NewScene(name string, renderer *sdl.Renderer, background *sdl.Texture, objects ...BaseObject) (*Scene, error) {
	return &Scene{Name: name, Background: background, Objects: objects}, nil
}

func (self *Scene) AddObj(obj BaseObject) {
	self.Objects = append(self.Objects, obj)
}

func (self *Scene) RenderBackground(renderer *sdl.Renderer) error {
	if self.Background != nil {
		renderer.Copy(self.Background, nil, nil)
	} else {
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
	}
	return nil
}

func (self *Scene) Refresh(event sdl.Event) {
	for _, obj := range self.Objects {
		obj.Tick()
	}
}
