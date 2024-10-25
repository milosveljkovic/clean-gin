package handlers

import (
	"cleangin/internal/app"
	"cleangin/internal/e"
	"cleangin/internal/models"
	"cleangin/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserSvc service.UserSvcI
}

func NewUserHandler(userSvc service.UserSvcI) *UserHandler {
	return &UserHandler{UserSvc: userSvc}
}

func (h UserHandler) Register(ctx *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	appG := app.Gin{C: ctx}

	httpCode, errCode := app.BindAndValid(ctx, &body)
	if errCode != e.SUCCESS {
		appG.FailResponse(httpCode, errCode)
		return
	}

	user := models.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}

	err := h.UserSvc.RegisterUser(user)
	if err != nil {
		appG.FailResponse(http.StatusInternalServerError, e.ERROR)
		return
	}
	appG.Response(http.StatusCreated, e.SUCCESS, nil)
}
