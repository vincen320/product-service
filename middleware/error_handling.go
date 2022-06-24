package middleware

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vincen320/product-service/exception"
	"github.com/vincen320/product-service/model/web"
)

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				log.Println(reflect.TypeOf(err), " ", err) //debugging
				switch err.(type) {
				case *exception.BadRequestError: //jangan lupa dicek dalam bentuk pointer
					BadRequestErrResponse(c, err)
				case *exception.NotFoundError: //jangan lupa dicek dalam bentuk pointer
					NotFoundErrResponse(c, err)
				case *exception.UnauthorizedError: //jangan lupa dicek dalam bentuk pointer
					UnauthorizedErrResponse(c, err)
				case validator.ValidationErrors:
					ValidationErrorResponse(c, err)
				default:
					InternalServerError(c, err)
				}
			}
		}()
		c.Next()
	}
}

func BadRequestErrResponse(c *gin.Context, err interface{}) {
	badRequest, _ := err.(*exception.BadRequestError) //diset ke bentuk pointer
	c.JSON(http.StatusBadRequest, web.WebResponse{
		Status:  http.StatusBadRequest,
		Message: badRequest.Error(),
	})
}

func NotFoundErrResponse(c *gin.Context, err interface{}) {
	notFound, _ := err.(*exception.NotFoundError) //diset ke bentuk pointer
	c.JSON(http.StatusNotFound, web.WebResponse{
		Status:  http.StatusNotFound,
		Message: notFound.Error(),
	})
}

func UnauthorizedErrResponse(c *gin.Context, err interface{}) {
	unauthorized, _ := err.(*exception.UnauthorizedError) //diset ke bentuk pointer
	c.Header("WWW-Authenticate", "not authorized user")   //jelaskan saja, kenapa mendapatkan error ini seperti usernya tidak boleh ke server ini, "sepertinya sih"(?)
	c.JSON(http.StatusUnauthorized, web.WebResponse{
		Status:  http.StatusUnauthorized,
		Message: unauthorized.Error(),
	})
}

func ValidationErrorResponse(c *gin.Context, err interface{}) {
	validationErr, _ := err.(validator.ValidationErrors)
	c.JSON(http.StatusBadRequest, web.WebResponse{
		Status:  http.StatusBadRequest,
		Message: validationErr.Error(),
	})
}

func InternalServerError(c *gin.Context, err interface{}) {
	internalErr, ok := err.(string)
	if !ok {
		internalErr = "unknown error"
	}
	c.JSON(http.StatusInternalServerError, web.WebResponse{
		Status:  http.StatusInternalServerError,
		Message: internalErr,
	})
}
