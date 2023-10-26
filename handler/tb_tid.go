package handler

import (
	"fmt"
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/constant"
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	"github.com/aditya3232/atmVideoPack-services.git/model/tb_tid"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type TbTidHandler struct {
	tbTidService tb_tid.Service
}

func NewTbTidHandler(tbTidService tb_tid.Service) *TbTidHandler {
	return &TbTidHandler{tbTidService}
}

func (h *TbTidHandler) GetAllTbEntries(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, tb_tid.TbTid{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	TbTids, page, err := h.tbTidService.GetAll(filter, page, sort)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIDataTableResponse(constant.CannotProcessRequest, http.StatusBadRequest, helper.Pagination{}, errorMessage)
		log_function.Error(err, " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	if len(TbTids) == 0 {
		errorMessage := gin.H{"errors": "Entry tidak ditemukan"}
		response := helper.APIDataTableResponse(constant.DataNotFound, http.StatusNotFound, helper.Pagination{}, errorMessage)
		log_function.Info("Entry tidak ditemukan", " from ip address: ", c.ClientIP())
		c.JSON(response.Meta.Code, response)
		return
	}

	response := helper.APIDataTableResponse(constant.DataFound, http.StatusOK, page, tb_tid.TbTidGetAllFormat(TbTids))
	c.JSON(response.Meta.Code, response)
}

func (h *TbTidHandler) CreateTbTid(c *gin.Context) {
	var input tb_tid.TbTidCreateInput

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

	newTbTid, err := h.tbTidService.Create(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessMessage, http.StatusOK, tb_tid.TbTidCreateFormat(newTbTid))
	log_function.Info("Tid berhasil dibuat")
	c.JSON(http.StatusOK, response)
}

func (h *TbTidHandler) GetOneByID(c *gin.Context) {
	var input tb_tid.GetOneByIDInput

	fmt.Println(input)

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

	getOneByID, err := h.tbTidService.GetOneByID(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse(constant.CannotProcessRequest, http.StatusBadRequest, errorMessage)
		log_function.Error(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(constant.SuccessMessage, http.StatusOK, tb_tid.TbTidCreateFormat(getOneByID))
	log_function.Info("Tid berhasil ditemukan")
	c.JSON(http.StatusOK, response)
}
