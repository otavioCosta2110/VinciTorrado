package cutscene

import (
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/system"
)

func (c *Cutscene) IntroCutscenes(player system.Live, gf system.Live, enemyManager *enemy.EnemyManager) {
	fullBelly := enemyManager.Enemies[len(enemyManager.Enemies)-1]

	posXBelly := fullBelly.Object.X
	posYBelly := fullBelly.Object.Y
	weaponBelly := fullBelly.Weapon

	fullBelly.Object.X = 1400
	fullBelly.Object.Y = 400
	fullBelly.Weapon = nil

	c.AddAction(NewObjectMoveAction(player, 200, float32(player.GetObject().Y), 1, "walk"))
	c.AddAction(NewObjectMoveAction(fullBelly, 1110, float32(fullBelly.GetObject().Y), 2, "walk"))
	c.AddAction(NewWaitAction(0.5))
	c.AddAction(NewObjectMoveAction(fullBelly, 1105, float32(fullBelly.GetObject().Y), 2, "fb_charge"))
	c.AddAction(NewWaitAction(0.5))
	c.AddAction(NewObjectMoveAction(fullBelly, 1100, float32(fullBelly.GetObject().Y), 2, "fb_attack"))
	c.AddAction(NewWaitAction(0.5))
	c.AddAction(NewCallbackAction(func() {
		gf.SetActive(false)
	}))
	c.AddAction(NewObjectMoveAction(fullBelly, 1300, float32(fullBelly.GetObject().Y), 2, "fb_walk_with_girl"))

	c.AddAction(NewCallbackAction(func() {
		fullBelly.Object.X = posXBelly
		fullBelly.Object.Y = posYBelly
		fullBelly.Weapon = weaponBelly
	}))
}

func (c *Cutscene) BarIntroCutscene(player system.Live, gf system.Live, enemyManager *enemy.EnemyManager) {
	mafiaBoss := enemyManager.Enemies[len(enemyManager.Enemies)-1]

	posXMafia := mafiaBoss.Object.X
	posYMafia := mafiaBoss.Object.Y
	weaponMafia := mafiaBoss.Weapon

	mafiaBoss.Object.X = 4000
	mafiaBoss.Object.Y = 400
	mafiaBoss.Weapon = nil

	println(player.GetObject().Y)
	println(mafiaBoss.Object.X)

		gf.SetActive(false)
		mafiaBoss.Object.FrameY=4
		mafiaBoss.UpdateAnimation("fb_walk_with_girl")
		c.AddAction(NewObjectMoveAction(mafiaBoss, 1300, float32(mafiaBoss.GetObject().Y), 2, "fb_walk_with_girl"))

	c.AddAction(NewCallbackAction(func() {
		mafiaBoss.Object.X = posXMafia
		mafiaBoss.Object.Y = posYMafia
		mafiaBoss.Weapon = weaponMafia
	}))
}
