package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	"github.com/aditya3232/atmVideoPack-services.git/model/auth"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthHandler struct {
	authService auth.Service
}

func NewAuthHandler(authService auth.Service) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input auth.LoginInput

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := fmt.Sprintf("Login gagal")
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	token, err := h.authService.Login(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := fmt.Sprintf("Login gagal")
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := fmt.Sprintf("Login berhasil")
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	newToken := token.RememberToken
	expires := time.Now().AddDate(0, 0, 30)
	response := helper.APIResponse(message, http.StatusOK, auth.LoginFormat(newToken, expires))
	c.JSON(response.Meta.Code, response)
}
