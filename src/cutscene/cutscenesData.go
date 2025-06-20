package cutscene

import (
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/props"
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

	mafiaBoss.Object.X = 600
	mafiaBoss.Object.Y = 400
	mafiaBoss.Weapon = nil

	gf.SetActive(false)
	mafiaBoss.Object.FrameY = 4
	mafiaBoss.UpdateAnimation("fb_walk_with_girl")
	c.AddAction(NewObjectMoveAction(mafiaBoss, 1300, float32(mafiaBoss.GetObject().Y), 3, "fb_walk_with_girl"))

	c.AddAction(NewCallbackAction(func() {
		mafiaBoss.Object.X = posXMafia
		mafiaBoss.Object.Y = posYMafia
		mafiaBoss.Weapon = weaponMafia
	}))
}

func (c *Cutscene) Transition(player system.Live, gf system.Live, enemyManager *enemy.EnemyManager) {
	monterMan := enemyManager.Enemies[0]

	monterMan.Object.X = 600
	monterMan.Object.Y = 400
	monterMan.Weapon = nil

	gf.SetActive(false)
	c.AddAction(NewObjectMoveAction(monterMan, 1500, float32(monterMan.GetObject().Y), 4, "fb_walk_with_girl"))
	c.AddAction(NewObjectMoveAction(player, 1200, float32(player.GetObject().Y), 4, "walk"))
}

func (c *Cutscene) GfMonster(player system.Live, gf system.Live, enemyManager *enemy.EnemyManager, prs []*props.Prop) {
	monsterGf := enemyManager.Enemies[0]

	monsterGf.Object.X = 5000
	monsterGf.Object.Y = 5000

	gf.SetActive(false)

	c.AddAction(NewWaitAction(1.0))

	c.AddAction(NewCallbackAction(func() {
		audio.PlayGlassBreakingSound()
		c.DrawBlackScreen = true
	}))

	c.AddAction(NewWaitAction(1.0))

	c.AddAction(NewCallbackAction(func() {
		c.DrawBlackScreen = false
		prs[0].HandleKick(nil, player.GetObject())
		monsterGf.Object.Flipped = true
		monsterGf.Object.X = 1100
		monsterGf.Object.Y = 500
	}))
	c.AddAction(NewWaitAction(1.0))
}
