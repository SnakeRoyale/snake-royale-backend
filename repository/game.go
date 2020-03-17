package repository

import "github.com/SnakeRoyale/snake-royale-backend/model"

type GameRepository struct {
	connections []model.Connection
	Status      string
	StatusCode  int
	TimeToStart int
}

func NewGameRepository() *GameRepository {
	return &GameRepository{
		Status:      "OK",
		StatusCode:  200,
		TimeToStart: 0,
	}
}
