package ui

import (
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/player"
	"otaviocosta2110/vincitorrado/src/sprites"

	rl "github.com/gen2brain/raylib-go/raylib"
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
	Item     *equipment.Equipment 
	IsEmpty  bool                
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
	SlotWidth       float32 
	SlotHeight      float32
	SlotSpacing     float32
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
		SlotWidth:       120,
		SlotHeight:      120,
		SlotSpacing:     50,
	}

	menu.initEquipmentSlots()
	return menu
}

func (m *Menu) initEquipmentSlots() {
	startX := float32(rl.GetScreenWidth())/2 - (float32(m.Columns)*m.SlotWidth)/2 + 50
	startY := float32(rl.GetScreenHeight() / 5)

	m.EquipmentSlots = make([]EquipmentSlot, 0, 9) 

	for i := range 9 {
		row := i / m.Columns
		col := i % m.Columns

		m.EquipmentSlots = append(m.EquipmentSlots, EquipmentSlot{
			Name: "Empty",
			Rect: rl.NewRectangle(
				startX+float32(col)*(m.SlotWidth+m.SlotSpacing),
				startY+float32(row)*(m.SlotHeight+m.SlotSpacing),
				m.SlotWidth,
				m.SlotHeight,
			),
			IsActive: false,
			IconPos:  rl.NewRectangle(0, 0, 32, 32), 
			Item:     nil,
			IsEmpty:  true,
		})
	}
}

func (m *Menu) AddToMenu(item *equipment.Equipment) bool {
	for i := range m.EquipmentSlots {
		if m.EquipmentSlots[i].IsEmpty {
			m.fillSlot(i, item)
			return true
		}
	}

	return m.addNewSlot(item)
}

func (m *Menu) fillSlot(index int, item *equipment.Equipment) {
	m.EquipmentSlots[index].Name = item.Name
	m.EquipmentSlots[index].IconPos = m.getItemIconPos(item)
	m.EquipmentSlots[index].Item = item
	m.EquipmentSlots[index].IsEmpty = false
}

func (m *Menu) addNewSlot(item *equipment.Equipment) bool {
	startX := float32(rl.GetScreenWidth())/2 - (float32(m.Columns)*m.SlotWidth)/2 + 50
	startY := float32(rl.GetScreenHeight() / 5)

	row := len(m.EquipmentSlots) / m.Columns
	col := len(m.EquipmentSlots) % m.Columns

	newSlot := EquipmentSlot{
		Name: item.Name,
		Rect: rl.NewRectangle(
			startX+float32(col)*(m.SlotWidth+m.SlotSpacing),
			startY+float32(row)*(m.SlotHeight+m.SlotSpacing),
			m.SlotWidth,
			m.SlotHeight,
		),
		IsActive: false,
		IconPos:  m.getItemIconPos(item),
		Item:     item,
		IsEmpty:  false,
	}

	m.EquipmentSlots = append(m.EquipmentSlots, newSlot)
	return true
}

func (m *Menu) RemoveFromMenu(index int) {
	if index >= 0 && index < len(m.EquipmentSlots) {
		m.EquipmentSlots[index].Name = "Empty"
		m.EquipmentSlots[index].IconPos = rl.NewRectangle(0, 0, 32, 32)
		m.EquipmentSlots[index].Item = nil
		m.EquipmentSlots[index].IsEmpty = true
	}
}

func (m *Menu) getItemIconPos(item *equipment.Equipment) rl.Rectangle {
	switch item.Name {
	case "Turbante":
		return rl.NewRectangle(32, 32, 32, 32)
	default:
		return rl.NewRectangle(0, 0, 32, 32)
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
	sourceRec := rl.NewRectangle(0, 0, float32(m.PlayerSprite.SpriteWidth), float32(m.PlayerSprite.SpriteHeight))

	destinationRec := rl.NewRectangle(
		playerPreviewX,
		playerPreviewY,
		float32(m.PlayerSprite.SpriteWidth)*float32(m.PlayerReference.Object.Scale*3),
		float32(m.PlayerSprite.SpriteHeight)*float32(m.PlayerReference.Object.Scale*3),
	)

	origin := rl.NewVector2(0, 0)

	rl.DrawTexturePro(
		m.PlayerSprite.Texture,
		sourceRec,
		destinationRec,
		origin,
		0,
		rl.White,
	)

	if m.PlayerReference.HasEquipment() {
		rl.DrawTexturePro(
			m.PlayerReference.HatSprite.Texture,
			sourceRec,
			destinationRec,
			origin,
			0.0,
			rl.White,
		)
	}

	for i, slot := range m.EquipmentSlots {
		if slot.IsEmpty {
			continue
		}

		color := rl.Gray
		textColor := rl.White
		if i == m.SelectedSlot {
			color = rl.White
			textColor = rl.White
		}

		if slot.IsEmpty {
			color = rl.Fade(rl.Gray, 0.3)
			textColor = rl.Fade(rl.White, 0.5)
		}

		rl.DrawRectangleRounded(slot.Rect, 0.1, 5, color)

		if !slot.IsEmpty {
			rl.DrawTexturePro(
				m.IconSheet,
				slot.IconPos,
				rl.NewRectangle(
					slot.Rect.X+slot.Rect.Width/2-42.5, 
					slot.Rect.Y+slot.Rect.Height/2-42.5,
					85,
					85,
				),
				rl.NewVector2(0, 0),
				0,
				rl.White,
			)
		}

		rl.DrawText(slot.Name, int32(slot.Rect.X+2), int32(slot.Rect.Y-20), 20, textColor)
	}
	rl.DrawText("Press U to unequip", int32(menuWidth)/5, int32(menuHeight), 20, rl.White)
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
		for {
			m.SelectedSlot++
			if m.SelectedSlot >= len(m.EquipmentSlots) {
				m.SelectedSlot = 0
			}
			if !m.EquipmentSlots[m.SelectedSlot].IsEmpty || m.SelectedSlot == prevSelected {
				break
			}
		}
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		for {
			m.SelectedSlot--
			if m.SelectedSlot < 0 {
				m.SelectedSlot = len(m.EquipmentSlots) - 1
			}
			if !m.EquipmentSlots[m.SelectedSlot].IsEmpty || m.SelectedSlot == prevSelected {
				break
			}
		}
	} else if rl.IsKeyPressed(rl.KeyDown) {
		for {
			m.SelectedSlot += m.Columns
			if m.SelectedSlot >= len(m.EquipmentSlots) {
				m.SelectedSlot %= len(m.EquipmentSlots)
			}
			if !m.EquipmentSlots[m.SelectedSlot].IsEmpty || m.SelectedSlot == prevSelected {
				break
			}
		}
	} else if rl.IsKeyPressed(rl.KeyUp) {
		for {
			m.SelectedSlot -= m.Columns
			if m.SelectedSlot < 0 {
				m.SelectedSlot += len(m.EquipmentSlots)
				if m.SelectedSlot < 0 {
					m.SelectedSlot = 0
				}
			}
			if !m.EquipmentSlots[m.SelectedSlot].IsEmpty || m.SelectedSlot == prevSelected {
				break
			}
		}
	}

	if prevSelected != m.SelectedSlot && m.SelectedSlot >= 0 {
		menu_move_sound := rl.LoadSound("assets/sounds/menu_move.mp3")
		rl.PlaySound(menu_move_sound)
	}

	if rl.IsKeyPressed(rl.KeyEnter) && m.SelectedSlot >= 0 && !m.EquipmentSlots[m.SelectedSlot].IsEmpty {
		m.PlayerReference.Equip(m.EquipmentSlots[m.SelectedSlot].Item)
		menu_select_sound := rl.LoadSound("assets/sounds/menu_selected.mp3")
		rl.PlaySound(menu_select_sound)
	}

	if rl.IsKeyPressed(rl.KeyU) && m.SelectedSlot >= 0 && !m.EquipmentSlots[m.SelectedSlot].IsEmpty {
		if m.PlayerReference.HasEquipment() {
			m.PlayerReference.Unequip()
			menu_select_sound := rl.LoadSound("assets/sounds/menu_selected.mp3")
			rl.PlaySound(menu_select_sound)
		}
	}
}
