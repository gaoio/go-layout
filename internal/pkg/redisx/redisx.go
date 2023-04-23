package redisx

import (
	"github.com/go-redis/redis"
)

// New https://github.com/go-redis/redis
func New(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//return redis.NewClusterClient(&redis.ClusterOptions{
	//	Addrs: []string{addr},
	//})
}
