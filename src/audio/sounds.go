package audio

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	PunchSound rl.Sound
)

func LoadSounds() {
	PunchSound = rl.LoadSound("assets/sounds/sound_punch.mp3")
}

func UnloadSounds() {
	rl.UnloadSound(PunchSound)
}
