package handler

import (
	"io"
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/constant"
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	"github.com/aditya3232/atmVideoPack-services.git/model/download_playback"
	"github.com/gin-gonic/gin"
)

type DownloadPlaybackHandler struct {
	downloadPlaybackService download_playback.Service
}

func NewDownloadPlaybackHandler(downloadPlaybackService download_playback.Service) *DownloadPlaybackHandler {
	return &DownloadPlaybackHandler{downloadPlaybackService}
}

func (h *DownloadPlaybackHandler) DownloadPlayback(c *gin.Context) {
	var input download_playback.ServiceDownloadPlaybackInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err, " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	buffer, err := h.downloadPlaybackService.DownloadPlayback(input)
	if err != nil {
		errorMessage := generateHTMLErrorMessageDownloadPlayback()
		c.Data(http.StatusInternalServerError, "text/html", []byte(errorMessage))
		return
	}

	// Salin isi response body ke response context Gin
	_, err = io.Copy(c.Writer, buffer)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

}

func generateHTMLErrorMessageDownloadPlayback() string {
	errorMessage := `
    <!DOCTYPE html>
	<html>
	<head>
		<title>Terjadi Kesalahan Server</title>
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
	</head>
	<body>
		<div class="container">
			<div class="row">
				<div class="col-12 text-center mt-5">
					<h1 class="display-4">Terjadi Kesalahan Server</h1>
					<p class="lead">Maaf, terjadi kesalahan saat memproses permintaan Anda.</p>
				</div>
			</div>
		</div>
	</body>
	</html>

	`

	return errorMessage
}
