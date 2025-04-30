package enemy

import (
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
	"slices"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyManager struct {
	Enemies         []*Enemy
	ActiveEnemies   []*Enemy
	InactiveEnemies []*Enemy
	NumEnemies      int
}

func (em *EnemyManager) Update(p system.Player, s screen.Screen) {
	cameraBounds := rl.Rectangle{
		X:      s.Camera.Target.X - float32(s.Width)/2,
		Y:      s.Camera.Target.Y - float32(s.Height)/2,
		Width:  float32(s.Width),
		Height: float32(s.Height),
	}

	for i := len(em.InactiveEnemies) - 1; i >= 0; i-- {
		enemy := em.InactiveEnemies[i]
		if isInCameraBounds(enemy, cameraBounds) {
			enemy.IsActive = true
			em.ActiveEnemies = append(em.ActiveEnemies, enemy)
			em.InactiveEnemies = slices.Delete(em.InactiveEnemies, i, i+1)
		}
	}

	for i := len(em.ActiveEnemies) - 1; i >= 0; i-- {
		enemy := em.ActiveEnemies[i]
		enemyRect := rl.Rectangle{
			X:      float32(enemy.Object.X),
			Y:      float32(enemy.Object.Y),
			Width:  float32(enemy.Object.Width),
			Height: float32(enemy.Object.Height),
		}

		if !rl.CheckCollisionRecs(enemyRect, cameraBounds) {
			enemy.IsActive = false
			em.InactiveEnemies = append(em.InactiveEnemies, enemy)
			em.ActiveEnemies = slices.Delete(em.ActiveEnemies, i, i+1)
		}
	}

	for i := len(em.ActiveEnemies) - 1; i >= 0; i-- {
		enemy := em.ActiveEnemies[i]
		enemy.Update(p, s)
		if enemy.Object.Destroyed {
			if enemy.HandleDrop != nil && enemy.EnemyType == "dwarf" {
				enemy.HandleDrop(enemy.Object.X, enemy.Object.Y)
			}
			em.ActiveEnemies = slices.Delete(em.ActiveEnemies, i, i+1)
		}
	}
}

func isInCameraBounds(enemy *Enemy, cameraBounds rl.Rectangle) bool {
	enemyRect := rl.Rectangle{
		X:      float32(enemy.Object.X),
		Y:      float32(enemy.Object.Y),
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
		enemy.Draw()
	}
}

func (em *EnemyManager) AddEnemy(e *Enemy) {
	em.Enemies = append(em.Enemies, e)
	em.InactiveEnemies = append(em.InactiveEnemies, e)
	em.NumEnemies++
}

func (em *EnemyManager) SetDropHandler(handler func(x, y int32)) {
	for _, e := range em.Enemies {
		e.HandleDrop = handler
	}
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
