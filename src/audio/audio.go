package audio

import rl "github.com/gen2brain/raylib-go/raylib"

func PlayPunch() {
	rl.PlaySound(PunchSound)
}

func PlayKick() {
	rl.PlaySound(KickSound)
}

func PlayWeaponBreaking() {
	rl.PlaySound(WeaponBreakingSound)
}

func SetVolume(volume float32) {
	rl.SetSoundVolume(PunchSound, volume)
	rl.SetSoundVolume(KickSound, volume)
	rl.SetSoundVolume(WeaponBreakingSound, volume)
}
