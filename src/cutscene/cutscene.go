package cutscene

import (
	"math"
	"otaviocosta2110/vincitorrado/src/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type CutsceneAction interface {
    Update() bool // Returns true when action is complete
    Draw()
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

func (c *Cutscene) Update() bool {
    if !c.playing || c.current >= len(c.actions) {
        return true // Cutscene is done
    }
    
    if c.actions[c.current].Update() {
        c.current++
    }
    
    return false
}

func (c *Cutscene) Draw() {
    if c.playing && c.current < len(c.actions) {
        c.actions[c.current].Draw()
    }
}

func (c *Cutscene) IsPlaying() bool {
    return c.playing
}

// Camera movement action
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
    
    progress := min(a.elapsed / a.duration, 1)
    
    startX := a.camera.Target.X
    startY := a.camera.Target.Y
    
    a.camera.Target.X = startX + (a.targetX-startX)*progress
    a.camera.Target.Y = startY + (a.targetY-startY)*progress
    
    return a.elapsed >= a.duration
}

func (a *CameraMoveAction) Draw() {
    // Camera actions don't need to draw anything
}

// Dialogue action
type DialogueAction struct {
    text      string
    duration  float32
    elapsed   float32
    textBox   rl.Rectangle
}

func NewDialogueAction(text string, duration float32, screenWidth, screenHeight int32) *DialogueAction {
    return &DialogueAction{
        text:     text,
        duration: duration,
        textBox: rl.NewRectangle(
            float32(screenWidth)*0.1,
            float32(screenHeight)*0.7,
            float32(screenWidth)*0.8,
            float32(screenHeight)*0.2,
        ),
    }
}

func (a *DialogueAction) Update() bool {
    a.elapsed += rl.GetFrameTime()
    return a.elapsed >= a.duration || rl.IsKeyPressed(rl.KeySpace)
}

func (a *DialogueAction) Draw() {
    // Draw text box background
    rl.DrawRectangleRec(a.textBox, rl.Fade(rl.Black, 0.7))
    
    // Draw text
    rl.DrawText(
        a.text,
        int32(a.textBox.X+20),
        int32(a.textBox.Y+20),
        20,
        rl.White,
    )
}

// Wait action (simple delay)
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


func (a *WaitAction) Draw() {}
// PlayerMoveAction makes the player character move during a cutscene
type PlayerMoveAction struct {
    player    *player.Player
    targetX   float32
    targetY   float32
    speed     float32
    completed bool
}

func NewPlayerMoveAction(p *player.Player, targetX, targetY, speed float32) *PlayerMoveAction {
    return &PlayerMoveAction{
        player:    p,
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
    
    // Calculate direction to target
    dx := a.targetX - float32(a.player.Object.X)
    dy := a.targetY - float32(a.player.Object.Y)
    distance := float32(math.Sqrt(float64(dx*dx + dy*dy)))
    
    // If we're close enough to the target, mark as completed
    if distance < 5 { // 5 pixels is close enough
        a.completed = true
        return true
    }
    
    // Normalize direction and apply speed
    dx = dx / distance * a.speed
    dy = dy / distance * a.speed
    
    // Update player position
    a.player.Object.X += int32(dx)
    a.player.Object.Y += int32(dy)
    
    // Update player animation based on direction
    if math.Abs(float64(dx)) > math.Abs(float64(dy)) {
        if dx > 0 {
            a.player.UpdatePlayerAnimation()
        } else {
            a.player.CurrentAnimation = a.player.Animations["walk_left"]
        }
    } else {
        if dy > 0 {
            a.player.CurrentAnimation = a.player.Animations["walk_down"]
        } else {
            a.player.CurrentAnimation = a.player.Animations["walk_up"]
        }
    }
    
    return false
}

func (a *PlayerMoveAction) Draw() {
    // Player movement doesn't need to draw anything extra
}
