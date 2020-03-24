package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Registerable interface {
	RegisterRoutes(engine *gin.Engine)
}

func New(registerables ...Registerable) *gin.Engine {
	ginEngine := gin.New()

	ginEngine.Use(gin.Recovery())

	// Allow cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization", "*")
	corsConfig.AddAllowMethods("HEAD", "GET", "POST", "PUT", "DELETE")
	ginEngine.Use(cors.New(corsConfig))

	//RegisterRoutes
	for _, registerable := range registerables {
		registerable.RegisterRoutes(ginEngine)
	}

	return ginEngine
}
