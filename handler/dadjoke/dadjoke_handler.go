package dadjoke

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/dhanifu/dadjokes/config"
	"github.com/dhanifu/dadjokes/models"
	"github.com/gofiber/fiber/v2"
)

type DadJokeHandler interface {
	GetDadRandomJokes(c *fiber.Ctx) error
}

type handler struct {
	config *config.Configuration
}

func InitDadJoke(config *config.Configuration) *handler {
	return &handler{
		config: config,
	}
}

func (h handler) GetDadRandomJokes(c *fiber.Ctx) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	candaanEndpoint := fmt.Sprintf("%s/%s", h.config.CandaanEndpoint, "api/text/random")
	JokesBapackEndpoint := fmt.Sprintf("%s/%s", h.config.JokesBapackEndpoint, "v1/text/random")
	endpoints := []string{
		candaanEndpoint,
		JokesBapackEndpoint,
	}
	randomEndpointIndex := r.Intn(len(endpoints))
	selectedEndpoint := endpoints[randomEndpointIndex]

	client := &http.Client{Timeout: h.config.AppDefaultTimeout}

	req, err := http.NewRequest("GET", selectedEndpoint, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return err
	}
	defer response.Body.Close()

	var dadJokeResponse models.DadJokeResponse
	if err := json.NewDecoder(response.Body).Decode(&dadJokeResponse); err != nil {
		log.Println("Error reading response body:", err)
		return err
	}

	return c.JSON(dadJokeResponse)
}