package audio

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	PunchSound          rl.Sound
	KickSound           rl.Sound
	WeaponBreakingSound rl.Sound
	Mission1Music       rl.Music
)

func LoadSounds() {
	PunchSound = rl.LoadSound("assets/sounds/sound_punch.mp3")
	KickSound = rl.LoadSound("assets/sounds/Kick_Trash.mp3")
	WeaponBreakingSound = rl.LoadSound("assets/sounds/weapon_breaking.mp3")
	Mission1Music = rl.LoadMusicStream("assets/sounds/mission.mp3")
}

func UnloadSounds() {
	rl.UnloadSound(PunchSound)
	rl.UnloadSound(KickSound)
	rl.UnloadMusicStream(Mission1Music)
}
