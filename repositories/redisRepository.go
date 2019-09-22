package repositories

import (

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// RedisStorage struct
type RedisStorage struct {
	Redis *redis.Client
}

// NewRedisStorage ctor
func NewRedisStorage() *RedisStorage {
	return &RedisStorage{}
}

// Connect connects to redis db
func (s *RedisStorage) Connect(config *viper.Viper) error {
	client := redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.url"),
		Password: config.GetString("redis.password"),
		DB:       0,
	})

	_, err := client.Ping().Result()

	if err != nil {
		panic("Failed to connect to redis.")
	}

	s.Redis = client

	return err
}

// Ping implementation
func (s *RedisStorage) Ping() error {
	_, err := s.Redis.Ping().Result()

	return err
}
