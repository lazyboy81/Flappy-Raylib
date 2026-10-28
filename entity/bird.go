package entity

import (
	"Flappy/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bird struct {
	width   int32
	height  int32
	X       float32
	Y       float32
	dy      float32
	gravity float32
}

func NewBird() Bird {
	return Bird{
		width:   constants.BirdTexture.Width,
		height:  constants.BirdTexture.Height,
		X:       float32((constants.VWidth - constants.BirdTexture.Width) / 2),
		Y:       float32((constants.VHeight - constants.BirdTexture.Height) / 2),
		dy:      0,
		gravity: 20,
	}
}

func (b *Bird) Update(dt float32) {
	b.dy = b.dy + b.gravity*dt

	if rl.IsKeyPressed(rl.KeySpace) {
		b.dy = -2
	}

	b.Y += b.dy
}

func (b *Bird) Render() {
	rl.DrawTexture(constants.BirdTexture, int32(b.X), int32(b.Y), rl.White)
}

func (b *Bird) Collides(p *Pipe) bool {
	if (b.X+2)+(float32(b.width)-4) >= p.x && b.X+2 <= p.x+pipeWidth {
		if (b.Y+2)+(float32(b.height)-4) >= p.y && b.Y+2 <= p.y+pipeHeight {
			return true
		}

	}

	return false
}
