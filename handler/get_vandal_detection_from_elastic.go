package handler

import (
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/constant"
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	"github.com/aditya3232/atmVideoPack-services.git/model/get_vandal_detection_from_elastic"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type GetVandalDetectionFromElasticHandler struct {
	getVandalDetectionFromElasticService get_vandal_detection_from_elastic.Service
}

func NewGetVandalDetectionFromElasticHandler(getVandalDetectionFromElasticService get_vandal_detection_from_elastic.Service) *GetVandalDetectionFromElasticHandler {
	return &GetVandalDetectionFromElasticHandler{getVandalDetectionFromElasticService}
}

func (h *GetVandalDetectionFromElasticHandler) FindAll(c *gin.Context) {
	var input get_vandal_detection_from_elastic.FindAllElasticVandalDetectionInput

	// input from form-data
	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.InvalidRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	elasticVandalDetections, err := h.getVandalDetectionFromElasticService.FindAll(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err, " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	if len(elasticVandalDetections) == 0 {
		errorMessage := gin.H{"errors": "Entry tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, errorMessage)
		log_function.Info("Entry tidak ditemukan", " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, get_vandal_detection_from_elastic.ElasticVandalDetectionGetAllFormat(elasticVandalDetections))
	c.JSON(response.Meta.Code, response)
}
