package state

import (
	"Flappy/constants"
	"Flappy/entity"
	"Flappy/util"
	"math"
	"math/rand/v2"
	"slices"
	"strconv"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var gamePaused = false

type State interface {
	Enter(map[string]any)
	Update(dt float32)
	Render()
}

type emptyState struct{}

func (es *emptyState) Enter(params map[string]any) {}

func (es *emptyState) Update(dt float32) {}

func (es *emptyState) Render() {}

type TitleScreenState struct{}

func NewTitleScreenState() *TitleScreenState {
	return &TitleScreenState{}
}

func (ts *TitleScreenState) Enter(params map[string]any) {

}

func (ts *TitleScreenState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeyEnter) {
		gStateMachine.Change("countdown", nil)
	}
}

func (ts *TitleScreenState) Render() {
	rl.DrawTextEx(constants.FlappyFont, "Flappy", rl.NewVector2(220, 64), 28, 0, rl.White)
	rl.DrawTextEx(constants.MediumFont, "Press Enter", rl.NewVector2(220, 100), 14, 0, rl.White)
}

func NewCountDownState() *CountDownState {
	return &CountDownState{count: 3, timer: 0}
}

type CountDownState struct {
	gameParams map[string]any
	count      int
	timer      float64
}

func (cs *CountDownState) Enter(params map[string]any) {
	if params != nil {
		cs.gameParams = params
	}
}

func (cs *CountDownState) Update(dt float32) {
	cs.timer += float64(dt)

	// to see if we have passed a second or not
	if cs.timer >= (1 * time.Second.Seconds()) {
		cs.timer = math.Mod(cs.timer, time.Second.Seconds()*1)
		cs.count -= 1

		if cs.count == 0 {
			gStateMachine.Change("play", cs.gameParams)
		}
	}
}

func (cs *CountDownState) Render() {
	rl.DrawTextEx(constants.HugeFont, strconv.Itoa(cs.count), rl.NewVector2(220, 64), 56, 0, rl.White)
}

type PlayState struct {
	bird           entity.Bird
	pipePairs      []*entity.PipePair
	timer          float32
	timerThreshold float32
	score          int
	lastY          float64
}

func NewPlayState() *PlayState {
	return &PlayState{
		bird:           entity.NewBird(),
		pipePairs:      []*entity.PipePair{},
		timer:          0,
		timerThreshold: util.RandRangeFloat32(2, 5), // return a number between [2, 5]
		score:          0,
		lastY: float64(
			-constants.PipeTexture.Height + rand.Int32N(80) + 21,
		),
	}
}

func (ps *PlayState) Enter(params map[string]any) {
	if params != nil {
		ps.bird = params["bird"].(entity.Bird)
		ps.pipePairs = params["pipePairs"].([]*entity.PipePair)
		ps.timer = params["timer"].(float32)
		ps.timerThreshold = params["timerThreshold"].(float32)
		ps.score = params["score"].(int)
		ps.lastY = params["lastY"].(float64)
	}
}

func (ps *PlayState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeyP) {
		gamePaused = true
		gStateMachine.Change("pause", map[string]any{
			"bird":           ps.bird,
			"pipePairs":      ps.pipePairs,
			"timer":          ps.timer,
			"timerThreshold": ps.timerThreshold,
			"score":          ps.score,
			"lastY":          ps.lastY,
		})
	}

	if gamePaused {
		return
	}

	ps.timer += dt
	if ps.timer > ps.timerThreshold {
		pipeHeight := float64(constants.PipeTexture.Height)

		minY := -pipeHeight + 10
		maxY := float64(constants.VHeight) - 90 - pipeHeight
		y := math.Max(
			minY,
			math.Min(
				ps.lastY+util.RandRangeFloat64(-20, 20),
				maxY,
			),
		)
		ps.lastY = y

		ps.pipePairs = append(ps.pipePairs, entity.NewPipePair(float32(y)))

		ps.timer = 0
		ps.timerThreshold = util.RandRangeFloat32(2, 5)
	}

	for _, p := range ps.pipePairs {
		if !p.Scored && p.X+float32(constants.PipeTexture.Width) < ps.bird.X {
			ps.score += 1
			p.Scored = true
			rl.PlaySound(constants.Sounds["score"])
		}

		p.Update(dt)
	}

	ps.pipePairs = slices.DeleteFunc(
		ps.pipePairs,
		func(pair *entity.PipePair) bool {
			return pair.Remove
		},
	)

	ps.bird.Update(dt)

	// we check if the bird collides with the pipes
	for _, pp := range ps.pipePairs {
		for _, p := range pp.Pipes {
			if ps.bird.Collides(p) {
				rl.PlaySound(constants.Sounds["explosion"])
				rl.PlaySound(constants.Sounds["hurt"])
				gStateMachine.Change("score", map[string]any{
					"score": ps.score,
				})
			}
		}
	}

	if ps.bird.Y > constants.VHeight+15 {
		rl.PlaySound(constants.Sounds["explosion"])
		rl.PlaySound(constants.Sounds["hurt"])
		gStateMachine.Change("score", map[string]any{
			"score": ps.score,
		})
	}

}

func (ps *PlayState) Render() {
	for _, p := range ps.pipePairs {
		p.Render()
	}

	rl.DrawTextEx(constants.FlappyFont, "Score: "+strconv.Itoa(ps.score), rl.NewVector2(8, 8), 28, 0, rl.White)

	ps.bird.Render()
}

type PauseState struct {
	gameParams map[string]any
}

func (ps *PauseState) Enter(params map[string]any) {
	ps.gameParams = params
}

func (ps *PauseState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeyR) {
		gamePaused = false
		gStateMachine.Change("play", ps.gameParams)
	}
}

func (ps *PauseState) Render() {
	rl.DrawTextEx(constants.FlappyFont, "Paused", rl.NewVector2(220, 64), 28, 0, rl.White)
	rl.DrawTextEx(constants.MediumFont, "Press r to resume", rl.NewVector2(220, 100), 14, 0, rl.White)
}

type ScoreState struct {
	score int
}

func (ss *ScoreState) Enter(params map[string]any) {
	ss.score = params["score"].(int)
}

func (ss *ScoreState) Update(dt float32) {
	if rl.IsKeyPressed(rl.KeyEnter) {
		gStateMachine.Change("countdown", nil)
	}
}

func (ss *ScoreState) Render() {
	rl.DrawTextEx(constants.FlappyFont, "Oof! you lost!", rl.NewVector2(220, 64), 28, 0, rl.White)
	rl.DrawTextEx(constants.MediumFont, "Score: "+strconv.Itoa(ss.score), rl.NewVector2(220, 100), 14, 0, rl.White)
	rl.DrawTextEx(constants.MediumFont, "Press Enter to play again", rl.NewVector2(220, 150), 14, 0, rl.White)
}
