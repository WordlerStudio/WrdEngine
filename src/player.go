package WrdEngine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

type Player struct {
	*PhysicalObj
	velocityY int32
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
		PhysicalObj: &PhysicalObj{
			BaseObject: &BaseObject{
				Texture:  texture,
				X:        x,
				Y:        y,
				Width:    img.W,
				Height:   img.H,
				Renderer: renderer,
			},
		},
		velocityY: 0,
	}, nil
}

var keyState = make(map[sdl.Keycode]bool)

func (self *Player) Tick(event sdl.Event) {
	keys := sdl.GetKeyboardState()
	keyState[sdl.K_SPACE] = keys[sdl.SCANCODE_SPACE] != 0
	keyState[sdl.K_a] = keys[sdl.SCANCODE_A] != 0
	keyState[sdl.K_d] = keys[sdl.SCANCODE_D] != 0
	self.PhysicTick()
	const speed = 7
	const jumpPower = -15
	const gravity = 1

	log.Printf("Player: %v", *self)
	log.Printf("OnGround? %v", self.OnGround())
	if keyState[sdl.K_SPACE] && self.OnGround() && !self.Jumping {
		log.Printf("Jumping, btw: %v", self)
		self.Jumping = true
		for i := 10; i > 0; i-- {
			self.Y += jumpPower / 10
		}
		self.Jumping = false
	}

	self.velocityY += gravity
	self.ChangePos(self.X, self.Y+self.velocityY)

	if keyState[sdl.K_a] && self.X-speed >= 0 {
		self.ChangePos(self.X-speed, self.Y)
	}
	if keyState[sdl.K_d] && self.X+speed <= screen.Width-self.Width {
		self.ChangePos(self.X+speed, self.Y)
	}
}
