package handler

import (
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/constant"
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	users_model "github.com/aditya3232/atmVideoPack-services.git/model/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

/*
	- dihandler akan menampilkan error untuk frontend, dan log juga ke elastic
	- namun di service berikan error asli, dan log juga ke elastic
	- karena user gk perlu tau error

	- time="2023-12-07T11:19:23+07:00" level=error msg="username must unique"
	- time="2023-12-07T11:19:23+07:00" level=error msg="Failed to create new data. (user)"
*/

type UsersHandler struct {
	usersService users_model.Service
}

func NewUsersHandler(usersService users_model.Service) *UsersHandler {
	return &UsersHandler{usersService}
}

func (h *UsersHandler) GetAll(c *gin.Context) {
	filter := helper.QueryParamsToMap(c, users_model.Users{})
	page := helper.NewPagination(helper.StrToInt(c.Query("page")), helper.StrToInt(c.Query("limit")))
	sort := helper.NewSort(c.Query("sort"), c.Query("order"))

	users, page, err := h.usersService.GetAll(filter, page, sort)
	if err != nil {
		// errors := helper.FormatError(err)
		// errorMessage := gin.H{"errors": errors}
		endpoint := c.Request.URL.Path
		message := constant.DataNotFound
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		log_function.Error(message, err, endpoint, errorCode, ipAddress)

		response := helper.APIDataTableResponse(message, http.StatusNotFound, helper.Pagination{}, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	if len(users) == 0 {
		// errorMessage := gin.H{"errors": "Entry tidak ditemukan"}
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		log_function.Error(message, err, endpoint, errorCode, ipAddress)

		response := helper.APIDataTableResponse(message, http.StatusNotFound, helper.Pagination{}, nil)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessCreateData
	infoCode := http.StatusCreated
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIDataTableResponse(message, http.StatusOK, page, users_model.UsersGetAllFormat(users))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) GetOne(c *gin.Context) {
	var input users_model.UsersGetOneByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		endpoint := c.Request.URL.Path
		errors := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	user, err := h.usersService.GetOne(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		errors := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusNotFound, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	if user.ID == 0 {
		endpoint := c.Request.URL.Path
		errors := constant.DataNotFound
		errorCode := http.StatusNotFound
		ipAddress := c.ClientIP()
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusNotFound, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	info := constant.DataFound
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(info, endpoint, infoCode, ipAddress)

	response := helper.APIResponse(info, http.StatusOK, users_model.UsersGetFormat(user))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) Create(c *gin.Context) {
	var input users_model.UsersInput

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errorMessage := gin.H{"message": message}
		log_function.Error(message, err, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	newUser, err := h.usersService.Create(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		message := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errorMessage := gin.H{"message": message}
		log_function.Error(message, err, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(message, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	message := constant.SuccessCreateData
	infoCode := http.StatusCreated
	ipAddress := c.ClientIP()
	log_function.Info(message, "", endpoint, infoCode, ipAddress)

	response := helper.APIResponse(message, http.StatusCreated, users_model.UsersCreateFormat(newUser))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) Update(c *gin.Context) {
	var id users_model.UsersGetOneByIdInput
	var input users_model.UsersUpdateInput

	err := c.ShouldBindUri(&id)
	if err != nil {
		endpoint := c.Request.URL.Path
		errors := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	input.ID = id.ID

	err = c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := c.Request.URL.Path
		errors := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	newUser, err := h.usersService.Update(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		errors := constant.FailedUpdateData
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	info := constant.SuccessUpdateData
	infoCode := http.StatusOK
	ipAddress := c.ClientIP()
	log_function.Info(info, endpoint, infoCode, ipAddress)

	response := helper.APIResponse(info, http.StatusOK, users_model.UsersUpdateFormat(newUser))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) Delete(c *gin.Context) {
	var input users_model.UsersGetOneByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		endpoint := c.Request.URL.Path
		errors := constant.InvalidRequest
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	err = h.usersService.Delete(input)
	if err != nil {
		endpoint := c.Request.URL.Path
		errors := constant.FailedDeleteData
		errorCode := http.StatusBadRequest
		ipAddress := c.ClientIP()
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := c.Request.URL.Path
	info := constant.SuccessDeleteData
	infoCode := http.StatusNoContent
	ipAddress := c.ClientIP()
	log_function.Info(info, endpoint, infoCode, ipAddress)

	response := helper.APIResponse(info, http.StatusNoContent, nil)
	c.JSON(response.Meta.Code, response)
}
