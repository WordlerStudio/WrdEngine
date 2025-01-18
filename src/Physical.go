package WrdEngine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type PhysicalObj struct {
	*BaseObject
	LockPosition bool
}

var physicalObjects []*PhysicalObj

func NewPhysicalObj(renderer *sdl.Renderer, imagePath string, x, y int32) (*PhysicalObj, error) {
	img, err := sdl.LoadBMP(imagePath)
	if err != nil {
		return nil, fmt.Errorf("Image load error: ", err)
	}
	defer img.Free()

	texture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, err
	}

	nObj := &PhysicalObj{
		BaseObject: &BaseObject{
			Texture:  texture,
			X:        x,
			Y:        y,
			Width:    img.W,
			Height:   img.H,
			Renderer: renderer,
		},
	}
	physicalObjects = append(physicalObjects, nObj)
	return nObj, nil
}

func NewPhysicalBaseObj(renderer *sdl.Renderer, color sdl.Color, x, y, w, h int32) (*PhysicalObj, error) {
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, w, h)
	if err != nil {
		return nil, fmt.Errorf("Texture creation error: %v", err)
	}

	if err := renderer.SetRenderTarget(texture); err != nil {
		return nil, fmt.Errorf("Set render target error: %v", err)
	}

	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	if err := renderer.Clear(); err != nil {
		return nil, fmt.Errorf("Clear texture error: %v", err)
	}

	renderer.SetRenderTarget(nil)

	nObj := &PhysicalObj{
		BaseObject: &BaseObject{
			Texture:  texture,
			X:        x,
			Y:        y,
			Width:    w,
			Height:   h,
			Renderer: renderer,
		},
	}
	physicalObjects = append(physicalObjects, nObj)
	return nObj, nil
}

func (self *PhysicalObj) Lock() {
	self.LockPosition = true
}

func (self *PhysicalObj) UnLock() {
	self.LockPosition = false
}

func (self *PhysicalObj) PhysicTick() {
	if !self.OnGround() && !self.LockPosition {
		self.Y += int32(physic.GravtyPower)
	}
	if self.Y > screen.Height {
		self.Y = screen.Height
	}
}

func (self *PhysicalObj) OnGround() bool {
	for _, obj := range physicalObjects {
		if obj == self {
			continue
		}

		fmt.Printf("Checking obj: X=%d, Y=%d, Width=%d, Height=%d\n", obj.X, obj.Y, obj.Width, obj.Height)
		fmt.Printf("Self: X=%d, Y=%d, Width=%d, Height=%d\n", self.X, self.Y, self.Width, self.Height)

		if self.X < obj.X+obj.Width && self.X+self.Width > obj.X && self.Y+self.Height >= obj.Y {
			//fmt.Println("OnGround=True")
			return true
		}
	}
	fmt.Println("OnGround=False")
	return false
}

func (self *PhysicalObj) CheckCollisionLeft() bool {
	for _, obj := range physicalObjects {
		if obj == self {
			continue
		}

		if self.X < obj.X+obj.Width && self.X+self.Width > obj.X && self.Y+self.Height > obj.Y && self.Y < obj.Y+obj.Height {
			return true
		}
	}
	return false
}

func (self *PhysicalObj) CheckCollisionRight() bool {
	for _, obj := range physicalObjects {
		if obj == self {
			continue
		}

		if self.X+self.Width > obj.X && self.X < obj.X+obj.Width && self.Y+self.Height > obj.Y && self.Y < obj.Y+obj.Height {
			return true
		}
	}
	return false
}
