package main

import (
	"context"
	"log/slog"

	"github.com/yugovtr/goredisbench/entity"
	"github.com/yugovtr/goredisbench/redis"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelInfo)

	client, callback := redis.New()
	defer callback()

	zipcode := entity.FactoryZipcode()
	jsonCache(client, zipcode)
	gobCache(client, zipcode)
}

func jsonCache(client redis.Client, zipcode entity.Zipcode) {
	ctx := context.Background()
	jsonZipcode := entity.JSONZipcode(zipcode)

	key := "json." + jsonZipcode.Number
	if err := client.Set(ctx, key, jsonZipcode, 0).Err(); err != nil {
		slog.Error("redis save", "description", err)
		return
	}
	slog.Debug("structure saved in redis", "key", key)

	sZipcode := &entity.JSONZipcode{}
	if err := client.Get(ctx, key).Scan(sZipcode); err != nil {
		slog.Error("redis get", "description", err)
		return
	}
	slog.Debug("structure retrieved from redis", "zipcode", sZipcode)

	size, _ := client.MemoryUsage(ctx, key).Result()
	slog.Info("JSON - memory usage", "key", key, "size", size)
}

func gobCache(client redis.Client, zipcode entity.Zipcode) {
	ctx := context.Background()
	gobZipcode := entity.GOBZipcode(zipcode)

	key := "gob." + gobZipcode.Number
	if err := client.Set(ctx, key, gobZipcode, 0).Err(); err != nil {
		slog.Error("redis save", "description", err)
		return
	}
	slog.Debug("structure saved in redis", "key", key)

	sZipcode := &entity.GOBZipcode{}
	if err := client.Get(ctx, key).Scan(sZipcode); err != nil {
		slog.Error("redis get", "description", err)
		return
	}
	slog.Debug("structure retrieved from redis", "zipcode", sZipcode)

	size, _ := client.MemoryUsage(ctx, key).Result()
	slog.Info("GOB  - memory usage", "key", key, "size", size)
}
