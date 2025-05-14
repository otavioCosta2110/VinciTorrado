package system

type GameMap struct {
	Buildings    string
	Floor        string
	EnemiesPath  string
	PropsPath    string
	NextMap      string
	PlayerStartX int32
	PlayerStartY int32
}

type MapManager struct {
	CurrentMap *GameMap
	Maps       map[string]*GameMap
}

func NewMapManager() *MapManager {
	return &MapManager{
		Maps: make(map[string]*GameMap),
	}
}

func (mm *MapManager) LoadMap(name string) {
	mm.CurrentMap = mm.Maps[name]
}
