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

func PlayMission2Music() {
	rl.PlayMusicStream(Mission2Music)
	rl.SetMusicVolume(Mission2Music, 0.5)
}

func PlayMission3Music() {
	rl.PlayMusicStream(Mission3Music)
	rl.SetMusicVolume(Mission3Music, 0.5)
}

func PlayEndingMusic() {
	rl.PlayMusicStream(EndingMusic)
	rl.SetMusicVolume(EndingMusic, 0.5)
}

func UpdateMusic(music string) {
	switch music {
	case "mission1":
		rl.UpdateMusicStream(Mission1Music)
	case "mission2":
		rl.UpdateMusicStream(Mission2Music)
	case "mission3":
		rl.UpdateMusicStream(Mission3Music)
	case "ending":
		rl.UpdateMusicStream(EndingMusic)
	case "full_belly":
		rl.UpdateMusicStream(FullBellyMusic)
	}
}

func StopMusic() {
	rl.StopMusicStream(Mission1Music)
	rl.StopMusicStream(Mission2Music)
	rl.StopMusicStream(FullBellyMusic)
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

func PlayShot() {
	rl.PlaySound(Shot)
}

func PlayBulletHittingTableSound() {
	rl.PlaySound(BulletHittingTableSound)
}

func PlayBombBippingSound() {
	rl.PlaySound(BombBippingSound)
}

func PlayExplosionSound() {
	rl.PlaySound(ExplosionSound)
}

func PauseBombBippingSound() {
	rl.PauseSound(BombBippingSound)
}

func ResumeBombBippingSound() {
	rl.ResumeSound(BombBippingSound)
}

func PlayGfRunningSound() {
	rl.PlaySound(GfRunningSound)
}

func PlayGfHittingWall() {
	rl.PlaySound(GfHittingWall)
}

func PlayGlassBreakingSound() {
	rl.PlaySound(GlassBreakingSound)
}

func StopGfRunningSound() {
	rl.StopSound(GfRunningSound)
}

func SetVolume(volume float32) {
	rl.SetSoundVolume(PunchSound, volume)
	rl.SetSoundVolume(KickSound, volume)
	rl.SetSoundVolume(WeaponBreakingSound, volume)
	rl.SetSoundVolume(FullBellyAttack, volume)
	rl.SetSoundVolume(FullBellyPrepare, volume)
}
