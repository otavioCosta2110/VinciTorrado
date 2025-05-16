package cutscene

import (
	"encoding/json"
	"fmt"
	"os"
)

type CutscenesConfig struct {
	Who       string `json:"who"`
	TargetX   int    `json:"targetX"`
	TargetY   int    `json:"targetY"`
	Animation string `json:"animation"`
}

func LoadCutscenesFromJSON(filename string) ([]*CutscenesConfig, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read animations file: %w", err)
	}

	var configs []*CutscenesConfig
	if err := json.Unmarshal(file, &configs); err != nil {
		return nil, fmt.Errorf("failed to parse animations JSON: %w", err)
	}

	return configs, nil
}
