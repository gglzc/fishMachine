package utils

import (
	"encoding/json"
	"log"
	"os"

	models "github.com/gglzc/fishMachine/models/fish"
	"github.com/joho/godotenv"
)

var FishConfig map[int]*models.Fish

func LoadFishConfig() {
    if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
    data, err := os.ReadFile("./config/fishConfig.json")
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
