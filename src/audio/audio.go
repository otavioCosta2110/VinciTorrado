package audio

import rl "github.com/gen2brain/raylib-go/raylib"

func PlayPunch() {
	rl.PlaySound(PunchSound)
}

func SetVolume(volume float32) {
	rl.SetSoundVolume(PunchSound, volume)
}
