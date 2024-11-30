package repository

import (
	"context"
	"errors"
	"fmt"
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
	key := fmt.Sprintf("blacklist:token:%s", jwt)
	expiration := 24 * time.Hour
	err := tr.Redis.Set(ctx, key, userID, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (tr *TokenRepository) FindJWT(ctx context.Context, jwt string) (*string, error) {
	key := fmt.Sprintf("blacklist:token:%s", jwt)
	token, err := tr.Redis.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &token, nil
}
