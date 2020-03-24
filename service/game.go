package service

import (
	"github.com/SnakeRoyale/snake-royale-backend/common"
	"github.com/SnakeRoyale/snake-royale-backend/model"
	"github.com/SnakeRoyale/snake-royale-backend/repository"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	var connection model.PostConnection
	errBinding := c.ShouldBindWith(&connection, binding.JSON)

	if errBinding != nil {
		common.WriteFailApiResponse(c, http.StatusBadRequest, "Error while parsing JSON body")
		return
	} else if connection.Token != "" {
		if err := s.gameRepository.StopConnection(connection.Token); err != nil {
			common.WriteFailApiResponse(c, http.StatusBadRequest, "Error while deleting old token")
			return
		}
	}
	token := s.gameRepository.Authenticate(connection.Name)

	c.JSON(http.StatusCreated, model.Connection{
		Name:  connection.Name,
		Token: token,
	})
}

func (s GameService) checkGameStatus(c *gin.Context) {
	c.JSON(http.StatusOK, model.StatusResponse{
		Status:      s.gameRepository.Status,
		StatusCode:  s.gameRepository.StatusCode,
		TimeToStart: s.gameRepository.TimeToStart,
	})
}

func (s GameService) leaveGame(c *gin.Context) {
	var connection model.PostConnection
	errBinding := c.ShouldBindWith(&connection, binding.JSON)

	if errBinding != nil {
		common.WriteFailApiResponse(c, http.StatusBadRequest, "Error while parsing JSON body")
	} else if connection.Token == "" {
		common.WriteFailApiResponse(c, http.StatusBadRequest, "missing token")
	} else {
		if err := s.gameRepository.StopConnection(connection.Token); err != nil {
			common.WriteFailApiResponse(c, http.StatusBadRequest, "Error while deleting token")
			return
		} else {
			common.WriteOKApiResponse(c, http.StatusOK, "successfully deleted" + connection.Name)
		}
	}
}

func NewGame(gameRepository *repository.GameRepository) *GameService {
	return &GameService{
		gameRepository: gameRepository,
	}
}
