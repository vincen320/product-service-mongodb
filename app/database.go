package app

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() (*mongo.Database, error) {
	clientOptions := options.Client().
		SetMaxPoolSize(100).
		SetMaxConnecting(20).
		SetConnectTimeout(60 * time.Minute).
		SetMaxConnIdleTime(10 * time.Minute).
		ApplyURI("mongodb://127.0.0.1:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client.Database("v_product"), nil
}

func ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		MaxRetries:   2,
		MaxConnAge:   60 * time.Minute,
		IdleTimeout:  10 * time.Minute,
		MinIdleConns: 5,
		PoolSize:     20,
	})
	return rdb
}
