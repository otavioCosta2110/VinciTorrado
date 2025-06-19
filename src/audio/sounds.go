package audio

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	PunchSound              rl.Sound
	KickSound               rl.Sound
	WeaponBreakingSound     rl.Sound
	Mission1Music           rl.Music
	CollectItemSound        rl.Sound
	FullBellyAttack         rl.Sound
	FullBellyPrepare        rl.Sound
	Shot                    rl.Sound
	BulletHittingTableSound rl.Sound
	BombBippingSound        rl.Sound
	ExplosionSound          rl.Sound
	GfRunningSound          rl.Sound
	GfHittingWall           rl.Sound

	FullBellyMusic rl.Music
	Mission2Music  rl.Music
	Mission3Music  rl.Music
)

func LoadSounds() {
	PunchSound = rl.LoadSound("assets/sounds/sound_punch.mp3")
	KickSound = rl.LoadSound("assets/sounds/Kick_Trash.mp3")
	WeaponBreakingSound = rl.LoadSound("assets/sounds/weapon_breaking.mp3")
	CollectItemSound = rl.LoadSound("assets/sounds/collect_item.mp3")
	FullBellyAttack = rl.LoadSound("assets/sounds/full_belly_attack.mp3")
	FullBellyPrepare = rl.LoadSound("assets/sounds/full_belly_prepare.mp3")
	Shot = rl.LoadSound("assets/sounds/shot.mp3")
	BulletHittingTableSound = rl.LoadSound("assets/sounds/bullet_hitting_table.mp3")
	BombBippingSound = rl.LoadSound("assets/sounds/bipping_bomb.mp3")
	ExplosionSound = rl.LoadSound("assets/sounds/explosion.mp3")
	GfRunningSound = rl.LoadSound("assets/sounds/gf_running.mp3")
	GfHittingWall = rl.LoadSound("assets/sounds/gf_hitting_wall.mp3")

	Mission1Music = rl.LoadMusicStream("assets/sounds/mission1.mp3")
	FullBellyMusic = rl.LoadMusicStream("assets/sounds/music_fullbelly.mp3")
	Mission2Music = rl.LoadMusicStream("assets/sounds/mission2.mp3")
	Mission3Music = rl.LoadMusicStream("assets/sounds/mission3.mp3")
}

func UnloadSounds() {
	rl.UnloadSound(PunchSound)
	rl.UnloadSound(KickSound)
	rl.UnloadSound(CollectItemSound)
	rl.UnloadSound(WeaponBreakingSound)
	rl.UnloadSound(FullBellyAttack)
	rl.UnloadSound(FullBellyPrepare)
	rl.UnloadSound(Shot)
	rl.UnloadSound(BulletHittingTableSound)
	rl.UnloadSound(BombBippingSound)
	rl.UnloadSound(GfRunningSound)
	rl.UnloadSound(GfHittingWall)

	rl.UnloadMusicStream(Mission1Music)
	rl.UnloadMusicStream(Mission2Music)
	rl.UnloadMusicStream(Mission3Music)
	rl.UnloadMusicStream(FullBellyMusic)
}

func UnloadMusic() {
	rl.UnloadMusicStream(Mission1Music)
	rl.UnloadMusicStream(FullBellyMusic)
	rl.UnloadMusicStream(Mission2Music)
	rl.UnloadMusicStream(Mission3Music)
}
