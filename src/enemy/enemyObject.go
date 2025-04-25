package enemy

import (
	"math/rand"
	"otaviocosta2110/vincitorrado/src/physics"
	"otaviocosta2110/vincitorrado/src/screen"
	"otaviocosta2110/vincitorrado/src/sprites"
	"otaviocosta2110/vincitorrado/src/system"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	animationDelay int32 = 300
)

type Enemy struct {
	system.LiveObject
	LastAttackTime time.Time
	HitCount       int32
	LastHitTime    time.Time
	IsStunned      bool
	IsActive       bool
	StunEndTime    time.Time
	Layer          int
}

func (e *Enemy) GetObject() system.Object {
	return e.Object
}

func (e *Enemy) SetObject(obj system.Object) {
	e.Object = obj
}

func NewEnemy(x, y, speed, width, height, scale int32, sprite sprites.Sprite) *Enemy {
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
			},
			MaxHealth: 5,
			Health:    5,
			Speed:     speed,
			Flipped:   false,
		},
		IsActive: false,
		Layer:  0,
	}
}

func (e *Enemy) Draw() {
	var width float32 = float32(e.Object.Sprite.SpriteWidth)
	if e.Flipped {
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
}

func (e *Enemy) CheckAtk(player system.Object) bool {
	punchX := e.Object.X
	punchY := e.Object.Y - e.Object.Height/3

	punchWidth := e.Object.Width / 2
	punchHeight := e.Object.Height / 2

	if e.Flipped {
		punchX -= punchWidth + punchWidth
	} else {
		punchX += punchWidth
	}

	punchObject := system.Object{
		X:      punchX,
		Y:      punchY,
		Width:  punchWidth,
		Height: punchHeight,
	}

	attackCooldown := int64(2000)
	// wind up time eh tipo o tempo que o boneco precisa esperar pra atacar
	// tipo, ele vai parar na frente do player e esperar 0.5 seg pra atacar de fato
	windUpTime := int64(500)

	timeSinceLastAttack := time.Since(e.Object.LastAttackTime).Milliseconds()

	if timeSinceLastAttack < windUpTime {
		return false
	}

	if physics.CheckCollision(punchObject, player) {
		if timeSinceLastAttack >= attackCooldown {
			e.Object.LastAttackTime = time.Now()

			framex := rand.Intn(2)
			e.Object.UpdateAnimation(50, []int{framex}, []int{1})

			return true
		}
	}

	e.Object.UpdateAnimation(300, []int{0, 1}, []int{0, 0})
	return false
}

func (e *Enemy) Update(p system.Player, screen screen.Screen) {
	if e.Object.Destroyed {
		e.Object.FrameX = 0
		e.Object.FrameY = 3
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

		if e.Object.KnockbackX == 0 || e.Object.KnockbackY == 0 {
			*e = MoveEnemyTowardPlayer(p, *e, screen)
		}
	} else {
		// todo: animacao pra quando o inimigo estiver stunado
	}
}

// ele devia na vdd soh mandar pra tras qnd levasse tipo 3 hit
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

func (e *Enemy) TakeDamage(damage int32, pX int32, pY int32) {
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

	if e.HitCount >= 3 {
		e.Object.UpdateAnimation(100, []int{1, 1}, []int{2, 2})
		e.setKnockback(pX)
		e.HitCount = 0
		e.IsStunned = true
		e.StunEndTime = time.Now().Add(700 * time.Millisecond)
	} else {
		e.Object.UpdateAnimation(100, []int{0, 0}, []int{2, 2})

	}

	e.LastDamageTaken = time.Now()
}

func (e *Enemy) TakeDamageFromBox(box system.Object) {
	damage := int32(1)
	e.TakeDamage(damage, box.X, box.Y)

	knockbackStrength := int32(15)
	if e.Object.X < box.X {
		e.Object.KnockbackX = -knockbackStrength
	} else {
		e.Object.KnockbackX = knockbackStrength
	}
	e.Object.KnockbackY = -knockbackStrength / 2
}
