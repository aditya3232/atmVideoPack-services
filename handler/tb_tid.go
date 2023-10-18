package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aditya3232/gatewatchApp-services.git/constant"
	"github.com/aditya3232/gatewatchApp-services.git/helper"
	"github.com/aditya3232/gatewatchApp-services.git/log"
	"github.com/aditya3232/gatewatchApp-services.git/model/tb_tid"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type TbTidHandler struct {
	tbTidService tb_tid.Service
}

func NewTbTidHandler(tbTidService tb_tid.Service) *TbTidHandler {
	return &TbTidHandler{tbTidService}
}

func (h *TbTidHandler) CreateTbTid(c *gin.Context) {
	var input tb_tid.TbTidCreateInput

	// input from form-data
	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newTbTid, err := h.tbTidService.Create(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessMessage, http.StatusOK, tb_tid.TbTidCreateFormat(newTbTid))
	log.Info("Tid berhasil dibuat")
	c.JSON(http.StatusOK, response)
}

func (h *TbTidHandler) GetOneByID(c *gin.Context) {
	var input tb_tid.GetOneByIDInput

	// input from form-data
	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getOneByID, err := h.tbTidService.GetOneByID(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessMessage, http.StatusOK, tb_tid.TbTidCreateFormat(getOneByID))
	log.Info("Tid berhasil ditemukan")
	c.JSON(http.StatusOK, response)
}

// get stream video
func (h *TbTidHandler) GetStreamVideo(c *gin.Context) {
	var input tb_tid.GetOneByIDInput

	// input from form-data
	// err := c.ShouldBindWith(&input, binding.Form)

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	getOneByID, err := h.tbTidService.GetOneByID(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// get ip address
	ipAddress := getOneByID.IpAddress
	streamingURL := fmt.Sprintf("http://%s:5000", ipAddress)

	// Ambil video dari URL pihak ketiga
	response, err := http.Get(streamingURL)
	if err != nil {
		log.Error(err)
		// Tampilkan pesan kesalahan HTML sebagai respons
		errorMessage := generateHTMLErrorMessage()
		// Mengirimkan respons dengan HTML
		c.Data(http.StatusInternalServerError, "text/html", []byte(errorMessage))
		return
	}

	defer response.Body.Close()

	contentType := response.Header.Get("Content-Type")
	c.Header("Content-Type", contentType)

	// Salin isi response body ke response context Gin
	_, err = io.Copy(c.Writer, response.Body)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

}

func generateHTMLErrorMessage() string {
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
