package handler

import (
	"io"
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/constant"
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	"github.com/aditya3232/atmVideoPack-services.git/model/streaming_cctv"
	"github.com/gin-gonic/gin"
)

type StreamingCctvHandler struct {
	streamingCctvService streaming_cctv.Service
}

func NewStreamingCctvHandler(streamingCctvService streaming_cctv.Service) *StreamingCctvHandler {
	return &StreamingCctvHandler{streamingCctvService}
}

func (h *StreamingCctvHandler) StreamingCctv(c *gin.Context) {
	var input streaming_cctv.StreamingCctvInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err, " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	buffer, err := h.streamingCctvService.StreamingCctv(input)
	if err != nil {
		errorMessage := generateHTMLErrorMessageCctv()
		c.Data(http.StatusInternalServerError, "text/html", []byte(errorMessage))
		return
	}

	// Salin isi response body ke response context Gin
	c.Header("Content-Type", "text/html")
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

func generateHTMLErrorMessageCctv() string {
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
