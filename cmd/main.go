package main

import (
	"fmt"

	"github.com/dhanifu/dadjokes/config"
	"github.com/dhanifu/dadjokes/handler"
	"github.com/dhanifu/dadjokes/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.LoadConfig(".")

	app := fiber.New()

	handler := handler.InitHandlers(config)

	app.Route("", routes.InitRouter(handler, config).Route)

	app.Listen(fmt.Sprintf(":%s", config.AppPort))
}