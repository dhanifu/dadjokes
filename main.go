package main

import (
	"fmt"
	"os"

	"github.com/dhanifu/dadjokes/config"
	"github.com/dhanifu/dadjokes/data"
	"github.com/dhanifu/dadjokes/handler"
	"github.com/dhanifu/dadjokes/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.LoadConfig(".")

	app := fiber.New()

	data.InitData(config)
	handler := handler.InitHandlers(config)

	app.Route("", routes.InitRouter(handler, config).Route)

	port := os.Getenv("PORT")
	if port == "" {
		port = config.AppPort
	}
	
	app.Listen(fmt.Sprintf("0.0.0.0:%s", port))
}