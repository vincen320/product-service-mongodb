package controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
	FindByIdCache(c *gin.Context)
	FindAllCache(c *gin.Context)
}
