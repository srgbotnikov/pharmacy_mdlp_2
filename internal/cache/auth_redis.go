package cache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	rdb *redis.Client
}

type TokenCache struct {
	Token     string    `json:"token"`
	TokenTime time.Time `json:"tt"`
}

func NewRedisCache(rdb *redis.Client) *RedisCache {
	return &RedisCache{rdb: rdb}
}

//GetTokenFromCache Получить токен. Проверить истекло ли время. Если истекло, вернуть пустую строку
func (c *RedisCache) GetTokenFromCache(thumbprint string) (string, error) {
	val, err := c.rdb.Get(thumbprint).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", errors.New("Error get token from cache:" + err.Error())
	}

	var tokenCache *TokenCache

	err = json.Unmarshal([]byte(val), &tokenCache)
	if err != nil {
		return "", err
	}

	timeNow := time.Now()

	subTime := int(timeNow.Sub(tokenCache.TokenTime) / time.Minute)

	if subTime > 28 {
		return "", nil
	}

	return tokenCache.Token, nil
}

//SetTokenCache записать токен в кеш
func (c *RedisCache) SetTokenCache(token, thumbprint string) error {
	json, err := json.Marshal(TokenCache{Token: token, TokenTime: time.Now()})
	if err != nil {
		return err
	}

	err = c.rdb.Set(thumbprint, json, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
