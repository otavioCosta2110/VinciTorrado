package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"otaviocosta2110/vincitorrado/src/player"
	"otaviocosta2110/vincitorrado/src/sprites"
)

type MenuState int

const (
	MenuMain MenuState = iota
	MenuEquipment
	MenuStats
)

type EquipmentSlot struct {
	Name     string
	Rect     rl.Rectangle
	IsActive bool
	IconPos  rl.Rectangle
}

type Menu struct {
	State           MenuState
	IsVisible       bool
	PlayerSprite    *sprites.Sprite
	PlayerReference *player.Player
	EquipmentSlots  []EquipmentSlot
	SelectedSlot    int
	Columns         int
	IconSheet       rl.Texture2D
}

func NewMenu(player *player.Player, sprite *sprites.Sprite) *Menu {
	menu := &Menu{
		State:           MenuMain,
		IsVisible:       false,
		PlayerSprite:    sprite,
		PlayerReference: player,
		SelectedSlot:    0,
		Columns:         3,
		IconSheet:       rl.LoadTexture("assets/ui/equipamentos.png"),
	}

	equipmentTypes := []struct {
		Name string
		X    int32
		Y    int32
	}{
		{"Nada", 0, 0},
		{"Turbante", 1, 0},
	}

	slotWidth := float32(120)
	slotHeight := float32(120)
	spacing := float32(50)
	startX := float32(rl.GetScreenWidth())/2 - (float32(menu.Columns)*slotWidth)/2 + 50
	startY := float32(rl.GetScreenHeight()/5)

	for i, eq := range equipmentTypes {
		row := i / menu.Columns
		col := i % menu.Columns
		
		menu.EquipmentSlots = append(menu.EquipmentSlots, EquipmentSlot{
			Name: eq.Name,
			Rect: rl.NewRectangle(
				startX + float32(col)*(slotWidth+spacing),
				startY + float32(row)*(slotHeight+spacing),
				slotWidth,
				slotHeight,
			),
			IsActive: false,
			IconPos: rl.NewRectangle(
				float32(eq.X)*32,
				float32(eq.Y)*32,
				32,
				32,
			),
		})
	}

	return menu
}

func (m *Menu) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		m.IsVisible = !m.IsVisible
	}

	if !m.IsVisible {
		return
	}

	prevSelected := m.SelectedSlot
	if rl.IsKeyPressed(rl.KeyRight) {
		m.SelectedSlot++
		if m.SelectedSlot >= len(m.EquipmentSlots) {
			m.SelectedSlot = 0
		}
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		m.SelectedSlot--
		if m.SelectedSlot < 0 {
			m.SelectedSlot = len(m.EquipmentSlots) - 1
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		m.SelectedSlot += m.Columns
		if m.SelectedSlot >= len(m.EquipmentSlots) {
			m.SelectedSlot %= len(m.EquipmentSlots)
		}
	} else if rl.IsKeyPressed(rl.KeyUp) {
		m.SelectedSlot -= m.Columns
		if m.SelectedSlot < 0 {
			m.SelectedSlot += len(m.EquipmentSlots)
			if m.SelectedSlot < 0 {
				m.SelectedSlot = 0
			}
		}
	}

	if prevSelected != m.SelectedSlot && m.SelectedSlot >= 0 {
	}

	if rl.IsKeyPressed(rl.KeyEnter) && m.SelectedSlot >= 0 {
		println("Selected:", m.EquipmentSlots[m.SelectedSlot].Name)
	}
}

func (m *Menu) Draw() {
	if !m.IsVisible {
		return
	}

	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Fade(rl.Black, 0.5))

	menuWidth := float32(rl.GetScreenWidth()) * 0.8
	menuHeight := float32(rl.GetScreenHeight()) * 0.8
	menuX := (float32(rl.GetScreenWidth()) - menuWidth) / 2
	menuY := (float32(rl.GetScreenHeight()) - menuHeight) / 2

	rl.DrawRectangleRounded(
		rl.NewRectangle(menuX, menuY, menuWidth, menuHeight),
		0.05, 10, rl.DarkGray,
	)

	playerPreviewX := menuX 
	playerPreviewY := menuY * 2
	rl.DrawTexturePro(
		m.PlayerSprite.Texture,
		rl.NewRectangle(0, 0, float32(m.PlayerSprite.SpriteWidth), float32(m.PlayerSprite.SpriteHeight)),
		rl.NewRectangle(
			playerPreviewX,
			playerPreviewY,
			float32(m.PlayerSprite.SpriteWidth)*float32(m.PlayerReference.Object.Scale * 3),
			float32(m.PlayerSprite.SpriteHeight)*float32(m.PlayerReference.Object.Scale * 3),
		),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)

	for i, slot := range m.EquipmentSlots {
		color := rl.Gray
		textColor := rl.White
		if i == m.SelectedSlot {
			color = rl.White
			textColor = rl.White
		}

		rl.DrawRectangleRounded(slot.Rect, 0.1, 5, color)
		rl.DrawTexturePro(
			m.IconSheet,
			slot.IconPos,
			rl.NewRectangle(
				slot.Rect.X + slot.IconPos.Width/2,
				slot.Rect.Y + slot.IconPos.Width/2,
				85,
				85,
			),
			rl.NewVector2(0, 0),
			0,
			rl.White,
		)
		rl.DrawText(slot.Name, int32(slot.Rect.X+2), int32(slot.Rect.Y-20), 20, textColor)
	}
}
