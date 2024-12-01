package models

type Fish struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    CaptureRate float64 `json:"capture_rate"` // 捕获概率，0-1之间的浮点数
    Reward      int64   `json:"reward"`       // 奖励金额
    // 可以添加其他属性，如外观、移动速度等
}
