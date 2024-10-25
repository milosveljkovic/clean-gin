package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func JwtVerify(ctx *gin.Context) {
	fmt.Println("Some logic for jwt-middleware...")
}
