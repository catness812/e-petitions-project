package repository

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) *RedisRepository {
	return &RedisRepository{redisClient: redisClient}
}
func (redisRepo *RedisRepository) InsertUserToken(key uint, value string, expires time.Duration) (error, string) {

	err := redisRepo.redisClient.Set(context.Background(), strconv.FormatUint(uint64(key), 10), value, expires).Err()

	if err != nil{
		log.Printf("failed to set token: %v\n", err)
	}

	return err, "Insertion successful."
}

func (redisRepo *RedisRepository) ReplaceToken(currentToken, newToken string, expires time.Duration) error {
	id, err := redisRepo.deleteToken(currentToken)
	if err != nil {
		return err
	}
	return redisRepo.redisClient.Set(context.Background(), newToken, id, expires).Err()
}

func (redisRepo *RedisRepository) deleteToken(token string) (string, error) {
	return redisRepo.redisClient.GetDel(context.Background(), token).Result()
}