package redis

import (
	"context"

	"github.com/arthasyou/utility-go/logger"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// Ctx context
var Ctx = context.Background()

// Cli simple client
var Cli *redis.Client

// Cluster client
var Cluster *redis.ClusterClient

// Connect to redis
func Connect(addr string, password string, db int) {
	Cli = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	_, err := Cli.Ping(Ctx).Result()
	if err != nil {
		logger.Error("redis Ping", zap.String("err", err.Error()))
		return
	}
}

// Close redis
func Close() {
	Cli.Close()
}

// ConnectCluster to redis
func ConnectCluster(addrs []string, poolSize int, password string) {
	Cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
		PoolSize: viper.GetInt("Redis.ClusterPoolSize"),
	})
	_, err := Cluster.Ping(Ctx).Result()
	if err != nil {
		logger.Error("redis Ping", zap.String("err", err.Error()))
	}
}

// CloseCluster redis
func CloseCluster() {
	Cluster.Close()
}
