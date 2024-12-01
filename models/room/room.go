package models

type Room struct {
    ID       int64    `json:"id"`
    Currency string   `json:"currency"`
    Players  []int64  `json:"players"`
}

func (r *Room) IsFull() bool {
    return len(r.Players) >= 4
}
