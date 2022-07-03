package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vincen320/product-service-mongodb/app"
	"github.com/vincen320/product-service-mongodb/controller"
	"github.com/vincen320/product-service-mongodb/middleware"
	"github.com/vincen320/product-service-mongodb/repository"
	"github.com/vincen320/product-service-mongodb/service"
)

func main() {

	db, err := app.ConnectMongo()
	if err != nil {
		panic(err)
	}

	rdb := app.ConnectRedis()
	validator := validator.New()
	productRepository := repository.NewProductRespository()
	productService := service.NewProductService(productRepository, validator, db, rdb)
	productController := controller.NewProductController(productService)

	router := gin.New()
	router.Use(middleware.PanicHandler())

	rgroup := router.Group("/", middleware.AuthenticateJWT())
	{
		rgroup.POST("/products", productController.Create)
		rgroup.PUT("/products/:idProduct", productController.Update)
		rgroup.PATCH("/products/:idProduct", productController.Update)
		rgroup.DELETE("/products/:idProduct", productController.Delete)
	}
	router.GET("products/:idProduct", productController.FindById)
	router.GET("products", productController.FindAll)
	router.GET("products/cache/:idProduct", productController.FindByIdCache)
	router.GET("products/cache", productController.FindAllCache)

	server := http.Server{
		Addr:           ":8082",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Product Service Start in 8082 port")
	err = server.ListenAndServe()
	if err != nil {
		panic("Cannot Start Server " + err.Error()) //500 Internal Server Error
	}
}
