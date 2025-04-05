package enemy

import (
	"otaviocosta2110/getTheBlueBlocks/src/screen"
	"otaviocosta2110/getTheBlueBlocks/src/system"
)

type EnemyManager struct {
	Enemies []*Enemy
}

// vou deixar assim pq acho q eh legal se ficar os corpos deles no chao todo desembrenhado
func (em *EnemyManager) Update(p system.Player, screen screen.Screen) {
	activeEnemies := em.Enemies[:0]
	for _, enemy := range em.Enemies {
		if enemy.Health > 0 {
			enemy.Update(p, screen)
			activeEnemies = append(activeEnemies, enemy)
		} else {
			enemy.Object.Destroyed = true
		}
	}
	em.Enemies = activeEnemies
}

func (em *EnemyManager) Draw() {
	for _, enemy := range em.Enemies {
		enemy.Draw()
	}
}

func (em *EnemyManager) AddEnemy(e *Enemy) {
	em.Enemies = append(em.Enemies, e)
}
