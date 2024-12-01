package utils

import (
	"encoding/json"
	"os"
	"log"

	models "github.com/gglzc/fishMachine/models/fish"
)

var FishConfig map[int]*models.Fish

func LoadFishConfig() {
    data, err := os.ReadFile("./config/fish_config.json")
    if err != nil {
        log.Fatalf("Failed to read fish config file: %v", err)
    }
    var fishes []*models.Fish
    err = json.Unmarshal(data, &fishes)
    if err != nil {
        log.Fatalf("Failed to parse fish config: %v", err)
    }

    FishConfig = make(map[int]*models.Fish)
    for _, fish := range fishes {
        FishConfig[fish.ID] = fish
    }
}
