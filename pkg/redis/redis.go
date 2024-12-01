package pkg

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)


func InitRedis() (*redis.Client,error){
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	addr := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,      
		Password: password,  
		DB:       0,         
	})
	// test connection
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	log.Println("Connected to Redis successfully")
	return redisClient, nil
}
