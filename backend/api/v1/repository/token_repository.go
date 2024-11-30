package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"stageflow/config/initializers"
	"time"
)

type TokenRepository struct {
	Redis *redis.Client
}

func NewTokenRepository() *TokenRepository {
	return &TokenRepository{
		Redis: initializers.GetRedisClient(),
	}
}

func (tr *TokenRepository) SetUserJWT(userID string, token string) error {
	ctx := context.Background()
	expiration := 24 * time.Hour
	err := tr.Redis.Set(ctx, userID, token, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
