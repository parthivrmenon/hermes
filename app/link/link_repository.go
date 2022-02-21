package link

import (
	"github.com/go-redis/redis"
)

type Repository interface {
	GetLink(key string) (string, error)
	SetLink(key string, value string) error
	DelLink(key string) error

	IncLinkHits(key string) error
	GetHits() ([]redis.Z, error)
}

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(r *redis.Client) Repository {
	return &RedisRepository{r}
}

func (r *RedisRepository) GetLink(key string) (string, error) {
	return r.client.Get(key).Result()

}

func (r *RedisRepository) SetLink(key string, value string) error {
	return r.client.Set(key, value, 0).Err()

}

func (r *RedisRepository) DelLink(key string) error {
	return r.client.Del(key).Err()
}

func (r *RedisRepository) IncLinkHits(key string) error {
	err := r.client.Do("ZINCRBY", "hits", "1", key).Err()
	return err
}

func (r *RedisRepository) GetHits() ([]redis.Z, error) {
	hits, err := r.client.ZRevRangeWithScores("hits", 0, 10).Result()
	return hits, err
}
