package WrdEngine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Event uint

const (
	LeftClicked Event = iota
	_                 // RightClicked
	MiddleClicked
	// TODO: Add more events
)

type BaseObject struct {
	X, Y              int32
	Width, Height     int32
	Texture           *sdl.Texture
	Renderer          *sdl.Renderer
	Addons            []Addon
	events            map[Event][]func()
	mousePressedLeft  bool
	mousePressedRight bool
	mousePressedMid   bool
}

func NewObj(renderer *sdl.Renderer, imagePath string, x, y int32) (*BaseObject, error) {
	img, err := sdl.LoadBMP(imagePath)
	if err != nil {
		return nil, ImageLoadError(fmt.Errorf("an error occurred while loading an image: %s", err))
	}
	defer img.Free()

	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, ImageLoadError(fmt.Errorf("an error occurred while loading an image: %s", err))
	}

	return &BaseObject{
		Texture:  texture,
		X:        x,
		Y:        y,
		Width:    img.W,
		Height:   img.H,
		Renderer: renderer,
		events:   make(map[Event][]func()),
	}, nil
}

func (self *BaseObject) Connect(event Event, handler func()) {
	self.events[event] = append(self.events[event], handler)
}

func (self *BaseObject) Emit(event Event) {
	if handlers, ok := self.events[event]; ok {
		for _, handler := range handlers {
			handler()
		}
	}
}

func (self *BaseObject) ChangePos(x, y int32) {
	self.X = x
	self.Y = y
}

func (self *BaseObject) Attach(addon Addon) {
	self.Addons = append(self.Addons, addon)
}

func (self *BaseObject) Render() error {
	for _, addon := range self.Addons {
		addon.Start(self)
	}
	dst := &sdl.Rect{X: self.X, Y: self.Y, W: self.Width, H: self.Height}
	if err := self.Renderer.Copy(self.Texture, nil, dst); err != nil {
		return SdlError(fmt.Errorf("an error occured while rendering object: %v", err))
	}
	return nil
}

func (self *BaseObject) Tick() {
	for _, addon := range self.Addons {
		addon.Tick(self)
	}
	mouseX, mouseY, mouseState := sdl.GetMouseState()
	if mouseState&sdl.BUTTON_LEFT != 0 {
		if !self.mousePressedLeft &&
			mouseX >= self.X && mouseX <= self.X+self.Width &&
			mouseY >= self.Y && mouseY <= self.Y+self.Height {
			self.Emit(LeftClicked)
		}
		self.mousePressedLeft = true
	} else {
		self.mousePressedLeft = false
	}

	//if mouseState&sdl.BUTTON_RIGHT != 0 {
	//	if !self.mousePressedRight &&
	//		mouseX >= self.X && mouseX <= self.X+self.Width &&
	//		mouseY >= self.Y && mouseY <= self.Y+self.Height {
	//		self.Emit(RightClicked)
	//	}
	//	self.mousePressedRight = true
	//} else {
	//	self.mousePressedRight = false
	//}

	if mouseState&sdl.BUTTON_MIDDLE != 0 {
		if !self.mousePressedMid &&
			mouseX >= self.X && mouseX <= self.X+self.Width &&
			mouseY >= self.Y && mouseY <= self.Y+self.Height {
			self.Emit(MiddleClicked)
		}
		self.mousePressedMid = true
	} else {
		self.mousePressedMid = false
	}
}
