package redis_repository

import (
	"github.com/catness812/e-petitions-project/security_service/internal/config"
	"github.com/gookit/slog"
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
		slog.Fatal("failed to connect to RedisDB")
	}
	slog.Info("Successfully connected to RedisDB")
	return client
}
