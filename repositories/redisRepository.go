package repositories

import (
	"github.com/gabrielerzinger/urlShrtnr/models"
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

func (s *RedisStorage) Store(entry *models.Entry) error {
	_, err := s.Redis.Set(entry.ShortUrl, entry.URL, 0).Result()
	return err
}

func (s *RedisStorage) 	RetrieveByShortString(shortString string) (URL string, err error){
	URL, err = s.Redis.Get(shortString).Result()
	return
}


// Ping implementation
func (s *RedisStorage) Ping() error {
	_, err := s.Redis.Ping().Result()

	return err
}
