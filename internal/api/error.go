package api

import (
	"gitee.com/plutoccc/devops_app/utils/errors"
)

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	err := errors.NewNotFound()
	this.ServeError(err)
}

func (this *ErrorController) Error405() {
	err := errors.NewMethodNotAllowed()
	this.ServeError(err)
}
