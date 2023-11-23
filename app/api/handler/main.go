package handler

import (
	"base-plate/app"
	"base-plate/app/api/handler/checker"
	checker2 "base-plate/repositories/checker"
	checkerSvc "base-plate/services/checker"
	"github.com/gofiber/fiber/v2"
	jwtFiber "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

func loadCheckerRoute(c *app.Container, router *fiber.App) *fiber.App {
	checkRepo := checker2.NewCheckerRepository(c.SqlClient, c.PSqlClient)
	checkSvc := checkerSvc.NewCheckerService(checkRepo, c.KafkaPublisher, c.KafkaSubscriber, c.Cache, c.MongoClient, c.SqlClient, c.PSqlClient, c.HttpClient)
	handler := checker.NewCheckerHandler(checkSvc)

	router.Get("/health-check", handler.HealthCheck)
	return router
}
func LoadHandler(router *fiber.App, c *app.Container) *fiber.App {
	//router without jwt
	router = loadCheckerRoute(c, router)

	//using jwt
	router.Use(setBearerRule)
	router.Use(jwtFiber.New(jwtFiber.Config{
		SigningMethod: "HS256",
		SigningKey:    []byte("secret"),
	}))
	router.Use(validateJWTClient)
	return router
}

func setBearerRule(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if strings.HasPrefix(tokenString, "Bearer") == false {
		c.Request().Header.Set("Authorization", "Bearer "+tokenString)
	}

	return c.Next()
}

func validateJWTClient(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	if claims, ok := user.Claims.(jwt.MapClaims); ok {
		c.Locals("user_data", map[string]interface{}(claims))
		return c.Next()
	}

	return c.SendStatus(http.StatusUnauthorized)
}
