package services

import (
    "sync"
	"errors"
    "github.com/gglzc/fishMachine/models/room"
)

type RoomService struct {
    Rooms sync.Map // key: room ID, value: *models.Room
}

func (rs *RoomService) GetOrCreateRoom(currency string) (*models.Room, error) {
    var room *models.Room

    // 查找已有未滿的房間
    rs.Rooms.Range(func(key, value interface{}) bool {
        r := value.(*models.Room)
        if r.Currency == currency && !r.IsFull() {
            room = r
            return false // 結束遍歷
        }
        return true
    })

    // 如果沒有可用的房間，創建新房間
    if room == nil {
        room = &models.Room{
            ID:       1,
            Currency: currency,
            Players:  []int64{},
        }
        rs.Rooms.Store(room.ID, room)
    }

    return room, nil
}

func (rs *RoomService) AddPlayerToRoom(roomID int64, playerID int64) error {
    value, ok := rs.Rooms.Load(roomID)
    if !ok {
        return errors.New("Room not found")
    }
    room := value.(*models.Room)
    if room.IsFull() {
        return errors.New("Room is full")
    }
    room.Players = append(room.Players, playerID)
    return nil
}
