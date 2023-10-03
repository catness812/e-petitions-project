package security_repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) *RedisRepository {
	return &RedisRepository{redisClient: redisClient}
}
func (redisRepo *RedisRepository) InsertUserToken(key string, value uint32, expires time.Duration) error {
	return redisRepo.redisClient.Set(context.Background(), key, value, expires).Err()
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

func (redisRepo *RedisRepository) InsertOTP(otp string, mail string, expires time.Duration) error {
	return redisRepo.redisClient.Set(context.Background(), mail, otp, expires).Err()
}

func (redisRepo *RedisRepository) GetOTP(mail string) (string, error) {
	otp, err := redisRepo.redisClient.Get(context.Background(), mail).Result()
	if err != nil {
		return "", err
	}
	return otp, nil
}
