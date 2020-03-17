package service

import (
	"github.com/SnakeRoyale/snake-royale-backend/model"
	"github.com/SnakeRoyale/snake-royale-backend/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GameService struct {
	gameRepository *repository.GameRepository
}

func (s GameService) RegisterRoutes(ginEngine *gin.Engine) {

	privateGroup := ginEngine.Group("/api")

	privateGroup.POST("/login", s.authenticate)
	privateGroup.GET("game/status", s.checkGameStatus)
	privateGroup.DELETE("game", s.leaveGame)
}

func (s GameService) authenticate(c *gin.Context) {
// send pos
// get other player
// spawn fruit
// change color
// loading screen
}

func (s GameService) checkGameStatus(c *gin.Context) {
	c.JSON(http.StatusOK, model.StatusResponse{
		Status:      s.gameRepository.Status,
		StatusCode:  s.gameRepository.StatusCode,
		TimeToStart: s.gameRepository.TimeToStart,
	})
}

func (s GameService) leaveGame(c *gin.Context) {

}

func NewGame(gameRepository *repository.GameRepository) *GameService {
	return &GameService{
		gameRepository: gameRepository,
	}
}
