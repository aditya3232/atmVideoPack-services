package handler

import (
	"fmt"
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/constant"
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	roles_model "github.com/aditya3232/atmVideoPack-services.git/model/roles"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RolesHandler struct {
	rolesService roles_model.Service
}

func NewRolesHandler(rolesService roles_model.Service) *RolesHandler {
	return &RolesHandler{rolesService}
}

func (h *RolesHandler) GetAll(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, roles_model.Roles{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	roles, page, err := h.rolesService.GetAll(filter, page, sort)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIDataTableResponse(message, http.StatusNotFound, helper.Pagination{}, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	if len(roles) == 0 {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := fmt.Sprintf("Users not found")
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIDataTableResponse(message, http.StatusNotFound, helper.Pagination{}, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIDataTableResponse(message, http.StatusOK, page, roles_model.RolesGetAllFormat(roles))
	c.JSON(response.Meta.Code, response)
}

func (h *RolesHandler) GetOne(c *gin.Context) {
	var input roles_model.RolesGetOneByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	role, err := h.rolesService.GetOne(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusNotFound, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	if role.ID == 0 {
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errors := fmt.Sprintf("Roles not found")
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusNotFound, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, roles_model.RolesGetFormat(role))
	c.JSON(response.Meta.Code, response)
}

func (h *RolesHandler) Create(c *gin.Context) {
	var input roles_model.RolesInput

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	newRole, err := h.rolesService.Create(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessCreateData
	infoCode := http.StatusCreated
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusCreated, roles_model.RolesCreateFormat(newRole))
	c.JSON(response.Meta.Code, response)
}

func (h *RolesHandler) Update(c *gin.Context) {
	var id roles_model.RolesGetOneByIdInput
	var input roles_model.RolesUpdateInput

	err := c.ShouldBindUri(&id)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	input.ID = id.ID

	err = c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	newRole, err := h.rolesService.Update(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.FailedUpdateData
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessUpdateData
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusOK, roles_model.RolesUpdateFormat(newRole))
	c.JSON(response.Meta.Code, response)
}

func (h *RolesHandler) Delete(c *gin.Context) {
	var input roles_model.RolesGetOneByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	err = h.rolesService.Delete(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.FailedDeleteData
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errors := helper.FormatError(err)
		log_function.Error(message, errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, nil)
		c.JSON(response.Meta.Code, response)
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessDeleteData
	infoCode := http.StatusNoContent
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusNoContent, nil)
	c.JSON(response.Meta.Code, response)
}
