package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-ecommerce-app2/config"
	"go-ecommerce-app2/internal/api/rest"
	"go-ecommerce-app2/internal/api/rest/handlers"
	"go-ecommerce-app2/internal/domain"
	"go-ecommerce-app2/internal/helper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func StartServer(config config.AppConfig) {
	app := fiber.New()

	//log.Println(config.Dsn)
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection error:%v", err)
	}
	log.Println("Database connected")

	app.Get("/health", HealthCheck)

	// run migration
	db.AutoMigrate(&domain.User{})
	auth := helper.SetupAuth(config.AppSecret)
	log.Println("User migration successfully")

	// init rest handler
	rh := &rest.RestHandler{
		App:  app,
		DB:   db,
		Auth: auth,
	}

	SetupRoutes(rh)

	httpBasePath := fmt.Sprintf("%s:%s", config.ServerHost, config.ServerPort)

	app.Listen(httpBasePath)
}

func SetupRoutes(rh *rest.RestHandler) {
	// user routes
	handlers.SetupUserRoutes(rh)
	// catalog routes
	// transaction routes
}

func HealthCheck(ctx *fiber.Ctx) error {
	ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "working fine",
	})
	return nil
}
