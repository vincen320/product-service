package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vincen320/product-service/app"
	"github.com/vincen320/product-service/controller"
	"github.com/vincen320/product-service/middleware"
	"github.com/vincen320/product-service/repository"
	"github.com/vincen320/product-service/service"
)

func main() {
	db := app.NewConnection()
	validator := validator.New()
	productRepository := repository.NewProductRespository()
	productService := service.NewProductService(productRepository, db, validator)
	productController := controller.NewProductController(productService)

	router := gin.New()
	router.Use(middleware.ErrorHandling()) //diatas emg
	rgroup := router.Group("/", middleware.AuthenticateJWT())
	{
		rgroup.POST("/products", productController.CreateProduct)
		rgroup.PUT("/products/:productId", productController.UpdateProduct)
		rgroup.PATCH("/products/:productId", productController.PatchProduct)
		rgroup.DELETE("/products/:productId", productController.DeleteProduct)
	}

	router.GET("/products/:productId", productController.FindProductById)
	router.GET("/products", productController.FindAllProducts)

	server := &http.Server{
		Addr:           ":8082",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Product Service Start in 8082 port")
	err := server.ListenAndServe()
	if err != nil {
		panic("Cannot Start Server " + err.Error()) //500 Internal Server Error
	}

	//Tinggal buat middleware authentication
	//middleware authentication:
	// cek valid token
	//cek exp
	//kalau exp, redirect ke auth-server/refresh [BUAT HANDLER LAGI]
	//implementasi refreshtoken asli nanti saja
	//ambil id user https://stackoverflow.com/questions/69600198/how-to-pass-a-struct-from-middleware-to-handler-in-gin-without-marshalling

	//buat fungsi refresh di auth
}
