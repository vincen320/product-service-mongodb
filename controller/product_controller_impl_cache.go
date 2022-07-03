package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincen320/product-service-mongodb/exception"
	"github.com/vincen320/product-service-mongodb/helper"
	"github.com/vincen320/product-service-mongodb/model/web"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Logicnya sama seperti controller.FindById
func (pc *ProductControllerImpl) FindByIdCache(c *gin.Context) {
	idProduct := c.Param("idProduct")
	idProductObject, err := primitive.ObjectIDFromHex(idProduct)
	if err != nil {
		errBadRequest := exception.NewBadRequestErr("id not valid")
		helper.ReturnError(c, errBadRequest)
		return
	}

	response, err := pc.Service.FindByIdCache(c, idProductObject)
	if err != nil {
		helper.ReturnError(c, err)
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Successfull Find product by Id",
		Data:    response,
	})
}
func (pc *ProductControllerImpl) FindAllCache(c *gin.Context) {
	responses, err := pc.Service.FindAllCache(c)
	if err != nil {
		helper.ReturnError(c, err)
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Successfull Find All product",
		Data:    responses,
	})
}
