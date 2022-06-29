package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincen320/product-service-mongodb/model/web"
)

func PanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			e := recover()
			if e != nil {
				c.JSON(http.StatusInternalServerError, web.WebResponse{
					Code:    http.StatusInternalServerError,
					Message: "unknown error, ",
					Data:    e,
				})
			}
		}()
		c.Next()
	}
}
