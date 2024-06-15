package routes

import (
	"github.com/dhanifu/dadjokes/config"
	"github.com/dhanifu/dadjokes/handler"
	"github.com/gofiber/fiber/v2"
)

type router struct {
	handler *handler.Handler
	config *config.Configuration
}

func InitRouter(handler *handler.Handler, config *config.Configuration) router {
	return router{
		handler: handler,
		config:  config,
	}
}

func (r router) Route(router fiber.Router) {
	router.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Hello World",
		})
	})

	v1 := router.Group("v1")

	r.DadJokeRouter(v1.Group("dadjokes"))
}