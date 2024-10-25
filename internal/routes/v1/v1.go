package routes

import (
	"cleangin/internal/api/handlers"
	"cleangin/internal/api/middleware"

	"github.com/gin-gonic/gin"
)

func CreateV1Routes(
	r *gin.Engine,
	userHandler *handlers.UserHandler,
) {
	// auth routes for sing-in/sing-up process
	authv1 := r.Group("/api/v1/auth")
	authv1.POST("register", userHandler.Register)

	// routes protected with middleware.JwtVerify
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JwtVerify)
	// define your protected routes
	// authv1.POST("protectedroute", userHandler.ProtectedRoute)
}
