package di

import (
	server "cleangin/internal/api"
	"cleangin/internal/api/handlers"
	"cleangin/internal/db"
	"cleangin/internal/repo"
	"cleangin/internal/service"
)

func InitialiseAPI() {
	DB := db.InitDB()

	userRepository := repo.NewUserRepo(DB)
	userSvc := service.NewUserSvc(userRepository)
	userHandler := handlers.NewUserHandler(userSvc)

	server.StartServer(
		userHandler,
	)
}
