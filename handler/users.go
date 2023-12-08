package handler

import (
	"fmt"
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
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/getall")
		errors := constant.DataNotFound
		errorCode := fmt.Sprintf("Status: %d", http.StatusBadRequest)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIDataTableResponse(errors, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	if len(users) == 0 {
		// errorMessage := gin.H{"errors": "Entry tidak ditemukan"}
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/getall")
		errors := constant.DataNotFound
		errorCode := fmt.Sprintf("Status: %d", http.StatusNotFound)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIDataTableResponse(errors, http.StatusNotFound, helper.Pagination{}, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/getall")
	info := constant.DataFound
	infoCode := fmt.Sprintf("Status: %d", http.StatusOK)
	ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
	log_function.Info(info, endpoint, infoCode, ipAddress)

	response := helper.APIDataTableResponse(info, http.StatusOK, page, users_model.UsersGetAllFormat(users))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) GetOne(c *gin.Context) {
	var input users_model.UsersGetOneByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/getone/:id")
		errors := constant.InvalidRequest
		errorCode := fmt.Sprintf("Status: %d", http.StatusBadRequest)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	user, err := h.usersService.GetOne(input)
	if err != nil {
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/getone/:id")
		errors := constant.DataNotFound
		errorCode := fmt.Sprintf("Status: %d", http.StatusNotFound)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusNotFound, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	if user.ID == 0 {
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/getone/:id")
		errors := constant.DataNotFound
		errorCode := fmt.Sprintf("Status: %d", http.StatusNotFound)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusNotFound, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/getone/:id")
	info := constant.DataFound
	infoCode := fmt.Sprintf("Status: %d", http.StatusOK)
	ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
	log_function.Info(info, endpoint, infoCode, ipAddress)

	response := helper.APIResponse(info, http.StatusOK, users_model.UsersGetFormat(user))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) Create(c *gin.Context) {
	var input users_model.UsersInput

	err := c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/create")
		errors := constant.InvalidRequest
		errorCode := fmt.Sprintf("Status: %d", http.StatusBadRequest)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	newUser, err := h.usersService.Create(input)
	if err != nil {
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/create")
		errors := constant.FailedCreateData
		errorCode := fmt.Sprintf("Status: %d", http.StatusBadRequest)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/create")
	info := constant.SuccessCreateData
	infoCode := fmt.Sprintf("Status: %d", http.StatusCreated)
	ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
	log_function.Info(info, endpoint, infoCode, ipAddress)

	response := helper.APIResponse(info, http.StatusCreated, users_model.UsersCreateFormat(newUser))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) Update(c *gin.Context) {
	var id users_model.UsersGetOneByIdInput
	var input users_model.UsersUpdateInput

	err := c.ShouldBindUri(&id)
	if err != nil {
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/update/:id")
		errors := constant.InvalidRequest
		errorCode := fmt.Sprintf("Status: %d", http.StatusBadRequest)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	input.ID = id.ID

	err = c.ShouldBindWith(&input, binding.Form)
	if err != nil {
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/update/:id")
		errors := constant.InvalidRequest
		errorCode := fmt.Sprintf("Status: %d", http.StatusBadRequest)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	newUser, err := h.usersService.Update(input)
	if err != nil {
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/update/:id")
		errors := constant.FailedUpdateData
		errorCode := fmt.Sprintf("Status: %d", http.StatusBadRequest)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/update/:id")
	info := constant.SuccessUpdateData
	infoCode := fmt.Sprintf("Status: %d", http.StatusOK)
	ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
	log_function.Info(info, endpoint, infoCode, ipAddress)

	response := helper.APIResponse(info, http.StatusOK, users_model.UsersUpdateFormat(newUser))
	c.JSON(response.Meta.Code, response)
}

func (h *UsersHandler) Delete(c *gin.Context) {
	var input users_model.UsersGetOneByIdInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/delete/:id")
		errors := constant.InvalidRequest
		errorCode := fmt.Sprintf("Status: %d", http.StatusBadRequest)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	err = h.usersService.Delete(input)
	if err != nil {
		endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/delete/:id")
		errors := constant.FailedDeleteData
		errorCode := fmt.Sprintf("Status: %d", http.StatusBadRequest)
		ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
		errorMessage := gin.H{"errors": errors}
		log_function.Error(errors, endpoint, errorCode, ipAddress)

		response := helper.APIResponse(errors, http.StatusBadRequest, errorMessage)
		c.JSON(response.Meta.Code, response)
		return
	}

	endpoint := fmt.Sprintf("endpoint: %s", "/api/atmvideopack/v1/users/delete/:id")
	info := constant.SuccessDeleteData
	infoCode := fmt.Sprintf("Status: %d", http.StatusNoContent)
	ipAddress := fmt.Sprintf("from ip address: %s", c.ClientIP())
	log_function.Info(info, endpoint, infoCode, ipAddress)

	response := helper.APIResponse(info, http.StatusNoContent, nil)
	c.JSON(response.Meta.Code, response)
}
