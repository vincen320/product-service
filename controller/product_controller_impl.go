package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincen320/product-service/exception"
	"github.com/vincen320/product-service/model/web"
	"github.com/vincen320/product-service/service"
)

type ProductControllerImpl struct {
	Service service.ProductService
}

func NewProductController(service service.ProductService) ProductController {
	return &ProductControllerImpl{
		Service: service,
	}
}

func (pc *ProductControllerImpl) CreateProduct(c *gin.Context) {
	var request web.ProductCreateRequestWeb
	err := c.ShouldBind(&request)
	if err != nil {
		panic(exception.NewBadRequestError("Error get Request " + err.Error())) //bad request
	}

	idUser, exists := c.Get("id_user")
	if !exists {
		panic(exception.NewUnauthorizedError("Unauthorized")) //401 Unauthorized
	}
	idUserInt, ok := idUser.(int)
	if !ok {
		panic(exception.NewUnauthorizedError("Unauthorized User")) // 401 Unauthorized
	}

	request.IdUser = idUserInt
	response := pc.Service.Create(c, request)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Success Create Product",
		Data:    response,
	})
}

func (pc *ProductControllerImpl) UpdateProduct(c *gin.Context) {
	var request web.ProductUpdateWebRequest
	err := c.ShouldBind(&request)
	if err != nil {
		panic(exception.NewBadRequestError("Error get Request " + err.Error())) //400 bad request
	}

	idUser, exists := c.Get("id_user")
	if !exists {
		panic(exception.NewUnauthorizedError("Unauthorized")) //401 Unauthorized
	}
	idUserInt, ok := idUser.(int)
	if !ok {
		panic(exception.NewUnauthorizedError("Unauthorized User")) // 401 Unauthorized
	}

	idProduct := c.Param("productId")
	idProductInt, err := strconv.Atoi(idProduct)
	if err != nil {
		panic(exception.NewBadRequestError("id product must contain number")) // 400 Bad Request
	}

	request.IdUser = idUserInt
	request.Id = idProductInt
	response := pc.Service.Update(c, request)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Success Update Product",
		Data:    response,
	})
}

func (pc *ProductControllerImpl) DeleteProduct(c *gin.Context) {
	idUser, exists := c.Get("id_user")
	if !exists {
		panic(exception.NewUnauthorizedError("Unauthorized")) //401 Unauthorized
	}
	idUserInt, ok := idUser.(int)
	if !ok {
		panic(exception.NewUnauthorizedError("Unauthorized User")) // 401 Unauthorized
	}

	idProduct := c.Param("productId")
	idProductInt, err := strconv.Atoi(idProduct)
	if err != nil {
		panic(exception.NewBadRequestError("id product must contain number")) // 400 Bad Request
	}

	pc.Service.Delete(c, idProductInt, idUserInt)

	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Success Delete Product",
		Data:    nil,
	})
}

func (pc *ProductControllerImpl) FindProductById(c *gin.Context) {
	idProduct := c.Param("productId")
	idProductInt, err := strconv.Atoi(idProduct)
	if err != nil {
		panic(exception.NewBadRequestError("id product must contain number")) // 400 Bad Request
	}

	response := pc.Service.FindById(c, idProductInt)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Success Find Product",
		Data:    response,
	})
}

func (pc *ProductControllerImpl) FindAllProducts(c *gin.Context) {
	response := pc.Service.FindAll(c)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Success Find All Product",
		Data:    response,
	})
}

func (pc *ProductControllerImpl) PatchProduct(c *gin.Context) {
	var request web.ProductUpdatePatchWebRequest
	err := c.ShouldBind(&request)
	if err != nil {
		panic(exception.NewBadRequestError("Error get Request " + err.Error())) //bad request
	}

	idUser, exists := c.Get("id_user")
	if !exists {
		panic(exception.NewUnauthorizedError("Unauthorized")) //401 Unauthorized
	}
	idUserInt, ok := idUser.(int)
	if !ok {
		panic(exception.NewUnauthorizedError("Unauthorized User")) // 401 Unauthorized
	}

	idProduct := c.Param("productId")
	idProductInt, err := strconv.Atoi(idProduct)
	if err != nil {
		panic(exception.NewBadRequestError("Error get Request " + err.Error())) //bad request
	}

	request.IdUser = idUserInt
	request.Id = idProductInt
	response := pc.Service.UpdatePatch(c, request)
	c.JSON(200, web.WebResponse{
		Status:  200,
		Message: "Success Patch Product",
		Data:    response,
	})
}
