package main

import (
	"Flappy/constants"
	"Flappy/state"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.SetTraceLogLevel(rl.LogDebug)
	rl.InitWindow(constants.WIDTH, constants.HEIGHT, "Flappy")

	defer rl.CloseWindow()

	gStateMachine := state.NewStateMachine(
		map[string]state.StateFn{
			"title": func() state.State {
				return state.NewTitleScreenState()
			},
			"play": func() state.State {
				return state.NewPlayState()
			},
			"pause": func() state.State {
				return &state.PauseState{}
			},
			"countdown": func() state.State {
				return state.NewCountDownState()
			},
			"score": func() state.State {
				return &state.ScoreState{}
			},
		},
	)
	gStateMachine.Change("title", nil)

	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	constants.LoadSounds()
	defer constants.UnloadSounds()

	musicStream := setupMusic()
	defer rl.UnloadMusicStream(*musicStream)

	constants.LoadTextures()
	defer constants.LoadTextures()

	constants.LoadFonts()
	defer constants.UnloadFonts()

	background := rl.LoadTexture("pictures/background.png")
	backgroundScroll := 0.0

	ground := rl.LoadTexture("pictures/ground.png")
	groundScroll := 0.0

	defer rl.UnloadTexture(ground)
	defer rl.UnloadTexture(background)

	renderTexture := rl.LoadRenderTexture(constants.VWidth, constants.VHeight)
	defer rl.UnloadRenderTexture(renderTexture)

	// All textures in OpenGL by default have the origin in the lower-left corner,
	// while the screen origin is by default configured in the upper left corner.
	srcTexture := rl.NewRectangle(0, 0, constants.VWidth, -constants.VHeight)
	dstTexture := rl.NewRectangle(0, 0, constants.WIDTH, constants.HEIGHT)

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(*musicStream) // Update music buffer with new stream data

		backgroundScroll = math.Mod(backgroundScroll-(constants.BackgroundScrollSpeed*float64(rl.GetFrameTime())), constants.ScrollLimit)
		groundScroll = math.Mod(groundScroll-(constants.GroundScrollSpeed*float64(rl.GetFrameTime())), constants.VWidth)

		rl.BeginTextureMode(renderTexture)

		rl.ClearBackground(rl.White)
		rl.DrawTexture(background, int32(backgroundScroll), 0, rl.White)

		gStateMachine.Update(rl.GetFrameTime())

		gStateMachine.Render()

		rl.DrawTexture(ground, int32(groundScroll), constants.VHeight-ground.Height, rl.White)

		rl.EndTextureMode()

		rl.BeginDrawing()

		rl.ClearBackground(rl.White)

		rl.DrawTexturePro(renderTexture.Texture, srcTexture, dstTexture, rl.Vector2{}, 0, rl.White)

		rl.EndDrawing()
	}

}

func setupMusic() *rl.Music {
	musicStream := rl.LoadMusicStream("music/marios_way.mp3")

	rl.PlayMusicStream(musicStream)
	rl.SetMusicPan(musicStream, 0.5)
	rl.SetMusicVolume(musicStream, 0.5)

	return &musicStream
}
