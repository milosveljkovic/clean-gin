package server

import (
	"cleangin/internal/api/handlers"
	"cleangin/internal/constants"
	"cleangin/internal/routes/v1"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(
	userHandler *handlers.UserHandler,
) {
	engine := setupEngine()

	routes.CreateV1Routes(engine, userHandler)

	port := os.Getenv("CLEANGIN_PORT")
	if port == "" {
		port = constants.CleanGinDefaultPort
	}
	engine.Run(fmt.Sprintf(":%s", port))
}

func setupEngine() *gin.Engine {

	mode := gin.Mode()
	gin.SetMode(mode)

	engine := gin.New()
	engine.Use(cors.Default())
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	return engine
}

func getRouterMode() string {
	mode := os.Getenv("CLEANGIN_MODE")
	if mode == "" {
		mode = constants.CleanGinDefaultMode
	}
	return mode
}
