package controller

import "github.com/gin-gonic/gin"

type ProductController interface {
	CreateProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
	FindProductById(c *gin.Context)
	FindAllProducts(c *gin.Context)
	PatchProduct(c *gin.Context)
}
