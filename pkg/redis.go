package pkg

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/eddoog/be-capstone/models"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() *redis.Client {
	url := os.Getenv("REDIS_URL")

	opts, err := redis.ParseURL(url)

	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}

func GetRedisClient() *redis.Client {
	if RedisClient == nil && os.Getenv("REDIS_URL") != "" {
		RedisClient = InitRedis()
	}

	return RedisClient
}

func SetKeyToRedisWithExpire(ctx context.Context, key string, value interface{}) error {

	if RedisClient == nil {
		return nil
	}

	location, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		return err
	}

	now := time.Now().In(location)

	expireTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, location)

	defer SendInfoLog("Set key to redis with expire: " + key + " with expire time: " + expireTime.String())

	jsonData, err := json.Marshal(value)

	if err != nil {
		return err
	}

	err = RedisClient.Set(ctx, key, jsonData, expireTime.Sub(now)).Err()

	if err != nil {
		return err
	}

	return nil
}

func GetKeyFromRedis(ctx context.Context, key string) (interface{}, error) {
	if RedisClient == nil {
		return "", nil
	}

	val, err := RedisClient.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	var predictionsFromRedis []models.LocationFloodPrediction

	err = json.Unmarshal([]byte(val), &predictionsFromRedis)

	if err != nil {
		return "", err
	}

	return predictionsFromRedis, nil
}
