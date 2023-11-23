package api

import (
	"base-plate/app"
	"base-plate/app/api/handler"
	"base-plate/config"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.elastic.co/apm/module/apmfiber/v2"
	"log"
	"os"
	"os/signal"
)

func configure(ac *app.Container) *fiber.App {
	apps := fiber.New()

	apps.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.JSON("PONG")
	})

	apps = loadRoute(apps, ac)

	return apps
}

func Serve(ctx context.Context, ac *app.Container) {
	apps := configure(ac)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = apps.Shutdown()

		_ = ac.Cache.Close()
		_ = ac.MongoClient.Client.Disconnect(ctx)
		_ = ac.KafkaSubscriber.Close()
		_ = ac.KafkaPublisher.Close()
	}()

	if err := apps.Listen(fmt.Sprintf(":%d", config.Cfg.GetInt("APP_PORT"))); err != nil {
		log.Println("Failed to start server", err.Error())
	}

	fmt.Println("Running cleanup tasks...")
}

func loadRoute(app *fiber.App, ac *app.Container) *fiber.App {
	app.Use(apmfiber.Middleware())
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(pprof.New())

	app = handler.LoadHandler(app, ac)

	return app
}
