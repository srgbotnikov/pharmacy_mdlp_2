package cache

import (
	"github.com/go-redis/redis"
)

type Authorization interface {
	//GetTokenFromCache получить токен из кеша
	GetTokenFromCache(thumbprint string) (string, error)
	//SetTokenCache записать токен в кеш
	SetTokenCache(token, thumbprint string) error
}

type Cache struct {
	Authorization
}

func NewCache(rdb *redis.Client) *Cache {
	return &Cache{
		Authorization: NewRedisCache(rdb),
	}
}
