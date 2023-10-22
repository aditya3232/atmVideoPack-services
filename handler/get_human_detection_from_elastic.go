package handler

import (
	"net/http"

	"github.com/aditya3232/gatewatchApp-services.git/constant"
	"github.com/aditya3232/gatewatchApp-services.git/helper"
	"github.com/aditya3232/gatewatchApp-services.git/log"
	"github.com/aditya3232/gatewatchApp-services.git/model/get_human_detection_from_elastic"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type GetHumanDetectionFromElasticHandler struct {
	getHumanDetectionFromElasticService get_human_detection_from_elastic.Service
}

func NewGetHumanDetectionFromElasticHandler(getHumanDetectionFromElasticService get_human_detection_from_elastic.Service) *GetHumanDetectionFromElasticHandler {
	return &GetHumanDetectionFromElasticHandler{getHumanDetectionFromElasticService}
}

func (h *GetHumanDetectionFromElasticHandler) FindAll(c *gin.Context) {
	var input get_human_detection_from_elastic.FindAllElasticHumanDetectionInput

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

	elasticHumanDetections, err := h.getHumanDetectionFromElasticService.FindAll(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log.Error(err, " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	if len(elasticHumanDetections) == 0 {
		errorMessage := gin.H{"errors": "Entry tidak ditemukan"}
		response := helper.APIResponse(constant.DataNotFound, http.StatusNotFound, errorMessage)
		log.Info("Entry tidak ditemukan", " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIResponse(constant.DataFound, http.StatusOK, get_human_detection_from_elastic.ElasticHumanDetectionGetAllFormat(elasticHumanDetections))
	c.JSON(response.Meta.Code, response)
}
