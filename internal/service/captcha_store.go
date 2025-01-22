package service

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisStore struct {
	client     *redis.Client
	expiration time.Duration
	keyPrefix  string
}

func NewRedisStore(client *redis.Client, expiration time.Duration) *redisStore {
	return &redisStore{
		client:     client,
		expiration: expiration,
		keyPrefix:  "captcha:",
	}
}

func (s *redisStore) Set(id string, value string) error {
	ctx := context.Background()
	return s.client.Set(ctx, s.keyPrefix+id, value, s.expiration).Err()
}

func (s *redisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	key := s.keyPrefix + id
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	if clear {
		s.client.Del(ctx, key)
	}
	return val
}

func (s *redisStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}
