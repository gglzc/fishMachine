package services

import (
    "errors"
    "sync"

    "github.com/gglzc/fishMachine/service/user"
)

type GameService struct {
    UserService *services.UserService
    // RoomService 
    // 使用 mutex 確保餘額操作的線程安全
    BalanceMutex sync.Mutex
}

func (gs *GameService) ShootBullet(userID int64, cost int64) error {
    // 確保餘額操作的原子性
    gs.BalanceMutex.Lock()
    defer gs.BalanceMutex.Unlock()

    user, err := gs.UserService.GetUserByID(userID)
    if err != nil {
        return err
    }
    if user.Balance < cost {
        return errors.New("insufficient balance")
    }

    // 扣除餘額
    err = gs.UserService.UpdateBalance(userID, -cost)
    if err != nil {
        return err
    }

    // 處理遊戲邏輯，例如碰撞檢測和獎勵計算（此處簡化）
    reward := calculateReward()

    // 更新玩家餘額
    err = gs.UserService.UpdateBalance(userID, reward)
    if err != nil {
        return err
    }

    return nil
}

func calculateReward() int64 {
    // 簡化的隨機獎勵計算
    return 100
}
