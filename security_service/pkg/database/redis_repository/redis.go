package redis_repository

import (
	"log"

	"github.com/catness812/e-petitions-project/security_service/internal/security/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisDBConnection() *redis.Client {
	cfg := config.LoadConfig()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + cfg.Redis.Port,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.RedisDB,
	})
	if client == nil {
		log.Fatalf("failed to connect to RedisDB")
	}
	log.Println("Successfully connected to RedisDB")
	return client
}
