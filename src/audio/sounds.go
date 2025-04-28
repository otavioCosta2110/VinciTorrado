package audio

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	PunchSound rl.Sound
	KickSound  rl.Sound
)

func LoadSounds() {
	PunchSound = rl.LoadSound("assets/sounds/sound_punch.mp3")
	KickSound = rl.LoadSound("assets/sounds/kick_box.mp3")

}

func UnloadSounds() {
	rl.UnloadSound(PunchSound)
	rl.UnloadSound(KickSound)
}
