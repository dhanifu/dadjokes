package models

type Jokes struct {
	Jokes []string `json:"jokes"`
}

type DadJokeResponse struct {
	Data string `json:"data"`
}