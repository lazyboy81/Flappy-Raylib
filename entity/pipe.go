package entity

import (
	"Flappy/constants"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	pipeWidth  float32 = 70
	pipeHeight float32 = 288
	pipeSpeed  float32 = 60
)

type orientation int

const (
	top orientation = iota
	bottom
)

type Pipe struct {
	x float32
	y float32
	o orientation
}

func newPipe(or orientation, y float32) *Pipe {
	return &Pipe{
		x: constants.VWidth + 32,
		y: y,
		o: or,
	}
}

func (p *Pipe) Update(_ float32) {

}

func (p *Pipe) Render() {
	texture := constants.PipeTexture

	var src rl.Rectangle
	if p.o == top {
		src = rl.NewRectangle(0, 0, float32(texture.Width), -float32(texture.Height))
	} else {
		src = rl.NewRectangle(0, 0, float32(texture.Width), float32(texture.Height))
	}
	
	rl.DrawTexturePro(
		texture,
		src,
		rl.NewRectangle(p.x, p.y, pipeWidth, pipeHeight),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
}

type PipePair struct {
	X      float32
	Y      float32
	Pipes  map[string]*Pipe
	Remove bool
	Scored bool
}

func NewPipePair(y float32) *PipePair {
	return &PipePair{
		X: constants.VWidth + 32,
		Y: y,
		Pipes: map[string]*Pipe{
			"top":    newPipe(top, y),
			"bottom": newPipe(bottom, y+pipeHeight+float32(rl.GetRandomValue(60, 100))),
		},
		Remove: false,
		Scored: false,
	}
}

func (pp *PipePair) Update(dt float32) {
	if pp.X > -pipeWidth {
		pp.X = pp.X - (pipeSpeed * dt)
		pp.Pipes["top"].x = pp.X
		pp.Pipes["bottom"].x = pp.X
	} else {
		pp.Remove = true
	}
}

func (pp *PipePair) Render() {
	for _, p := range pp.Pipes {
		p.Render()
	}
}
