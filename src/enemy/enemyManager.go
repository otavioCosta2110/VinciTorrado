package enemy

import (
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/props"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
	"slices"
	"sort"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyManager struct {
	Enemies         []*Enemy
	ActiveEnemies   []*Enemy
	InactiveEnemies []*Enemy
	NumEnemies      int
	lastBossShot    time.Time
	CurrentMap      string
}

func (em *EnemyManager) Update(p system.Player, s screen.Screen, m *string, prps []*props.Prop, buildings *rl.Texture2D, isPaused *bool) {
	for _, enemy := range em.Enemies {
		if !enemy.Object.Destroyed {
			enemyMidpoint := enemy.Object.Y + (enemy.Object.Height / 2)

			if p.GetObject().Y < enemyMidpoint {
				enemy.Object.Layer = 2 
			} else {
				enemy.Object.Layer = 0 
			}
		}

		if enemy.Object.Destroyed && enemy.EnemyType == "mafia_boss" {
			em.updateBossExplosion(enemy, p, isPaused, buildings)
		}
	}

	if *isPaused {
		return
	}
	cameraBounds := rl.Rectangle{
		X:      s.Camera.Target.X - float32(s.Width)/2,
		Y:      s.Camera.Target.Y - float32(s.Height)/2,
		Width:  float32(s.Width),
		Height: float32(s.Height),
	}

	for i := len(em.InactiveEnemies) - 1; i >= 0; i-- {
		enemy := em.InactiveEnemies[i]
		if isInCameraBounds(enemy, cameraBounds) {
			if enemy.EnemyType == "full_belly" {
				*m = "full_belly"
				audio.PlayFullBellyMusic()
			}
			enemy.Active = true
			em.ActiveEnemies = append(em.ActiveEnemies, enemy)
			em.InactiveEnemies = slices.Delete(em.InactiveEnemies, i, i+1)
		}
	}

	for i := len(em.ActiveEnemies) - 1; i >= 0; i-- {
		enemy := em.ActiveEnemies[i]
		enemy.Update(p, s, prps)
		if enemy.Object.Destroyed {
			em.ActiveEnemies = slices.Delete(em.ActiveEnemies, i, i+1)
		}
	}
}

func isInCameraBounds(enemy *Enemy, cameraBounds rl.Rectangle) bool {
	enemyRect := rl.Rectangle{
		X:      float32(enemy.Activate_pos_X),
		Y:      float32(enemy.Activate_pos_Y),
		Width:  float32(enemy.Object.Width),
		Height: float32(enemy.Object.Height),
	}
	return rl.CheckCollisionRecs(enemyRect, cameraBounds)
}

func (em *EnemyManager) RemoveActiveEnemy(enemy *Enemy) {
	for i, e := range em.ActiveEnemies {
		if e == enemy {
			em.ActiveEnemies = slices.Delete(em.ActiveEnemies, i, i+1)
			break
		}
	}
}

func (em *EnemyManager) Draw() {
	sort.Slice(em.Enemies, func(i, j int) bool {
		return em.Enemies[i].Object.Layer < em.Enemies[j].Object.Layer
	})
	for _, enemy := range em.Enemies {
		if !enemy.Object.Destroyed {
			enemy.Draw()
		}
	}
}

func (em *EnemyManager) DrawDead() {
	enemies := em.Enemies
	sort.Slice(enemies, func(i, j int) bool {
		return enemies[i].Object.Layer < enemies[j].Object.Layer
	})
	for _, enemy := range enemies {
		if enemy.Object.Destroyed {
			enemy.Draw()
		}
	}
}

func (em *EnemyManager) AddEnemy(e *Enemy) {
	em.Enemies = append(em.Enemies, e)
	em.InactiveEnemies = append(em.InactiveEnemies, e)
	em.NumEnemies++
}

func (em *EnemyManager) updateBossExplosion(enemy *Enemy, p system.Player, isPaused *bool, buildings *rl.Texture2D) {
	if !enemy.Exploded {
		if !enemy.HasExplosionPlayedSound {
			audio.PlayBombBippingSound()
			enemy.ExplosionStart = time.Now()
			enemy.HasExplosionPlayedSound = true
			enemy.ExplosionElapsed = 0
		}

		if *isPaused && !enemy.ExplosionPaused {
			enemy.ExplosionPaused = true
			enemy.ExplosionPauseTime = time.Now()
			audio.PauseBombBippingSound()
		} else if !*isPaused && enemy.ExplosionPaused {
			enemy.ExplosionPaused = false
			enemy.ExplosionElapsed += time.Since(enemy.ExplosionPauseTime)
			audio.ResumeBombBippingSound()
		}

		if !*isPaused {
			elapsed := time.Since(enemy.ExplosionStart) - enemy.ExplosionElapsed
			if elapsed >= time.Duration(5.071*float64(time.Second)) && !enemy.Exploded {
				enemy.Explode(p)
				enemy.Exploded = true
				explodedBuildingsPath := "assets/scenes/bar_exploded.png"
				*buildings = system.LoadScaledTexture(explodedBuildingsPath, enemy.Object.Scale)
			}
		}
	}
}
