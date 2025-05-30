package enemy

import (
	// "math/rand"
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
	"otaviocosta2110/vincitorrado/src/weapon"
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
	BossProjectiles []*weapon.BossProjectile
	CurrentMap      string
}

func (em *EnemyManager) Update(p system.Player, s screen.Screen, m *string) {
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
		enemy.Update(p, s)
		if enemy.Object.Destroyed {
			em.ActiveEnemies = slices.Delete(em.ActiveEnemies, i, i+1)
		}
	}

	if em.CurrentMap == "bar" {
		if time.Since(em.lastBossShot) > 2*time.Second {
			em.lastBossShot = time.Now()

			// yPos := 400 + rand.Int31n(200)
			const bossProjectileScale = 4.0
			// bossBullet := weapon.NewBossProjectile(5000, yPos, bossProjectileScale)
			// bossBullet.Object.Flipped = true
			// em.BossProjectiles = append(em.BossProjectiles, bossBullet)
		}

		for i := 0; i < len(em.BossProjectiles); i++ {
			bullet := em.BossProjectiles[i]
			bullet.Update()

			if bullet.Object.X < -100 {
				em.BossProjectiles = slices.Delete(em.BossProjectiles, i, i+1)
				i--
			}
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
	enemies := em.Enemies
	sort.Slice(enemies, func(i, j int) bool {
		return enemies[i].Layer < enemies[j].Layer
	})
	for _, enemy := range enemies {
		if !enemy.Object.Destroyed {
			enemy.Draw()
		}
	}
}
func (em *EnemyManager) DrawDead() {
	enemies := em.Enemies
	sort.Slice(enemies, func(i, j int) bool {
		return enemies[i].Layer < enemies[j].Layer
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

func (em *EnemyManager) CheckBoxCollisions(box system.Object) {
	for _, e := range em.ActiveEnemies {
		if physics.CheckCollision(box, e.GetObject()) {
			if abs(box.KnockbackX) > 5 || abs(box.KnockbackY) > 5 {
				e.TakeDamageFromBox(box)
			}
		}
	}
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}
