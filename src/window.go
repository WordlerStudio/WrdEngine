package WrdEngine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	Window        *sdl.Window
	Renderer      *sdl.Renderer
	Scene         *Scene
	Width, Height int32
}

func NewWindow(title string, size ...int32) (*Window, error) {
	var width, height int32
	if len(size) == 2 {
		width = size[0]
		height = size[1]
	} else if len(size) == 1 {
		width = size[0]
		height = size[0]
	} else {
		width = 800
		height = 600
	}
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, WindowError(fmt.Errorf("failed to create window: %v", err))
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		window.Destroy()
		return nil, WindowError(fmt.Errorf("failed to create renderer: %v", err))
	}

	return &Window{Window: window, Renderer: renderer, Scene: nil}, nil
}

func (self *Window) SetScene(scene *Scene) {
	self.Scene = scene
	self.RenderScene()
}

func (self *Window) RenderScene() error {
	if self.Scene == nil {
		return fmt.Errorf("no scene is set to render")
	}

	if err := self.Scene.RenderBackground(self.Renderer); err != nil {
		return err
	}

	for _, obj := range self.Scene.Objects {
		if err := obj.Render(); err != nil {
			return err
		}
	}

	self.Renderer.Present()
	return nil
}

func (self *Window) Destroy() {
	if self.Renderer != nil {
		self.Renderer.Destroy()
	}
	if self.Window != nil {
		self.Window.Destroy()
	}
}
