package enemy

import (
	"otaviocosta2110/getTheBlueBlocks/src/screen"
	"otaviocosta2110/getTheBlueBlocks/src/system"
	"slices"
)

type EnemyManager struct {
	Enemies       []*Enemy
	ActiveEnemies []*Enemy
	NumEnemies    int
}

// vou deixar assim pq acho q eh legal se ficar os corpos deles no chao todo desembrenhado
func (em *EnemyManager) Update(p system.Player, screen screen.Screen) {
	for _, enemy := range em.ActiveEnemies {
		if enemy.Health > 0 {
			enemy.Update(p, screen)
		} else {
			em.RemoveActiveEnemy(enemy)
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
