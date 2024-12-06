package repositories

import (
	// models "github.com/gglzc/fishMachine/models/room"
	"github.com/redis/go-redis/v9"
)

type RoomRepository struct {
    Redis *redis.Client
}

func(r *RoomRepository)CreateRoom(){
	
}