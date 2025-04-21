package enemy

import (
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
	"slices"
)

type EnemyManager struct {
	Enemies       []*Enemy
	ActiveEnemies []*Enemy
	NumEnemies    int
}

// vou deixar assim pq acho q eh legal se ficar os corpos deles no chao todo desembrenhado
func (em *EnemyManager) Update(p system.Player, screen screen.Screen) {
	for i := 0; i < len(em.ActiveEnemies); i++ {
		enemy := em.ActiveEnemies[i]
		if enemy.Health > 0 {
			enemy.Update(p, screen)
		} else {
			em.RemoveActiveEnemy(enemy)
			i--
		}
	}
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
	for _, enemy := range em.Enemies {
		enemy.Draw()
	}
}

func (em *EnemyManager) AddEnemy(e *Enemy) {
	em.Enemies = append(em.Enemies, e)
	em.ActiveEnemies = append(em.ActiveEnemies, e)
	em.NumEnemies++
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
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
