package repository

import (
	"github.com/go-redis/redis/v8"
)

type RedisRepository struct{}

func (u RedisRepository) CheckIfTokenExists(rdb *redis.Client, token string) error {
	return rdb.Get(token).Err()
}
