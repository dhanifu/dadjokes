package routes

import "github.com/gofiber/fiber/v2"

func (r router) DadJokeRouterV2(group fiber.Router) fiber.Router {
	group.Get("/", r.handler.DadJoke.GetRandomDadJokes)

	return group
} 