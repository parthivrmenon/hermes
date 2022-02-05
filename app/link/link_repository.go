package link

import "github.com/go-redis/redis"

type Repository interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Del(key string) error
}

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(r *redis.Client) Repository {
	return &RedisRepository{r}
}

func (r *RedisRepository) Get(key string) (string, error) {
	return r.client.Get(key).Result()

}

func (r *RedisRepository) Set(key string, value string) error {
	return r.client.Set(key, value, 0).Err()
}

func (r *RedisRepository) Del(key string) error {
	return r.client.Del(key).Err()
}
