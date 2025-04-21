package system

import (
	"otaviocosta2110/vincitorrado/src/sprites"
	"time"
)

type Object struct {
	X, Y           int32
	Width          int32
	Height         int32
	KnockbackX     int32
	KnockbackY     int32
	FrameX         int32
	FrameY         int32
	LastFrameTime  time.Time
	LastAttackTime time.Time
	Sprite         sprites.Sprite
	Scale          int32
	Destroyed      bool
}

func (obj *Object) UpdateAnimation(animationDelay int, framesX, framesY []int) {
	if time.Since(obj.LastFrameTime).Milliseconds() > int64(animationDelay) {
		currentIndex := -1
		for i := range framesX {
			if obj.FrameX == int32(framesX[i]) && obj.FrameY == int32(framesY[i]) {
				currentIndex = i
				break
			}
		}

		if currentIndex == -1 {
			obj.FrameX = int32(framesX[0])
			obj.FrameY = int32(framesY[0])
		} else {
			nextIndex := (currentIndex + 1) % len(framesX)
			obj.FrameX = int32(framesX[nextIndex])
			obj.FrameY = int32(framesY[nextIndex])
		}

		obj.LastFrameTime = time.Now()
	}
}

func (obj *Object) SetKnockback(attackingObj Object) {
	knockbackStrengthX := int32(10)
	knockbackStrengthY := int32(10)

	if obj.X < attackingObj.X {
		obj.KnockbackX = -knockbackStrengthX
	} else {
		obj.KnockbackX = knockbackStrengthX
	}

	if obj.Y < attackingObj.Y {
		obj.KnockbackY = -knockbackStrengthY
	} else {
		obj.KnockbackY = knockbackStrengthY
	}

}
