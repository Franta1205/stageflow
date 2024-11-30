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

func (tr *TokenRepository) BlackListJWT(ctx context.Context, userID string, jwt string) error {
	expiration := 24 * time.Hour
	err := tr.Redis.SAdd(ctx, "blacklist:user:"+userID, jwt, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (tr *TokenRepository) FindJWT(ctx context.Context, userID string) ([]string, error) {
	tokens, err := tr.Redis.SMembers(ctx, "blacklist:user:"+userID).Result()
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
