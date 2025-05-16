package cutscene

import (
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/system"
)

func (c *Cutscene) IntroCutscenes(player system.Live, enemyManager *enemy.EnemyManager) {
	fullBelly:=enemyManager.Enemies[len(enemyManager.Enemies)-1]

	posXPanca := fullBelly.Object.X
	posYPanca := fullBelly.Object.Y
	fullBelly.Object.X = 1300
	fullBelly.Object.Y = 400

	c.AddAction(NewObjectMoveAction(fullBelly, 500, float32(fullBelly.GetObject().Y), 4))
	c.AddAction(NewObjectMoveAction(player, 200, float32(player.GetObject().Y), 4))

	c.AddAction(NewCallbackAction(func() {
		fullBelly.Object.X = posXPanca
		fullBelly.Object.Y = posYPanca
	}))
}
