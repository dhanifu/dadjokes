package routes

import "github.com/gofiber/fiber/v2"

func (r router) DadJokeRouter(group fiber.Router) fiber.Router {
	group.Get("/", r.handler.DadJoke.GetDadRandomJokes)

	return group
}