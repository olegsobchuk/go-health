package configs

import (
	"errors"

	"github.com/go-redis/redis"
)

// KVClient is key-value pool
var KVClient *redis.Client

// InitKV initialize new Key-Value pool
func InitKV() {
	KVClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if KVClient == nil {
		panic(errors.New("Redis client is absent, please check Redis connection"))
	}
}
