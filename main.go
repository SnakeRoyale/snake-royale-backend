package main

import (
	"fmt"
	"github.com/SnakeRoyale/snake-royale-backend/repository"
	"github.com/SnakeRoyale/snake-royale-backend/service"
	"github.com/SnakeRoyale/snake-royale-backend/config"
)

func main() {
	Register()

}

func Register() {
	// init repositories & facades
	gameRepository := repository.NewGameRepository()

	// api for frontend
	gameApi := service.NewGame(gameRepository)

	ginEngine := config.New(gameApi,)

	if err := ginEngine.Run(); err != nil {
		fmt.Println("Failed to start webservice: ", err)
	}
}