package audio

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	PunchSound          rl.Sound
	KickSound           rl.Sound
	WeaponBreakingSound rl.Sound
	Mission1Music       rl.Music
	CollectItemSound    rl.Sound
	FullBellyAttack     rl.Sound
	FullBellyPrepare    rl.Sound
)

func LoadSounds() {
	PunchSound = rl.LoadSound("assets/sounds/sound_punch.mp3")
	KickSound = rl.LoadSound("assets/sounds/Kick_Trash.mp3")
	WeaponBreakingSound = rl.LoadSound("assets/sounds/weapon_breaking.mp3")
	CollectItemSound = rl.LoadSound("assets/sounds/collect_item.mp3")
	Mission1Music = rl.LoadMusicStream("assets/sounds/mission1.mp3")
	FullBellyAttack = rl.LoadSound("assets/sounds/full_belly_attack.mp3")
	FullBellyPrepare = rl.LoadSound("assets/sounds/full_belly_prepare.mp3")
}

func UnloadSounds() {
	rl.UnloadSound(PunchSound)
	rl.UnloadSound(KickSound)
	rl.UnloadSound(CollectItemSound)
	rl.UnloadMusicStream(Mission1Music)
	rl.UnloadSound(WeaponBreakingSound)
	rl.UnloadSound(FullBellyAttack)
	rl.UnloadSound(FullBellyPrepare)
}
