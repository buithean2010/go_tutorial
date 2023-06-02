package controllers

import (
	"web-go-gin/dto"
	"web-go-gin/services"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginSvc services.LoginService
	jWtSvc   services.JWTService
}

func NewLoginController(loginSvc services.LoginService, jWtSvc services.JWTService) *loginController {
	return &loginController{
		loginSvc: loginSvc,
		jWtSvc:   jWtSvc,
	}
}

func (c *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := c.loginSvc.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return c.jWtSvc.CreateToken(credentials.Username, true)
	}
	return ""
}
