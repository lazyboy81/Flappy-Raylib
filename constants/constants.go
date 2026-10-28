package constants

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH   = 1280
	HEIGHT  = 720
	VWidth  = 512
	VHeight = 288
	//VWidth                = 1280
	//VHeight               = 720
	ScrollLimit           = 413
	BackgroundScrollSpeed = 30
	GroundScrollSpeed     = 60
)

var BirdTexture rl.Texture2D
var PipeTexture rl.Texture2D
var textureOnce sync.Once

func LoadTextures() {
	textureOnce.Do(func() {
		BirdTexture = rl.LoadTexture("pictures/bird.png")
		PipeTexture = rl.LoadTexture("pictures/pipe.png")
	})
}

func UnloadTextures() {
	rl.UnloadTexture(BirdTexture)
	rl.UnloadTexture(PipeTexture)
}

var SmallFont rl.Font
var MediumFont rl.Font
var FlappyFont rl.Font
var HugeFont rl.Font
var fontsOnce sync.Once

func LoadFonts() {
	fontsOnce.Do(func() {
		SmallFont = rl.LoadFontEx("font/font.ttf", 8, nil)
		MediumFont = rl.LoadFontEx("font/flappy.ttf", 14, nil)
		FlappyFont = rl.LoadFontEx("font/flappy.ttf", 28, nil)
		HugeFont = rl.LoadFontEx("font/flappy.ttf", 56, nil)
	})
}

func UnloadFonts() {
	rl.UnloadFont(SmallFont)
	rl.UnloadFont(MediumFont)
	rl.UnloadFont(FlappyFont)
	rl.UnloadFont(HugeFont)
}

var Sounds = make(map[string]rl.Sound)
var soundsOnce sync.Once

func LoadSounds() {
	soundsOnce.Do(func() {
		Sounds["jump"] = rl.LoadSound("music/jump.wav")
		Sounds["explosion"] = rl.LoadSound("music/explosion.wav")
		Sounds["hurt"] = rl.LoadSound("music/hurt.wav")
		Sounds["score"] = rl.LoadSound("music/score.wav")
	})
}

func UnloadSounds() {
	rl.UnloadSound(Sounds["jump"])
	rl.UnloadSound(Sounds["explosion"])
	rl.UnloadSound(Sounds["hurt"])
	rl.UnloadSound(Sounds["score"])
}
