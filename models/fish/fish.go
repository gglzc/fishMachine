package models

type Fish struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    CaptureRate float64 `json:"capture_rate"` 
    Reward      int64   `json:"reward"`      

}
