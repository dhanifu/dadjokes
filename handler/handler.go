package handler

import (
	"github.com/dhanifu/dadjokes/config"
	"github.com/dhanifu/dadjokes/handler/dadjoke"
)

type Handler struct {
	DadJoke dadjoke.DadJokeHandler
}

func InitHandlers(config *config.Configuration) *Handler {
	return &Handler{
		dadjoke.InitDadJoke(config),
	}
}