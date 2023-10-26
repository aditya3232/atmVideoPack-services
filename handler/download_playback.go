package handler

import (
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/constant"
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	"github.com/aditya3232/atmVideoPack-services.git/model/download_playback"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type DownloadPlaybackHandler struct {
	downloadPlaybackService download_playback.Service
}

func NewDownloadPlaybackHandler(downloadPlaybackService download_playback.Service) *DownloadPlaybackHandler {
	return &DownloadPlaybackHandler{downloadPlaybackService}
}

func (h *DownloadPlaybackHandler) DownloadPlayback(c *gin.Context) {
	var input download_playback.ServiceDownloadPlaybackInput

	// input from form-data
	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err, " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	DownloadPlaybacks, err := h.downloadPlaybackService.DownloadPlayback(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err, " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	if len(DownloadPlaybacks) == 0 {
		errorMessage := gin.H{"errors": "Video playback tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, errorMessage)
		log_function.Info("Entry tidak ditemukan", " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, download_playback.ServiceDownloadPlaybackFormatMany(DownloadPlaybacks))
	// response := helper.APIResponse(constant.DataFound, http.StatusOK, DownloadPlaybacks)
	c.JSON(response.Meta.Code, response)

}
