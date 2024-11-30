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

func (tr *TokenRepository) BlackListJWT(userID string, token string, ctx context.Context) error {
	expiration := 24 * time.Hour
	err := tr.Redis.SAdd(ctx, "blacklist:user:"+userID, token, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (tr *TokenRepository) FindJWT(userID string, ctx context.Context) ([]string, error) {
	tokens, err := tr.Redis.SMembers(ctx, "blacklist:user:"+userID).Result()
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
