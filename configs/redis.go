package configs

import (
	"errors"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

// KVClient is key-value pool
var KVClient *redis.Client

// InitKV initialize new Key-Value pool
func InitKV() {
	KVClient = redis.NewClient(&redis.Options{
		// Addr:     "redis:6379", // for production
		Addr:     "127.0.0.1:6379", // for development
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
	if KVClient == nil {
		panic(errors.New("Redis client is absent, please check Redis connection"))
	}
}

// KVRegisterSource registers source
func KVRegisterSource(ID uint, URL string) {
	strID := strconv.FormatUint(uint64(ID), 10)
	exist := KVClient.Exists(strID).Val()
	if exist == 0 {
		KVClient.RPush(strID, nil, URL)
		KVClient.Expire(strID, 30*time.Second) // temporary just for development
		KVClient.SAdd("availableSources", strID)
		KVClient.Incr("sourcesCounter")
	}
}

// KVUnregisterSource adds source to list
func KVUnregisterSource(ID uint) {
	strID := strconv.FormatUint(uint64(ID), 10)
	exist := KVClient.Exists(strID).Val()
	if exist != 0 {
		KVClient.Del(strID) // temporary just for development
		KVClient.SRem("availableSources", strID)
		KVClient.Decr("sourcesCounter")
	}
}
