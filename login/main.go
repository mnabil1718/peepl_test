package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
)

type config struct {
	fiberConfig fiber.Config
	redisConfig *redis.Options
}

type application struct {
	rdb    *redis.Client
	server *fiber.App
}

func bootstrapApp(cfg *config) *application {
	srv := fiber.New(cfg.fiberConfig)
	rdb := redis.NewClient(cfg.redisConfig)

	app := &application{
		server: srv,
		rdb:    rdb,
	}

	app.seedRedis()
	app.registerRoutes()
	return app
}

func main() {
	cfg := loadConfig()
	app := bootstrapApp(cfg)
	log.Fatal(app.server.Listen(":8080"))
}
