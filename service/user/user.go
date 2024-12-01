package services

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gglzc/fishMachine/models/user"
	"github.com/gglzc/fishMachine/repository/user"
)

type UserService struct {
    UserRepo *repositories.UserRepository
}

func (us *UserService) GetUserByID(id int64) (*models.User, error) {
	ctx := context.Background()
    user := &models.User{}

    // 將 int64 的 id 轉為字符串作為 Redis 的鍵
    redisKey := strconv.FormatInt(id, 10)

    // 嘗試從 Redis 獲取數據
    if err := us.UserRepo.Redis.Get(ctx, redisKey).Scan(user); err == nil {
        return user, nil
    }

    // 如果 Redis 沒有數據，查詢 MySQL
    user, err := us.UserRepo.GetUserByID(id)
    if err != nil {
        return nil, err
    }

    // 將查詢結果緩存到 Redis
    userJson, err := json.Marshal(user) // 將 user 結構序列化為 JSON
	if err != nil {
        return nil, err
    }
	//ttl = 600 second
    err = us.UserRepo.Redis.Set(ctx, redisKey, userJson, 10*time.Minute).Err()
    if err != nil {
        return nil, fmt.Errorf("failed to cache user data in Redis: %v", err)
    }

    return user, nil
}
func (us *UserService) CreateUser(user models.User) (error) {
    err:=us.UserRepo.CreateUser(user)
	if(err!=nil){
		return err
	}
	return nil
}

func (us *UserService) UpdateBalance(userID int64, amount int64) error {
    return us.UserRepo.UpdateBalance(userID, amount)
}

