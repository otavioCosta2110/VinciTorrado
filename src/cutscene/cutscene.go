package cutscene

import (
	"math"
	"otaviocosta2110/vincitorrado/src/system"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CutsceneAction interface {
	Update() bool
}

type Cutscene struct {
	actions []CutsceneAction
	current int
	playing bool
}

func NewCutscene() *Cutscene {
	return &Cutscene{
		actions: make([]CutsceneAction, 0),
		current: 0,
		playing: false,
	}
}

func (c *Cutscene) AddAction(action CutsceneAction) {
	c.actions = append(c.actions, action)
}

func (c *Cutscene) Start() {
	c.playing = true
	c.current = 0
}
func (c *Cutscene) Stop() {
	c.playing = false
}

func (c *Cutscene) Update() bool {
	if !c.playing || c.current >= len(c.actions) {
		return true
	}

	if c.current >= len(c.actions) {
		c.Stop()
		return true
	}

	if c.actions[c.current].Update() {
		c.current++

		if c.current >= len(c.actions) {
			c.Stop()
			return true
		}
	}
	return false
}

func (c *Cutscene) IsPlaying() bool {
	return c.playing
}

type CameraMoveAction struct {
	targetX, targetY float32
	duration         float32
	elapsed          float32
	camera           *rl.Camera2D
}

func NewCameraMoveAction(camera *rl.Camera2D, targetX, targetY, duration float32) *CameraMoveAction {
	return &CameraMoveAction{
		camera:   camera,
		targetX:  targetX,
		targetY:  targetY,
		duration: duration,
	}
}

func (a *CameraMoveAction) Update() bool {
	a.elapsed += rl.GetFrameTime()

	progress := min(a.elapsed/a.duration, 1)

	startX := a.camera.Target.X
	startY := a.camera.Target.Y

	a.camera.Target.X = startX + (a.targetX-startX)*progress
	a.camera.Target.Y = startY + (a.targetY-startY)*progress

	return a.elapsed >= a.duration
}

type WaitAction struct {
	duration float32
	elapsed  float32
}

func NewWaitAction(duration float32) *WaitAction {
	return &WaitAction{duration: duration}
}

func (a *WaitAction) Update() bool {
	a.elapsed += rl.GetFrameTime()
	return a.elapsed >= a.duration
}

type PlayerMoveAction struct {
	object    system.Live
	targetX   float32
	targetY   float32
	speed     float32
	completed bool
}

func NewObjectMoveAction(o system.Live, targetX, targetY, speed float32) *PlayerMoveAction {
	return &PlayerMoveAction{
		object:    o,
		targetX:   targetX,
		targetY:   targetY,
		speed:     speed,
		completed: false,
	}
}

func (a *PlayerMoveAction) Update() bool {
	if a.completed {
		return true
	}

	dx := a.targetX - float32(a.object.GetObject().X)
	dy := a.targetY - float32(a.object.GetObject().Y)
	distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))

	if distance < 5 {
		a.completed = true
		return true
	}

	dx = dx / distance * a.speed
	dy = dy / distance * a.speed

	a.object.SetObject(
		system.Object{
			X:              a.object.GetObject().X + int32(dx),
			Y:              a.object.GetObject().Y + int32(dy),
			Width:          a.object.GetObject().Width,
			Height:         a.object.GetObject().Height,
			KnockbackX:     a.object.GetObject().KnockbackX,
			KnockbackY:     a.object.GetObject().KnockbackY,
			FrameX:         a.object.GetObject().FrameX,
			FrameY:         a.object.GetObject().FrameY,
			LastFrameTime:  a.object.GetObject().LastFrameTime,
			LastAttackTime: a.object.GetObject().LastAttackTime,
			Sprite:         a.object.GetObject().Sprite,
			Scale:          a.object.GetObject().Scale,
			Destroyed:      a.object.GetObject().Destroyed,
			IsKicking:      a.object.GetObject().IsKicking,
			Flipped:        a.object.GetObject().Flipped,
		},
	)
	if math.Abs(float64(dx)) > math.Abs(float64(dy)) {
		a.object.UpdateAnimation("walk")
		if dx > 0 {
			a.object.SetObject(
				system.Object{
					X:              a.object.GetObject().X + int32(dx),
					Y:              a.object.GetObject().Y + int32(dy),
					Width:          a.object.GetObject().Width,
					Height:         a.object.GetObject().Height,
					KnockbackX:     a.object.GetObject().KnockbackX,
					KnockbackY:     a.object.GetObject().KnockbackY,
					FrameX:         a.object.GetObject().FrameX,
					FrameY:         a.object.GetObject().FrameY,
					LastFrameTime:  a.object.GetObject().LastFrameTime,
					LastAttackTime: a.object.GetObject().LastAttackTime,
					Sprite:         a.object.GetObject().Sprite,
					Scale:          a.object.GetObject().Scale,
					Destroyed:      a.object.GetObject().Destroyed,
					IsKicking:      a.object.GetObject().IsKicking,
					Flipped:        false,
				},
			)
		} else {
			a.object.SetObject(
				system.Object{
					X:              a.object.GetObject().X + int32(dx),
					Y:              a.object.GetObject().Y + int32(dy),
					Width:          a.object.GetObject().Width,
					Height:         a.object.GetObject().Height,
					KnockbackX:     a.object.GetObject().KnockbackX,
					KnockbackY:     a.object.GetObject().KnockbackY,
					FrameX:         a.object.GetObject().FrameX,
					FrameY:         a.object.GetObject().FrameY,
					LastFrameTime:  a.object.GetObject().LastFrameTime,
					LastAttackTime: a.object.GetObject().LastAttackTime,
					Sprite:         a.object.GetObject().Sprite,
					Scale:          a.object.GetObject().Scale,
					Destroyed:      a.object.GetObject().Destroyed,
					IsKicking:      a.object.GetObject().IsKicking,
					Flipped:        true,
				},
			)
		}
	}
	return false
}
type CallbackAction struct {
    callback func()
}

func NewCallbackAction(callback func()) *CallbackAction {
    return &CallbackAction{callback: callback}
}

func (a *CallbackAction) Update() bool {
    a.callback()
    return true 
}
