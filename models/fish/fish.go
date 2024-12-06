package models

type Fish struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    CaptureRate float64 `json:"capture_rate"`
    Reward      int64   `json:"reward"`
}

type Position struct {
    X float64
    Y float64
}

type Path struct {
    Points []Position
}

type FishInstance struct {
    Fish     *Fish
    Position Position
    Path     Path
}
