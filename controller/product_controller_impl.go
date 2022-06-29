package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincen320/product-service-mongodb/exception"
	"github.com/vincen320/product-service-mongodb/helper"
	"github.com/vincen320/product-service-mongodb/model/web"
	"github.com/vincen320/product-service-mongodb/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductControllerImpl struct {
	Service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &ProductControllerImpl{
		Service: service,
	}
}

func (pc *ProductControllerImpl) Create(c *gin.Context) {
	var createRequest web.ProductCreateRequest
	err := c.ShouldBind(&createRequest)
	if err != nil {
		errBad := exception.NewBadRequestErr("Can't bind request, " + err.Error())
		helper.ReturnError(c, errBad)
		return
	}

	//Waktu get disini, userId tipenya sudah ObjectID, sehingga kalo di konversi ke string akan error
	userId, exist := c.Get("user-id")
	if !exist {
		errNoLogin := exception.NewUnauthorizedErr("please, login first")
		helper.ReturnError(c, errNoLogin)
		return
	}

	userIdObject, ok := userId.(primitive.ObjectID)
	if !ok {
		errBadRequest := exception.NewBadRequestErr("id not valid")
		helper.ReturnError(c, errBadRequest)
		return
	}

	createRequest.IdUser = userIdObject

	response, err := pc.Service.Create(c, createRequest)
	if err != nil {
		helper.ReturnError(c, err)
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Successfull create product",
		Data:    response,
	})
}

func (pc *ProductControllerImpl) Update(c *gin.Context) {
	var updateRequest web.ProductUpdateRequest
	err := c.ShouldBind(&updateRequest)
	if err != nil {
		errBad := exception.NewBadRequestErr("Can't bind request, " + err.Error())
		helper.ReturnError(c, errBad)
		return
	}

	idProduct := c.Param("idProduct")
	idProductObject, err := primitive.ObjectIDFromHex(idProduct)
	if err != nil {
		errBadRequest := exception.NewBadRequestErr("id not valid")
		helper.ReturnError(c, errBadRequest)
		return
	}

	//Waktu get disini, userId tipenya sudah ObjectID, sehingga kalo di konversi ke string akan error
	userId, exist := c.Get("user-id")
	if !exist {
		errNoLogin := exception.NewUnauthorizedErr("please, login first")
		helper.ReturnError(c, errNoLogin)
		return
	}

	userIdObject, ok := userId.(primitive.ObjectID)
	if !ok {
		errBadRequest := exception.NewBadRequestErr("id not valid")
		helper.ReturnError(c, errBadRequest)
		return
	}

	updateRequest.IdUser = userIdObject
	updateRequest.Id = idProductObject

	response, err := pc.Service.Update(c, updateRequest)
	if err != nil {
		helper.ReturnError(c, err)
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Successfull update product",
		Data:    response,
	})
}

func (pc *ProductControllerImpl) Delete(c *gin.Context) {
	idProduct := c.Param("idProduct")
	idProductObject, err := primitive.ObjectIDFromHex(idProduct)
	if err != nil {
		errBadRequest := exception.NewBadRequestErr("id not valid")
		helper.ReturnError(c, errBadRequest)
		return
	}

	//Waktu get disini, userId tipenya sudah ObjectID, sehingga kalo di konversi ke string akan error
	userId, exist := c.Get("user-id")
	if !exist {
		errNoLogin := exception.NewUnauthorizedErr("please, login first")
		helper.ReturnError(c, errNoLogin)
		return
	}

	userIdObject, ok := userId.(primitive.ObjectID)
	if !ok {
		errBadRequest := exception.NewBadRequestErr("id not valid")
		helper.ReturnError(c, errBadRequest)
		return
	}

	err = pc.Service.Delete(c, idProductObject, userIdObject)
	if err != nil {
		helper.ReturnError(c, err)
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Message: "Successfull delete product",
	})
}

func (pc *ProductControllerImpl) FindById(c *gin.Context) {
	idProduct := c.Param("idProduct")
	idProductObject, err := primitive.ObjectIDFromHex(idProduct)
	if err != nil {
		errBadRequest := exception.NewBadRequestErr("id not valid")
		helper.ReturnError(c, errBadRequest)
		return
	}

	response, err := pc.Service.FindById(c, idProductObject)
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

func (pc *ProductControllerImpl) FindAll(c *gin.Context) {
	responses, err := pc.Service.FindAll(c)
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
