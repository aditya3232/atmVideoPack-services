package handler

import (
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
