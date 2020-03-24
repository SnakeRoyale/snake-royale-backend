package service

import (
	"github.com/SnakeRoyale/snake-royale-backend/model"
	"github.com/SnakeRoyale/snake-royale-backend/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SnakeService struct {
	snakeRepository *repository.SnakeRepository
}

func (s SnakeService) RegisterRoutes(ginEngine *gin.Engine) {

	privateGroup := ginEngine.Group("/api")

	privateGroup.POST("/snake", s.authenticate)
	privateGroup.GET("snake/status", s.checkGameStatus)
	privateGroup.DELETE("snake", s.leaveGame)
}

func (s SnakeService) authenticate(c *gin.Context) {
	// send pos
	// get other player
	// spawn fruit
	// change color
	// loading screen
}

func (s SnakeService) checkGameStatus(c *gin.Context) {
	c.JSON(http.StatusOK, model.StatusResponse{
		Status:      s.snakeRepository.Status,
		StatusCode:  s.snakeRepository.StatusCode,
		TimeToStart: s.snakeRepository.TimeToStart,
	})
}

func (s SnakeService) leaveGame(c *gin.Context) {

}

func NewSnake(gameRepository *repository.SnakeRepository) *SnakeService {
	return &SnakeService{
		snakeRepository: gameRepository,
	}
}
