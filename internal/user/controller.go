package user

import (
	"initial_project_go/pkg/config"
	"initial_project_go/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

var errParams = config.GetMessage("error.params")

type UserController struct {
	service UserServiceinterface
}

func NewUserController(service UserServiceinterface) *UserController {
	return &UserController{service: service}
}

func (c *UserController) AddUsers(cgin *gin.Context) {
	request := new(RequestAddUsers)
	if err := cgin.ShouldBindJSON(request); err != nil {
		r := utils.Response{
			Context:    cgin,
			Request:    request,
			StatusCode: http.StatusBadRequest,
		}
		r.ErrorHandler(err)
		return
	}

	result, httpStatus, err := c.service.ServiceAddUsers(cgin.Request.Context(), request)
	if err != nil {
		r := utils.Response{
			Message:    err.Error(),
			Context:    cgin,
			StatusCode: httpStatus,
			Data:       result,
			Request:    request,
		}
		r.Error()
		return
	}

	r := utils.Response{
		Context:    cgin,
		Message:    "success",
		Data:       result,
		StatusCode: 200,
		Request:    request,
	}
	r.Success()
}

func (c *UserController) GetUsers(cgin *gin.Context) {
	request := new(RequestGetUsers)
	if err := cgin.ShouldBindJSON(request); err != nil {
		r := utils.Response{
			Context:    cgin,
			Request:    request,
			StatusCode: http.StatusBadRequest,
		}
		r.ErrorHandler(err)
		return
	}

	result, httpStatus, err := c.service.ServiceGetUsers(cgin.Request.Context(), request)
	if err != nil {
		r := utils.Response{
			Message:    err.Error(),
			Context:    cgin,
			StatusCode: httpStatus,
			Data:       result,
			Request:    request,
		}
		r.Error()
		return
	}

	r := utils.Response{
		Context:    cgin,
		Message:    "success",
		Data:       result,
		StatusCode: 200,
		Request:    request,
	}
	r.Success()
}
