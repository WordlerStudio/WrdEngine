package WrdEngine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	Renderer          *sdl.Renderer
	Window            *sdl.Window
	BackgroundTexture *sdl.Texture
}

func NewWindow(title string, width, height int32, background string) (*Window, error) {
	screen = &Screen{Width: width, Height: height}
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, SdlError(fmt.Errorf("an error occurred while initializing SDL: %v", err))
	}

	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, width, height, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, WindowError(fmt.Errorf("an error occurred while creating the window: %v", err))
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, WindowError(fmt.Errorf("an error occurred while creating the renderer: %v", err))
	}

	// Load the background image
	image, err := sdl.LoadBMP(background)
	if err != nil {
		return nil, ImageLoadError(fmt.Errorf("an error occurred while loading the background image: %v", err))
	}
	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		return nil, ImageCreateError(fmt.Errorf("an error occurred while creating the texture from the background image: %v", err))
	}

	return &Window{Renderer: renderer, Window: window, BackgroundTexture: texture}, nil
}

func (self *Window) Render() {
	self.Renderer.Copy(self.BackgroundTexture, nil, nil)
}

func (self *Window) Destroy() {
	self.BackgroundTexture.Destroy()
	self.Renderer.Destroy()
	self.Window.Destroy()
	sdl.Quit()
}
