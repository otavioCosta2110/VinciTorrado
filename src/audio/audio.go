package audio

import rl "github.com/gen2brain/raylib-go/raylib"

func PlayPunch() {
	rl.PlaySound(PunchSound)
}

func PlayCollectItemSound() {
	rl.PlaySound(CollectItemSound)
}

func PlayKick() {
	rl.PlaySound(KickSound)
}

func PlayMissionMusic() {
	rl.PlayMusicStream(Mission1Music)
	rl.SetMusicVolume(Mission1Music, 0.5)
}

func UpdateMusic(music string) {
	switch(music){
	case "mission1":
		rl.UpdateMusicStream(Mission1Music)
	case "full_belly":
		rl.UpdateMusicStream(FullBellyMusic)
	}
}

func PauseMusic() {
	rl.PauseMusicStream(Mission1Music)
}

func ResumeMusic() {
	rl.ResumeMusicStream(Mission1Music)
}

func PlayFullBellyMusic() {
	rl.PlayMusicStream(FullBellyMusic)
	rl.SetMusicVolume(FullBellyMusic, 0.5)
}

func PlayWeaponBreaking() {
	rl.PlaySound(WeaponBreakingSound)
}

func PlayFullBellyAttack() {
	rl.PlaySound(FullBellyAttack)
}

func PlayFullBellyPrepare() {
	rl.PlaySound(FullBellyPrepare)
}

func SetVolume(volume float32) {
	rl.SetSoundVolume(PunchSound, volume)
	rl.SetSoundVolume(KickSound, volume)
	rl.SetSoundVolume(WeaponBreakingSound, volume)
	rl.SetSoundVolume(FullBellyAttack, volume)
	rl.SetSoundVolume(FullBellyPrepare, volume)
}
