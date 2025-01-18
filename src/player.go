package WrdEngine

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	*PhysicalObj
	Jumping      bool
	VelocityY    int32
	Speed        float64
	baseSpeed    float64
	maxSpeed     float64
	acceleration float64
	deceleration float64
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
		baseSpeed:    10.0,
		maxSpeed:     20.0,
		acceleration: 0.5,
		deceleration: 0.8,
		VelocityY:    0,
	}, nil
}

var keyState = make(map[sdl.Keycode]bool)

func (self *Player) Tick(event sdl.Event) {
	keys := sdl.GetKeyboardState()
	keyState[sdl.K_SPACE] = keys[sdl.SCANCODE_SPACE] != 0
	keyState[sdl.K_a] = keys[sdl.SCANCODE_A] != 0
	keyState[sdl.K_d] = keys[sdl.SCANCODE_D] != 0
	const speed = 10
	const jumpPower = -70

	if keyState[sdl.K_SPACE] && self.OnGround() && !self.Jumping {
		self.Jumping = true
		self.VelocityY = jumpPower
	}

	moving := false

	if keyState[sdl.K_a] && self.X-int32(self.Speed) >= 0 && !self.CheckCollisionLeft() {
		self.Speed = min(self.Speed+self.acceleration, self.maxSpeed)
		self.ChangePos(self.X-int32(self.Speed), self.Y)
		moving = true
	}

	if keyState[sdl.K_d] && self.X+int32(self.Speed) <= screen.Width-self.Width && !self.CheckCollisionRight() {
		self.Speed = min(self.Speed+self.acceleration, self.maxSpeed)
		self.ChangePos(self.X+int32(self.Speed), self.Y)
		moving = true
	}

	if !moving {
		self.Speed = max(self.Speed-self.deceleration, self.baseSpeed)
	}
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// PhysicTick - Override method from PhysicalObject to add jumping and Velocity
func (self *Player) PhysicTick() {
	if !self.OnGround() {
		self.VelocityY += int32(physic.GravtyPower)
	} else if self.VelocityY > 0 {
		self.VelocityY = 0
		self.Jumping = false
	}

	self.Y += self.VelocityY

	if self.OnGround() {
		self.Y = 600 - self.Height
	}
}
