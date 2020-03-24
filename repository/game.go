package repository

import (
	"github.com/SnakeRoyale/snake-royale-backend/common"
	"github.com/SnakeRoyale/snake-royale-backend/model"
	"strconv"
)

type GameRepository struct {
	connections []model.Connection
	Status      string
	StatusCode  int
	TimeToStart int
}

func (r *GameRepository) StopConnection(tokenString string) error {
	token, err := strconv.ParseInt(tokenString, 10, 64)
	if err != nil {
		return err
	}
	for i, s := range r.connections {
		if s.Token == token {
			r.connections[i] = r.connections[len(r.connections)-1]
			r.connections = r.connections[:len(r.connections)-1]
			return nil
		}
	}
}

func (r *GameRepository) Authenticate(name string) int64 {
	newToken := common.CreateId()
	for _, s := range r.connections {
		if s.Token == newToken {
			return r.Authenticate(name)
		}
	}
	if len(r.connections) < 1 {
		// TODO: start game mechanics
	}
	r.connections = append(r.connections, model.Connection{
		Name:  name,
		Token: newToken,
	})
	return newToken
}

func NewGameRepository() *GameRepository {
	return &GameRepository{
		Status:      "Start",
		StatusCode:  200,
		TimeToStart: 0,
	}
}
