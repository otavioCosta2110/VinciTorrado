package audio

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	PunchSound          rl.Sound
	KickSound           rl.Sound
	WeaponBreakingSound rl.Sound
)

func LoadSounds() {
	PunchSound = rl.LoadSound("assets/sounds/sound_punch.mp3")
	KickSound = rl.LoadSound("assets/sounds/Kick_Trash.mp3")
	WeaponBreakingSound = rl.LoadSound("assets/sounds/weapon_breaking.mp3")
}

func UnloadSounds() {
	rl.UnloadSound(PunchSound)
	rl.UnloadSound(KickSound)
}
