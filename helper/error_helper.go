package helper

import (
	"errors"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vincen320/product-service-mongodb/exception"
	"github.com/vincen320/product-service-mongodb/model/web"
)

func ReturnError(c *gin.Context, e error) {
	var (
		badRequest      *exception.BadRequestError
		notFound        *exception.NotFoundError
		unauthorized    *exception.UnauthorizedError
		validationError validator.ValidationErrors
	)

	log.Println(reflect.TypeOf(e))
	if errors.As(e, &badRequest) {
		ErrorResponse(c, http.StatusBadRequest, badRequest.Error())
	} else if errors.As(e, &notFound) {
		ErrorResponse(c, http.StatusNotFound, notFound.Error())
	} else if errors.As(e, &validationError) {
		ErrorResponse(c, http.StatusBadRequest, validationError.Error())
	} else if errors.As(e, &unauthorized) {
		ErrorResponse(c, http.StatusUnauthorized, unauthorized.Error())
	} else {
		ErrorResponse(c, http.StatusInternalServerError, e.Error())
	}
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, web.WebResponse{
		Code:    code,
		Message: message,
	})
}
