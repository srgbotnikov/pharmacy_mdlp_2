package cache

import (
	"github.com/go-redis/redis"
)

type ConfigRedis struct {
	Host string
	Port string
	// Username string
	// Password string
}

// var ctx = context.Background()

func NewRedisDB(cfg ConfigRedis) (*redis.Client, error) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil

}
