package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SnakeRoyale/snake-royale-backend/model"
	"github.com/SnakeRoyale/snake-royale-backend/repository"
	socketio "github.com/googollee/go-socket.io"
	"time"
)

type GameService struct {
	gameRepository *repository.GameRepository
	isGameRunning  bool
	server         *socketio.Server
	gameInterval chan bool
}

func (s GameService) RegisterEvents(server *socketio.Server) {
	s.server = server
	server.OnEvent("/login", "login", func(socket socketio.Conn, msg string) {
		conn, err := s.authenticate(msg)
		if err != nil {
			socket.Emit("Error in login: " + err.Error())
		}
		socket.Emit("Logged in", conn)
	})

	server.OnEvent("/game/status", "checkGameStatus", func(socket socketio.Conn, msg string) {
		json, err := s.checkGameStatus()
		if err != nil {
			socket.Emit("Error in checkGameStatus: " + err.Error())
		}
		socket.Emit("Game Status", json)
	})

	server.OnEvent("/game/leave", "leaveGame", func(socket socketio.Conn, msg string) {
		response, err := s.authenticate(msg)
		if err != nil {
			socket.Emit("Error in leaveGame: " + err.Error())
		}
		socket.Emit("Leave Game", response)
	})
}

func (s GameService) authenticate(msg string) (string, error) {
	data := model.PostConnection{}

	if err := json.Unmarshal([]byte(msg), data); err != nil {
		return "", err
	} else if len(data.Name) <= 1 {
		return "", errors.New("invalid name")
	} else if data.Token != "" {
		return "", errors.New("already authenticated")
	}
	byteArrayJson, err := json.Marshal(model.Connection{
		Name:  data.Name,
		Token: s.gameRepository.Authenticate(data.Name),
	})
	if err != nil {
		return "", err
	}
	if !s.isGameRunning {
		go s.startGame()
	}
	return string(byteArrayJson), err
}

func (s GameService) checkGameStatus() (string, error) {
	byteArrayJson, err := json.Marshal(model.StatusResponse{
		Status:      s.gameRepository.Status,
		StatusCode:  s.gameRepository.StatusCode,
		TimeToStart: s.gameRepository.TimeToStart,
	})
	if err != nil {
		return "", err
	}
	return string(byteArrayJson), err
}

func (s GameService) leaveGame(msg string) (string, error) {
	data := model.PostConnection{}

	if err := json.Unmarshal([]byte(msg), data); err != nil {
		return "", err
	} else if data.Token == "" {
		return "", errors.New("missing token")
	}

	if err := s.gameRepository.StopConnection(data.Token); err != nil {
		return "", errors.New("error while deleting token")
	}
	return "successfully deleted" + data.Name, nil

}

func (s GameService) sendSnakes(msg string) (string, error) {
	data := model.PostConnection{}

	if err := json.Unmarshal([]byte(msg), data); err != nil {
		return "", err
	} else if data.Token == "" {
		return "", errors.New("missing token")
	}

	if err := s.gameRepository.StopConnection(data.Token); err != nil {
		return "", errors.New("error while deleting token")
	}
	return "successfully deleted" + data.Name, nil

}

func (s GameService) startGame() {
	s.isGameRunning = true
	s.gameInterval := s.setInterval(s.updateGame, 1000)
	// TODO: Interval not starting
}

func (s GameService) updateGame() {
	s.setInterval()
	s.isGameRunning = true
	for s.isGameRunning {

	}
}

func NewGame(gameRepository *repository.GameRepository) *GameService {
	return &GameService{
		gameRepository: gameRepository,
	}
}
func (s GameService) setInterval(someFunc func(), milliseconds int) chan bool {
	interval := time.Duration(milliseconds) * time.Millisecond
	ticker := time.NewTicker(interval)
	clear := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				go someFunc()
			case <-clear:
				ticker.Stop()
				return
			}
		}
	}()
	return clear

}
