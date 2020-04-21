package service

import (
	"github.com/SnakeRoyale/snake-royale-backend/model"
	"github.com/SnakeRoyale/snake-royale-backend/repository"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"net/http"
)

type SnakeService struct {
	snakeRepository *repository.SnakeRepository
	server *socketio.Server
}

func (s SnakeService) RegisterRoutes(server *socketio.Server) {
	s.server = server
	server.OnEvent("/snake", "sendSnakePos", func(socket socketio.Conn, msg string) {
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

func (s SnakeService) authenticate(c *gin.Context) {
	// send pos
	// get other player
	// spawn fruit
	// change color
	// loading screen
}


func (s SnakeService) leaveGame(c *gin.Context) {

}

func NewSnake(gameRepository *repository.SnakeRepository) *SnakeService {
	return &SnakeService{
		snakeRepository: gameRepository,
	}
}
