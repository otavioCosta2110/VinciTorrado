package ui

import (
	"otaviocosta2110/vincitorrado/src/equipment"
	"otaviocosta2110/vincitorrado/src/player"
	"otaviocosta2110/vincitorrado/src/sprites"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuState int

const (
	MenuMain MenuState = iota
	MenuEquipment
	MenuStats
)

type EquipmentSlot struct {
	Rect    rl.Rectangle
	IconPos rl.Rectangle
	Item    *equipment.Equipment
	IsEmpty bool
}

type Menu struct {
	State           MenuState
	IsVisible       bool
	PlayerSprite    *sprites.Sprite
	PlayerReference *player.Player

	EquipmentSlots []EquipmentSlot
	SelectedSlot   int
	Columns        int
	Rows           int
	VisibleRows    int
	ScrollOffset   int

	IconSheet       rl.Texture2D
	ConsumableSheet rl.Texture2D

	SlotWidth   float32
	SlotHeight  float32
	SlotSpacing float32

	menuMoveSound   rl.Sound
	menuSelectSound rl.Sound

	screenWidth  int
	screenHeight int
}

func NewMenu(player *player.Player, sprite *sprites.Sprite) *Menu {
	menu := &Menu{
		State:           MenuMain,
		IsVisible:       false,
		PlayerSprite:    sprite,
		PlayerReference: player,
		SelectedSlot:    -1,
		Columns:         2,
		Rows:            0,
		VisibleRows:     3,
		ScrollOffset:    0,
		SlotWidth:       120,
		SlotHeight:      120,
		SlotSpacing:     50,
		screenWidth:     rl.GetScreenWidth(),
		screenHeight:    rl.GetScreenHeight(),
	}

	menu.loadResources()
	menu.initEquipmentSlots()

	return menu
}

func (m *Menu) loadResources() {
	m.IconSheet = rl.LoadTexture("assets/ui/equipamentos.png")
	m.ConsumableSheet = rl.LoadTexture("assets/ui/items.png")
	m.menuMoveSound = rl.LoadSound("assets/sounds/menu_move.mp3")
	m.menuSelectSound = rl.LoadSound("assets/sounds/menu_selected.mp3")
}

func (m *Menu) Unload() {
	rl.UnloadTexture(m.IconSheet)
	rl.UnloadTexture(m.ConsumableSheet)
	rl.UnloadSound(m.menuMoveSound)
	rl.UnloadSound(m.menuSelectSound)
}

func (m *Menu) initEquipmentSlots() {
	startX := float32(m.screenWidth)/2 - (float32(m.Columns)*m.SlotWidth)/2 + 50
	startY := float32(m.screenHeight / 5)

	m.EquipmentSlots = make([]EquipmentSlot, 0, len(m.PlayerReference.Equipment))
	m.Rows = (len(m.PlayerReference.Equipment) + m.Columns - 1) / m.Columns

	if m.SelectedSlot >= 0 {
		selectedRow := m.SelectedSlot / m.Columns
		if selectedRow < m.ScrollOffset {
			m.ScrollOffset = selectedRow
		} else if selectedRow >= m.ScrollOffset+m.VisibleRows {
			m.ScrollOffset = selectedRow - m.VisibleRows + 1
		}
	}

	maxScroll := max(0, m.Rows-m.VisibleRows)
	m.ScrollOffset = min(m.ScrollOffset, maxScroll)

	for i, item := range m.PlayerReference.Equipment {
		row := i / m.Columns
		col := i % m.Columns

		if row >= m.ScrollOffset && row < m.ScrollOffset+m.VisibleRows {
			m.EquipmentSlots = append(m.EquipmentSlots, EquipmentSlot{
				Rect: rl.NewRectangle(
					startX+float32(col)*(m.SlotWidth+m.SlotSpacing),
					startY+float32(row-m.ScrollOffset)*(m.SlotHeight+m.SlotSpacing),
					m.SlotWidth,
					m.SlotHeight,
				),
				IconPos: m.getItemIconPos(item),
				Item:    item,
				IsEmpty: false,
			})
		}
	}

	if m.SelectedSlot >= len(m.EquipmentSlots) || m.SelectedSlot < 0 {
		m.SelectedSlot = -1
		for i, slot := range m.EquipmentSlots {
			if !slot.IsEmpty {
				m.SelectedSlot = i
				break
			}
		}
	}
}

func (m *Menu) getItemIconPos(item *equipment.Equipment) rl.Rectangle {
	if item == nil {
		return rl.NewRectangle(0, 0, 32, 32)
	}

	iconMap := map[string]rl.Rectangle{
		"Turbante": rl.NewRectangle(32, 32, 32, 32),
		"Suit":     rl.NewRectangle(64, 32, 32, 32),
		"Shoes":    rl.NewRectangle(96, 32, 32, 32),
	}

	consumableMap := map[string]rl.Rectangle{
		"Hamburgão":  rl.NewRectangle(0, 0, 32, 32),
		"Saunduiche": rl.NewRectangle(32, 0, 32, 32),
		"Cachaca":    rl.NewRectangle(0, 32, 32, 32),
		"guarana":    rl.NewRectangle(0, 64, 32, 32),
	}

	if rect, exists := iconMap[item.Name]; exists {
		return rect
	}

	if item.Type == "consumable" {
		if rect, exists := consumableMap[item.Name]; exists {
			return rect
		}
		return rl.NewRectangle(64, 0, 32, 32)
	}

	return rl.NewRectangle(0, 0, 32, 32)
}

func (m *Menu) Draw() {
	if !m.IsVisible {
		return
	}

	rl.DrawRectangle(0, 0, int32(m.screenWidth), int32(m.screenHeight), rl.Fade(rl.Black, 0.5))

	menuWidth := float32(m.screenWidth) * 0.8
	menuHeight := float32(m.screenHeight) * 0.8
	menuX := (float32(m.screenWidth) - menuWidth) / 2
	menuY := (float32(m.screenHeight) - menuHeight) / 2

	rl.DrawRectangleRounded(
		rl.NewRectangle(menuX, menuY, menuWidth, menuHeight),
		0.05, 10, rl.DarkGray,
	)

	m.drawPlayerPreview(menuX, menuY)

	m.drawEquipmentSlots()

	if m.hasValidSelection() {
		m.drawItemStats(menuX, menuY, menuWidth)
	}

	m.drawInstructions(menuWidth, menuHeight)
	m.drawScrollBar(menuX, menuY, menuWidth, menuHeight)
}

func (m *Menu) drawPlayerPreview(x, y float32) {
	playerPreviewY := y * 2
	sourceRec := rl.NewRectangle(0, 0,
		float32(m.PlayerSprite.SpriteWidth),
		float32(m.PlayerSprite.SpriteHeight))

	scale := float32(m.PlayerReference.Object.Scale * 3)
	destinationRec := rl.NewRectangle(
		x,
		playerPreviewY,
		float32(m.PlayerSprite.SpriteWidth)*scale,
		float32(m.PlayerSprite.SpriteHeight)*scale,
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
}

func (m *Menu) drawEquipmentSlots() {
	for i, slot := range m.EquipmentSlots {
		color := rl.Gray
		textColor := rl.White
		if i == m.SelectedSlot {
			color = rl.White
		}
		if slot.IsEmpty {
			color = rl.Fade(rl.Gray, 0.3)
			textColor = rl.Fade(rl.White, 0.5)
		}

		rl.DrawRectangleRounded(slot.Rect, 0.1, 5, color)

		if !slot.IsEmpty {
			texture := m.IconSheet
			if slot.Item.Type == "consumable" {
				texture = m.ConsumableSheet
			}

			iconRect := rl.NewRectangle(
				slot.Rect.X+slot.Rect.Width/2-42.5,
				slot.Rect.Y+slot.Rect.Height/2-42.5,
				85,
				85,
			)

			rl.DrawTexturePro(
				texture,
				slot.IconPos,
				iconRect,
				rl.NewVector2(0, 0),
				0,
				rl.White,
			)

			rl.DrawText(slot.Item.Name,
				int32(slot.Rect.X+2),
				int32(slot.Rect.Y-20),
				20, textColor)
		} else {
			rl.DrawText("Empty",
				int32(slot.Rect.X+2),
				int32(slot.Rect.Y-20),
				20, textColor)
		}
	}
}

func (m *Menu) drawItemStats(x, y, width float32) {
	item := m.EquipmentSlots[m.SelectedSlot].Item
	statsX := x + width - 250
	statsY := y + 50

	rl.DrawRectangleRounded(
		rl.NewRectangle(statsX, statsY, 200, 150),
		0.1, 5, rl.Fade(rl.Black, 0.7),
	)

	rl.DrawText(item.Name, int32(statsX+10), int32(statsY+35), 20, rl.Gold)

	yOffset := int32(60)
	m.drawStat("Health: ", item.Stats.Life, statsX, statsY, &yOffset)
	m.drawStat("Damage: ", item.Stats.Damage, statsX, statsY, &yOffset)
	m.drawStat("Speed: ", item.Stats.Speed, statsX, statsY, &yOffset)

	if item.Type == "consumable" {
		rl.DrawText("Heal: "+strconv.Itoa(int(item.Stats.Heal)),
			int32(statsX+10), int32(statsY)+yOffset+50, 20, rl.Green)
	}
}

func (m *Menu) drawStat(label string, value int32, x, y float32, yOffset *int32) {
	if value != 0 {
		rl.DrawText(label, int32(x+10), int32(y)+*yOffset, 18, rl.White)
		rl.DrawText(formatStat(int(value)),
			int32(x+100), int32(y)+*yOffset, 18,
			getStatColor(int(value)))
		*yOffset += 25
	}
}

func (m *Menu) drawInstructions(width, height float32) {
	if m.PlayerReference.HasEquipment() {
		rl.DrawText("Press U to unequip",
			int32(width)/5, int32(height), 20, rl.White)
	}

	rl.DrawText("ESC to close",
		int32(width)/5, int32(height)+30, 20, rl.White)
}

func (m *Menu) drawScrollBar(x, y, width, height float32) {
	if m.Rows <= m.VisibleRows {
		return
	}

	barWidth := float32(20)
	barHeight := height * 0.8
	barX := x + width - barWidth - 10
	barY := y + (height-barHeight)/2

	rl.DrawRectangleRounded(
		rl.NewRectangle(barX, barY, barWidth, barHeight),
		0.5, 10, rl.Fade(rl.Black, 0.5),
	)

	thumbHeight := barHeight * float32(m.VisibleRows) / float32(m.Rows)
	thumbY := barY + (barHeight-thumbHeight)*float32(m.ScrollOffset)/float32(m.Rows-m.VisibleRows)

	rl.DrawRectangleRounded(
		rl.NewRectangle(barX, thumbY, barWidth, thumbHeight),
		0.5, 10, rl.White,
	)
}

func (m *Menu) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		m.toggleVisibility()
	}

	if !m.IsVisible {
		return
	}

	m.handleNavigation()

	if rl.IsKeyPressed(rl.KeyEnter) && m.hasValidSelection() {
		m.handleItemSelection()
	}

	if rl.IsKeyPressed(rl.KeyU) && m.PlayerReference.HasEquipment() {
		m.PlayerReference.Unequip()
		rl.PlaySound(m.menuSelectSound)
	}

	if rl.IsKeyPressed(rl.KeyPageUp) {
		m.ScrollOffset = max(0, m.ScrollOffset-m.VisibleRows)
		m.Refresh()
	}

	if rl.IsKeyPressed(rl.KeyPageDown) {
		m.ScrollOffset = min(m.Rows-m.VisibleRows, m.ScrollOffset+m.VisibleRows)
		m.Refresh()
	}
}

func (m *Menu) toggleVisibility() {
	m.IsVisible = !m.IsVisible
	if m.IsVisible {
		m.Refresh()
	}
}

func (m *Menu) handleNavigation() {
	prevSelected := m.SelectedSlot

	if rl.IsKeyPressed(rl.KeyRight) {
		m.findNextValidSlot(1)
	} else if rl.IsKeyPressed(rl.KeyLeft) {
		m.findNextValidSlot(-1)
	} else if rl.IsKeyPressed(rl.KeyDown) {
		m.findNextValidSlot(m.Columns)
	} else if rl.IsKeyPressed(rl.KeyUp) {
		m.findNextValidSlot(-m.Columns)
	}

	if prevSelected != m.SelectedSlot && m.SelectedSlot >= 0 {
		rl.PlaySound(m.menuMoveSound)
	}
}

func (m *Menu) handleItemSelection() {
	item := m.EquipmentSlots[m.SelectedSlot].Item

	if item.Type == "consumable" {
		m.PlayerReference.UseConsumable(m.SelectedSlot)
		m.Refresh()
		rl.PlaySound(m.menuSelectSound)
	} else {
		if m.PlayerReference.Equipped != item {
			m.PlayerReference.Unequip()
		}
		if m.PlayerReference.Equipped != item {
			m.PlayerReference.Equip(item)
			rl.PlaySound(m.menuSelectSound)
		}
	}
}

func (m *Menu) findNextValidSlot(step int) {
	if len(m.EquipmentSlots) == 0 {
		m.SelectedSlot = -1
		return
	}

	currentPos := max(m.SelectedSlot, 0)
	currentCol := currentPos % m.Columns
	currentRow := currentPos / m.Columns

	if step == 1 || step == -1 {
		next := currentPos + step

		if next >= 0 && next < len(m.EquipmentSlots) && (next/m.Columns) == currentRow {
			for i := next; i >= 0 && i < len(m.EquipmentSlots) && (i/m.Columns) == currentRow; i += step {
				if !m.EquipmentSlots[i].IsEmpty {
					m.SelectedSlot = i
					rl.PlaySound(m.menuMoveSound)
					return
				}
			}
		}
		return
	}

	if step == m.Columns || step == -m.Columns {
		next := currentPos + step
		nextRow := next / m.Columns
		if next < 0 {
			if m.ScrollOffset > 0 {
				m.ScrollOffset--
				m.Refresh()
				m.SelectedSlot = m.ScrollOffset*m.Columns + currentCol
				if m.SelectedSlot >= len(m.EquipmentSlots) {
					m.SelectedSlot = len(m.EquipmentSlots) + 1
				}
				rl.PlaySound(m.menuMoveSound)
			}
			return
		} else if next >= len(m.EquipmentSlots) {
			if m.ScrollOffset < m.Rows-m.VisibleRows {
				m.ScrollOffset++
				m.Refresh()
				m.SelectedSlot = min((m.ScrollOffset+m.VisibleRows-1)*m.Columns+currentCol, len(m.EquipmentSlots)-1)
				newPos := nextRow*m.Columns + currentCol
				if newPos < len(m.EquipmentSlots)+1 {
					m.SelectedSlot = newPos - 2
				}
				rl.PlaySound(m.menuMoveSound)
			}
			return
		}

		if nextRow < m.ScrollOffset {
			if m.ScrollOffset > 0 {
				m.ScrollOffset = nextRow + 1
				m.Refresh()
			}
		}

		for i := next; i >= 0 && i < len(m.EquipmentSlots); i += step {
			if !m.EquipmentSlots[i].IsEmpty && (i%m.Columns) == currentCol {
				m.SelectedSlot = i
				rl.PlaySound(m.menuMoveSound)
				return
			}
		}
	}
}

func (m *Menu) hasValidSelection() bool {
	return m.SelectedSlot >= 0 &&
		m.SelectedSlot < len(m.EquipmentSlots) &&
		!m.EquipmentSlots[m.SelectedSlot].IsEmpty
}

func (m *Menu) Refresh() {
	m.screenWidth = rl.GetScreenWidth()
	m.screenHeight = rl.GetScreenHeight()
	m.initEquipmentSlots()
}

func formatStat(value int) string {
	if value > 0 {
		return "+" + strconv.Itoa(value)
	}
	return strconv.Itoa(value)
}

func getStatColor(value int) rl.Color {
	if value > 0 {
		return rl.Green
	} else if value < 0 {
		return rl.Red
	}
	return rl.White
}
