package cutscene

import (
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/system"
)

func (c *Cutscene) IntroCutscenes(player system.Live, enemyManager *enemy.EnemyManager) {
	fullBelly:=enemyManager.Enemies[len(enemyManager.Enemies)-1]

	posXBelly := fullBelly.Object.X
	posYBelly := fullBelly.Object.Y
	weaponBelly := fullBelly.Weapon

	fullBelly.Object.X = 1400
	fullBelly.Object.Y = 400
	fullBelly.Weapon = nil

	// c.AddAction(NewObjectMoveAction(fullBelly, 500, float32(fullBelly.GetObject().Y), 2, "fb_walk_with_girl"))
	c.AddAction(NewObjectMoveAction(player, 200, float32(player.GetObject().Y), 1, "walk"))
	player.IsActive()

	c.AddAction(NewCallbackAction(func() {
		fullBelly.Object.X = posXBelly
		fullBelly.Object.Y = posYBelly
		fullBelly.Weapon = weaponBelly
	}))
}
