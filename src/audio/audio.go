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

func UpdateMusic() {
	rl.UpdateMusicStream(Mission1Music)
}

func PauseMusic() {
	rl.PauseMusicStream(Mission1Music)
}

func ResumeMusic() {
	rl.ResumeMusicStream(Mission1Music)
}

func PlayWeaponBreaking() {
	rl.PlaySound(WeaponBreakingSound)
}

func SetVolume(volume float32) {
	rl.SetSoundVolume(PunchSound, volume)
	rl.SetSoundVolume(KickSound, volume)
	rl.SetSoundVolume(WeaponBreakingSound, volume)
}
