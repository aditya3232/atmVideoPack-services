package handler

import (
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/constant"
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	"github.com/aditya3232/atmVideoPack-services.git/model/profile"
	"github.com/aditya3232/atmVideoPack-services.git/model/users"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	profileService profile.Service
}

func NewProfileHandler(profileService profile.Service) *ProfileHandler {
	return &ProfileHandler{profileService}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID := c.MustGet("currentUser").(users.Users).ID

	profileData, err := h.profileService.GetProfile(userID)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	user := profileData.User
	role := profileData.Role
	permissions := profileData.Permissions

	endpoint := c.Request.URL.Path
	message := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, profile.ProfileGetFormat(user, role, permissions))
	c.JSON(response.Meta.Code, response)

}
