// Common tools and helper functions
package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
)

type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs, result := err.(validator.ValidationErrors)
	if result {
		for _, v := range errs {
			if v.Param != "" {
				res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
			} else {
				res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
			}

		}
	} else {
		res = NewError("error", err)
	}
	return res
}

func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}
