package data

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/dhanifu/dadjokes/config"
	"github.com/dhanifu/dadjokes/models"
)

type data struct {
	config *config.Configuration
}

func InitData(config *config.Configuration) data {
	return data{
		config: config,
	}
}

func GetDadJokesData() models.Jokes {
	file, err := os.Open("./storage/jokes.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	var jokes models.Jokes
	if err = json.Unmarshal(byteValue, &jokes); err != nil {
		log.Fatalf("Failed to unmarshal JSON data: %v", err)
	}

	return jokes
}