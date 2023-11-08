package middleware

import (
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/storage/redis/v3"
)

func redisStorage() *redis.Storage {
	return redis.New(redis.Config{
		Host:      "redis",
		Port:      6379,
		Username:  "",
		Password:  "redispass",
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})
}

func Cache() func(*fiber.Ctx) error {
	return cache.New(cache.Config{
		Storage:    redisStorage(),
		Expiration: 10 * time.Minute,
		MaxBytes:   8192,
	})
}
