package player

import (
	"otaviocosta2110/vincitorrado/src/enemy"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	// "slices"
)

const (
	invencibilityDuration = 1000
)

func (p *Player) Update(em *enemy.EnemyManager, screen screen.Screen) {
	p.adjustKnockbackToScreenBounds()
	physics.TakeKnockback(&p.Object)
	if p.Object.KnockbackX == 0 || p.Object.KnockbackY == 0 {
		p.CheckMovement(screen)
	}

	p.HandleAttackInput()
	// p.UpdateProjectiles(em)

	for _, enemy := range em.Enemies {
		enemyObject := enemy.GetObject()
		if p.CheckAtk(enemyObject) {
			enemy.TakeDamage(p.Damage, p.Object)
		}
	}
}

func (p *Player) Draw() {
	color := rl.White

	if p.isInvincible(invencibilityDuration) {
		if time.Since(p.LastDamageTaken).Milliseconds()/100%2 == 0 {
			color = rl.Fade(rl.White, 0.3)
		}
	}

	var playerWidth float32 = float32(p.Object.Sprite.SpriteWidth)
	if p.Object.Flipped {
		playerWidth = -playerWidth
	}

	sourceRec := rl.NewRectangle(
		float32(p.Object.FrameX)*float32(p.Object.Sprite.SpriteWidth),
		float32(p.Object.FrameY)*float32(p.Object.Sprite.SpriteHeight),
		playerWidth,
		float32(p.Object.Sprite.SpriteHeight),
	)

	destinationRec := rl.NewRectangle(
		float32(p.Object.X),
		float32(p.Object.Y),
		float32(p.Object.Sprite.SpriteWidth)*float32(p.Object.Scale),
		float32(p.Object.Sprite.SpriteHeight)*float32(p.Object.Scale),
	)

	origin := rl.NewVector2(
		destinationRec.Width/2,
		destinationRec.Height/2,
	)

	rl.DrawTexturePro(
		p.Object.Sprite.Texture,
		sourceRec,
		destinationRec,
		origin,
		0.0,
		color,
	)

	if p.Equipped != nil && p.Equipped.IsEquipped && p.HatSprite.Texture.ID != 0 {
		rl.DrawTexturePro(
			p.HatSprite.Texture,
			sourceRec,
			destinationRec,
			origin,
			0.0,
			color,
		)
	}

	if p.Weapon != nil {
		p.Weapon.DrawEquipped(&p.Object)
	}
	// p.DrawProjectiles()
}

func (p *Player) TakeDamage(damage int32, eObj system.Object) {
	if !p.isInvincible(invencibilityDuration) {
		p.UpdateAnimation("hit")
		p.Health -= damage
		p.LastDamageTaken = time.Now()

		p.Object.SetKnockback(eObj)
		if p.Health < 1 {
			system.GameOverFlag = true
		}
	}
}

func (p *Player) adjustKnockbackToScreenBounds() {
	screenLeft := int32(p.Screen.Camera.Target.X - float32(p.Screen.Width)/2 + float32(p.Object.Width/2))
	screenRight := int32(p.Screen.Camera.Target.X + float32(p.Screen.Width)/2 - float32(p.Object.Width/2))

	screenTop := p.Object.Height - p.Object.Y + (p.Screen.ScenaryHeight + p.Object.Height)
	screenBottom := p.Screen.Height - (p.Object.Height)/2

	newX := p.Object.X + p.Object.KnockbackX
	newY := p.Object.Y + p.Object.KnockbackY

	if newX < screenLeft {
		p.Object.KnockbackX = screenLeft - p.Object.X
	} else if newX > screenRight {
		p.Object.KnockbackX = screenRight - p.Object.X
	}

	if newY < screenTop {
		p.Object.KnockbackY = screenTop - p.Object.Y
	} else if newY > screenBottom {
		p.Object.KnockbackY = screenBottom - p.Object.Y
	}
}

func (p *Player) isInvincible(duration int) bool {
	return time.Since(p.LastDamageTaken).Milliseconds() <= int64(duration)
}

func (p *Player) UpdateAnimation(animationName string) {
	switch animationName {
	case "walk":
		p.runAnimation(300, []int{0, 1}, []int{0, 0})
	case "punch":
		p.runAnimation(50, []int{0, 1}, []int{1, 1})
	case "kick":
		p.runAnimation(50, []int{0}, []int{3})
	case "hit":
		p.runAnimation(100, []int{0, 1}, []int{2, 2})
	case "default":
		p.runAnimation(int(animationDelay), []int{0}, []int{0})
	}
}

func (p *Player) runAnimation(animationDelay int, framesX, framesY []int) {
	p.Object.UpdateAnimation(animationDelay, framesX, framesY)
	if p.Weapon != nil {
		p.Weapon.Object.UpdateAnimation(animationDelay, framesX, framesY)
	}
}

func (p *Player) SetObject(obj system.Object) {
	p.Object = obj
}

func (p *Player) IsActive() bool {
	return true
}
func (p *Player) SetActive(bool) {}

// func (p *Player) Shoot() {
// 	println(p.Weapon.Ammo)
// 	if p.Weapon == nil || !p.Weapon.IsGun || p.Weapon.Ammo <= 0 {
// 		return
// 	}
//
// 	direction := rl.Vector2{X: 1.0, Y: 0.0}
// 	if p.Object.Flipped {
// 		direction.X = -1.0
// 	}
//
// 	startX := float32(p.Object.X)
// 	startY := float32(p.Object.Y)
//
// 	projectile := p.Weapon.Shoot(startX, startY, direction)
// 	if projectile != nil {
// 		p.Projectiles = append(p.Projectiles, projectile)
// 	}
// }

// func (p *Player) UpdateProjectiles(em *enemy.EnemyManager) {
// 	for i := 0; i < len(p.Projectiles); {
// 		proj := p.Projectiles[i]
// 		proj.Update()
//
// 		hitEnemy := false
//
// 		for _, enemy := range em.Enemies {
// 			if proj.IsActive && !enemy.GetObject().Destroyed &&
// 				physics.CheckCollision(*proj.Object, enemy.GetObject()) {
// 				enemy.TakeDamage(proj.Damage, *proj.Object)
// 				proj.IsActive = false
// 				hitEnemy = true
// 				break
// 			}
// 		}
//
// 		if !proj.IsActive || hitEnemy {
// 			p.Projectiles = slices.Delete(p.Projectiles, i, i+1)
// 		} else {
// 			i++
// 		}
// 	}
// }
//
// func (p *Player) DrawProjectiles() {
// 	for _, proj := range p.Projectiles {
// 		proj.Draw()
// 	}
// }
