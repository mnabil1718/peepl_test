package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

func loadConfig() *config {
	return &config{
		fiberConfig: fiber.Config{
			ErrorHandler:    GlobalErrorHandler,
			StructValidator: &structValidator{validate: validator.New()},
		},
		redisConfig: &redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	}
}
