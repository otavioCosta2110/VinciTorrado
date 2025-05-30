package enemy

import (
	"math/rand"
	"otaviocosta2110/vincitorrado/src/audio"
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/props"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"
	"otaviocosta2110/vincitorrado/src/weapon"
	"slices"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	animationDelay int32 = 300
)

type Enemy struct {
	system.LiveObject
	Activate_pos_X int32
	Activate_pos_Y int32
	LastAttackTime time.Time
	HitCount       int32
	LastHitTime    time.Time
	IsStunned      bool
	Active         bool
	StunEndTime    time.Time
	Layer          int
	CanMove        bool
	WindUpTime     int64
	isSpawning     bool
	EnemyType      string
	Drop           *equipment.Equipment
	DropCollected  bool
	Weapon         *weapon.Weapon
	AttackCooldown int64
	IsCharging     bool
	Projectiles    []*weapon.Projectile
	LastShotTime   time.Time
}

func (e *Enemy) GetObject() system.Object {
	return e.Object
}

func (e *Enemy) SetObject(obj system.Object) {
	e.Object = obj
}

func NewEnemy(x, y, aX, aY, speed, width, height, scale int32, sprite sprites.Sprite, windUpTime int64, enemyType string, attackCooldown int64, drops *equipment.Equipment, weapon *weapon.Weapon) *Enemy {
	return &Enemy{
		LiveObject: system.LiveObject{
			Object: system.Object{
				X:              x,
				Y:              y,
				Width:          width * scale / 2,
				Height:         height * scale,
				KnockbackX:     0,
				KnockbackY:     0,
				FrameY:         0,
				FrameX:         0,
				LastFrameTime:  time.Now(),
				LastAttackTime: time.Now(),
				Sprite: sprites.Sprite{
					SpriteWidth:  width,
					SpriteHeight: height,
					Texture:      sprite.Texture,
				},
				Scale:     scale,
				Destroyed: false,
				Flipped:   false,
			},
			MaxHealth: 5,
			Health:    5,
			Speed:     speed,
		},
		Activate_pos_X: aX,
		Activate_pos_Y: aY,
		Active:         false,
		Layer:          0,
		CanMove:        true,
		WindUpTime:     windUpTime,
		isSpawning:     true,
		EnemyType:      enemyType,
		Drop:           drops,
		Weapon:         weapon,
		AttackCooldown: attackCooldown,
		IsCharging:     false,
		LastShotTime:   time.Now(),
	}
}

func (e *Enemy) Draw() {
	if e.Object.Destroyed {
		if e.Drop != nil && !e.DropCollected {
			e.Drop.DrawAnimated(&e.Object)
		}
		if e.Weapon != nil && e.Weapon.IsDropped {
			e.Weapon.DrawAnimated()
		}

	}

	var width float32 = float32(e.Object.Sprite.SpriteWidth)
	if e.Object.Flipped {
		width = -float32(width)
	}

	sourceRec := rl.NewRectangle(
		float32(e.Object.FrameX)*float32(e.Object.Sprite.SpriteWidth),
		float32(e.Object.FrameY)*float32(e.Object.Sprite.SpriteWidth),
		width,
		float32(e.Object.Sprite.SpriteHeight),
	)

	destinationRec := rl.NewRectangle(
		float32(e.Object.X),
		float32(e.Object.Y),
		float32(e.Object.Sprite.SpriteWidth)*float32(e.Object.Scale),
		float32(e.Object.Sprite.SpriteHeight)*float32(e.Object.Scale),
	)

	origin := rl.NewVector2(
		destinationRec.Width/2,
		destinationRec.Height/2,
	)

	rl.DrawTexturePro(e.Object.Sprite.Texture, sourceRec, destinationRec, origin, 0.0, rl.White)
	if e.Object.Destroyed && e.Drop != nil {
		e.Drop.DrawAnimated(&e.Object)
	}
	if e.Weapon != nil && !e.Weapon.IsDropped {
		e.Weapon.DrawEquipped(&e.Object)
	}

	if e.Weapon != nil && e.Weapon.IsGun {
		e.DrawProjectiles()
	}
}

func (e *Enemy) Shoot() {
	if e.Weapon == nil || !e.Weapon.IsGun || e.Weapon.Ammo <= 0 {
		return
	}

	direction := rl.Vector2{X: 1.0, Y: 0.0}
	if e.Object.Flipped {
		direction.X = -1.0
	}

	startX := float32(e.Object.X)
	startY := float32(e.Object.Y)

	projectile := e.Weapon.Shoot(startX, startY, direction)
	if projectile != nil {
		e.Projectiles = append(e.Projectiles, projectile)
	}
}

func (e *Enemy) CheckAtk(player system.Object) bool {
	currentTime := time.Now()
	timeSinceLastAttack := time.Since(e.Object.LastAttackTime).Milliseconds()
	timeSinceLastShot := time.Since(e.LastShotTime).Seconds()

	if timeSinceLastAttack < e.AttackCooldown {
		e.CanMove = false
		return false
	}
	e.CanMove = true

	if e.Weapon != nil && e.Weapon.IsGun && timeSinceLastShot >= 4.0 {
		e.Shoot()
		e.LastShotTime = currentTime
	}

	punchX := e.Object.X
	punchY := e.Object.Y - e.Object.Height/3
	punchWidth := e.Object.Width / 2
	punchHeight := e.Object.Height / 2

	if e.Object.Flipped {
		punchX -= (punchWidth + punchWidth) - 35
	} else {
		punchX += punchWidth
	}

	punchObject := system.Object{
		X:      punchX,
		Y:      punchY,
		Width:  punchWidth,
		Height: punchHeight,
	}

	if e.EnemyType == "full_belly" {
		playerObj_margin := &system.Object{
			X:              player.X,
			Y:              player.Y,
			Width:          player.Width + e.Weapon.HitboxX,
			Height:         player.Height + e.Weapon.HitboxY,
			LastFrameTime:  player.LastFrameTime,
			LastAttackTime: player.LastAttackTime,
			Scale:          4,
		}
		if physics.CheckCollision(punchObject, *playerObj_margin) {
			if !e.IsCharging && timeSinceLastAttack >= e.AttackCooldown {
				e.IsCharging = true
				audio.PlayFullBellyPrepare()
				e.CanMove = false
				e.UpdateAnimation("fb_charge")
				e.Object.LastAttackTime = currentTime
				return false
			}

			if e.IsCharging && timeSinceLastAttack >= e.WindUpTime {
				e.IsCharging = false
				e.UpdateAnimation("fb_attack")
				e.Object.LastAttackTime = currentTime
				audio.PlayFullBellyAttack()
				return true
			}
		} else {
			e.IsCharging = false
			e.CanMove = true
		}
	}

	if physics.CheckCollision(punchObject, player) && timeSinceLastAttack >= e.AttackCooldown {
		framex := rand.Intn(2)
		e.Object.FrameX = int32(framex)
		e.UpdateAnimation("punch")
		e.Object.LastAttackTime = time.Now()
		audio.PlayPunch()
		return true
	}

	if e.EnemyType == "full_belly" && e.IsCharging {
		e.UpdateAnimation("fb_charge")
	} else {
		e.UpdateAnimation("walk")
	}
	return false
}
func (e *Enemy) Update(p system.Player, screen screen.Screen, prps []*props.Prop) {
	if e.isSpawning {
		e.isSpawning = false
	}
	if e.Object.Destroyed {
		e.Object.FrameX = 0
		e.Object.FrameY = 3
		if e.EnemyType != "full_belly" {
			e.DropWeapon()
		} else {
			e.Weapon = nil
		}
		return
	}

	if e.IsStunned && time.Now().After(e.StunEndTime) {
		e.IsStunned = false
	}

	physics.TakeKnockback(&e.Object)

	if !e.IsStunned {
		if e.CheckAtk(p.GetObject()) {
			p.TakeDamage(e.Damage, e.Object)
			return
		}

		if (e.Object.KnockbackX == 0 || e.Object.KnockbackY == 0) && !e.IsCharging {
			*e = MoveEnemyTowardPlayer(p, *e, screen)
		}
	}
	e.UpdateProjectiles(p, prps)
}

func (e *Enemy) setKnockback(pX int32) {
	knockbackStrengthX := int32(20)
	knockbackStrengthY := int32(0)

	if e.Object.X < pX {
		e.Object.KnockbackX = -knockbackStrengthX
	} else {
		e.Object.KnockbackX = knockbackStrengthX
	}

	e.Object.KnockbackY = knockbackStrengthY
}

func (e *Enemy) TakeDamage(damage int32, obj system.Object) {
	if e.Health <= 0 {
		e.Object.Destroyed = true
		e.Layer = -1
		return
	}

	hitWindow := time.Millisecond * 500
	if time.Since(e.LastHitTime) > hitWindow {
		e.HitCount = 0
	}

	e.Health -= damage
	e.LastHitTime = time.Now()
	e.HitCount++

	if e.EnemyType != "full_belly" {
		if e.HitCount >= 3 {
			e.UpdateAnimation("fb_hit")
			e.setKnockback(obj.X)
			e.HitCount = 0
			e.IsStunned = true
			e.StunEndTime = time.Now().Add(700 * time.Millisecond)
		} else {
			e.UpdateAnimation("hit")
		}
	}

	e.LastDamageTaken = time.Now()
}

func (e *Enemy) TakeDamageFromBox(box system.Object) {
	damage := int32(1)
	e.TakeDamage(damage, e.Object)

	knockbackStrength := int32(15)
	if e.Object.X < box.X {
		e.Object.KnockbackX = -knockbackStrength
	} else {
		e.Object.KnockbackX = knockbackStrength
	}
	e.Object.KnockbackY = -knockbackStrength / 2
}

func (e *Enemy) UpdateAnimation(animationName string) {
	switch animationName {
	case "walk":
		e.runAnimation(300, []int{0, 1}, []int{0, 0})
	case "punch":
		e.runAnimation(50, []int{0, 1}, []int{1, 1})
	case "kick":
		e.runAnimation(50, []int{0}, []int{3})
	case "hit":
		e.runAnimation(100, []int{0, 1}, []int{2, 2})
	case "fb_charge":
		e.runAnimation(100, []int{0}, []int{1})
	case "fb_attack":
		e.runAnimation(0, []int{1}, []int{1})
	case "fb_hit":
		e.runAnimation(100, []int{1, 1}, []int{2, 2})
	case "fb_walk_with_girl":
		e.runAnimation(300, []int{0, 1}, []int{4, 4})
	case "default":
		e.runAnimation(int(animationDelay), []int{0}, []int{0})
	}
}

func (e *Enemy) runAnimation(animationDelay int, framesX, framesY []int) {
	e.Object.UpdateAnimation(animationDelay, framesX, framesY)
	if e.Weapon != nil {
		e.Weapon.Object.UpdateAnimation(animationDelay, framesX, framesY)
	}
}

func (e *Enemy) DropWeapon() {
	if e.Weapon != nil {

		e.Weapon.Object.X = e.Object.X + 20
		e.Weapon.Object.Y = e.Object.Y - 20
		e.Weapon.IsDropped = true
		e.Weapon.IsEquipped = false
		e.Weapon.Object.FrameX = 0
		e.Weapon.Object.FrameY = 0
	}
}
func (e *Enemy) GetDropCollisionBox() system.Object {
	if e.Drop == nil {
		return system.Object{}
	}
	dropWidth := int32(32 * e.Object.Scale)
	dropHeight := int32(32 * e.Object.Scale)
	dropY := e.Object.Y - 20
	return system.Object{
		X:      e.Object.X,
		Y:      dropY,
		Width:  dropWidth / 2,
		Height: dropHeight / 2,
	}
}
func (e *Enemy) IsActive() bool {
	return e.Active
}
func (p *Enemy) SetActive(bool) {}

func (e *Enemy) UpdateProjectiles(p system.Player, prs []*props.Prop) {
	for i := 0; i < len(e.Projectiles); {
		proj := e.Projectiles[i]
		proj.Update()

		hitPlayer := false

		if proj.IsActive && !p.GetObject().Destroyed &&
			physics.CheckCollision(*proj.Object, p.GetObject()) {
			p.TakeDamage(proj.Damage, *proj.Object)
			proj.IsActive = false
			hitPlayer = true
			break
		}

		for _, prop := range prs {
			if prop.Kicked {
				propProjHitbox := system.Object{
					X:      prop.GetObject().X,
					Y:      prop.GetObject().Y + 50,
					Width:  prop.GetObject().Width + 200,
					Height: prop.GetObject().Height + 50,
				}
				if physics.CheckCollision(*proj.Object, propProjHitbox) {
					audio.PlayBulletHittingTableSound()
					proj.IsActive = false
				}
			}
		}

		if !proj.IsActive || hitPlayer {
			e.Projectiles = slices.Delete(e.Projectiles, i, i+1)
		} else {
			i++
		}
	}
}

func (e *Enemy) DrawProjectiles() {
	for _, proj := range e.Projectiles {
		proj.Draw()
	}
}
