package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
	"github.com/vincen320/product-service/helper"
	"github.com/vincen320/product-service/model/domain"
	"github.com/vincen320/product-service/model/web"
	"github.com/vincen320/product-service/repository"
)

type ProductServiceImpl struct {
	Repository repository.ProductRepository
	DB         *sql.DB
	Validator  *validator.Validate
}

func NewProductService(repo repository.ProductRepository, db *sql.DB, validator *validator.Validate) ProductService {
	return &ProductServiceImpl{
		Repository: repo,
		DB:         db,
		Validator:  validator,
	}
}

func (ps *ProductServiceImpl) Create(ctx context.Context, productRequest web.ProductCreateRequestWeb) web.ProductResponse {
	err := ps.Validator.Struct(productRequest)
	if err != nil {
		panic(err) //401 Bad Request  || Validation Error
	}
	tx, err := ps.DB.Begin()
	if err != nil {
		panic("DB Error : " + err.Error()) //500 Internal Server Error
	}
	defer helper.CommitOrRollBack(tx)

	productResponse := ps.Repository.Save(ctx, tx, domain.Product{
		IdUser:       productRequest.IdUser,
		NamaProduk:   productRequest.NamaProduk,
		Harga:        productRequest.Harga,
		Kategori:     productRequest.Kategori,
		Deskripsi:    productRequest.Deskripsi,
		Stok:         productRequest.Stok,
		LastModified: time.Now().UTC().UnixMilli(),
	})

	return web.ProductResponse{
		Id:           productResponse.Id,
		NamaProduk:   productResponse.NamaProduk,
		Harga:        productResponse.Harga,
		Kategori:     productResponse.Kategori,
		Deskripsi:    productResponse.Deskripsi,
		Stok:         productResponse.Stok,
		LastModified: productResponse.LastModified,
	}
}

func (ps *ProductServiceImpl) Update(ctx context.Context, productUpdate web.ProductUpdateWebRequest) web.ProductResponse {
	err := ps.Validator.Struct(productUpdate)
	if err != nil {
		panic(err) // 401 Bad Request  || Validation Error
	}
	tx, err := ps.DB.Begin()
	if err != nil {
		panic("DB Error : " + err.Error()) //500 Internal Server Error
	}
	defer helper.CommitOrRollBack(tx)

	_, err = ps.Repository.FindById(ctx, tx, productUpdate.Id)
	if err != nil {
		panic(err) // 404 Not Found
	}

	productResponse, err := ps.Repository.Update(ctx, tx, domain.Product{
		Id:           productUpdate.Id,
		IdUser:       productUpdate.IdUser,
		NamaProduk:   productUpdate.NamaProduk,
		Harga:        productUpdate.Harga,
		Kategori:     productUpdate.Kategori,
		Deskripsi:    productUpdate.Deskripsi,
		Stok:         productUpdate.Stok,
		LastModified: time.Now().UTC().UnixMilli(),
	})
	if err != nil {
		panic(err) //Tergantung dari Repository Errornya apa
	}

	return web.ProductResponse{
		Id:           productResponse.Id,
		NamaProduk:   productResponse.NamaProduk,
		Harga:        productResponse.Harga,
		Kategori:     productResponse.Kategori,
		Deskripsi:    productResponse.Deskripsi,
		Stok:         productResponse.Stok,
		LastModified: productResponse.LastModified,
	}
}

func (ps *ProductServiceImpl) Delete(ctx context.Context, productId int, userId int) {
	tx, err := ps.DB.Begin()
	if err != nil {
		panic("DB Error : " + err.Error()) //500 Internal Server Error || Validation Error
	}
	defer helper.CommitOrRollBack(tx)

	_, err = ps.Repository.FindById(ctx, tx, productId)
	if err != nil {
		panic(err) // 404 Not Found
	}

	err = ps.Repository.Delete(ctx, tx, productId, userId)
	if err != nil {
		panic(err) //Tergantung dari Repository Errornya apa
	}
}

func (ps *ProductServiceImpl) FindById(ctx context.Context, productId int) web.ProductResponse {
	tx, err := ps.DB.Begin()
	if err != nil {
		panic("DB Error : " + err.Error()) //500 Internal Server Error
	}
	defer helper.CommitOrRollBack(tx)
	product, err := ps.Repository.FindById(ctx, tx, productId)

	if err != nil {
		panic(err) // 404 Not Found
	}
	return web.ProductResponse{
		Id:           product.Id,
		NamaProduk:   product.NamaProduk,
		Harga:        product.Harga,
		Kategori:     product.Kategori,
		Deskripsi:    product.Deskripsi,
		Stok:         product.Stok,
		LastModified: product.LastModified,
	}
}

func (ps *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := ps.DB.Begin()
	if err != nil {
		panic("DB Error : " + err.Error()) //500 Internal Server Error
	}
	defer helper.CommitOrRollBack(tx)
	products := ps.Repository.FindAll(ctx, tx)

	var productsResponse []web.ProductResponse

	for _, product := range products {
		productsResponse = append(productsResponse, web.ProductResponse{
			Id:           product.Id,
			NamaProduk:   product.NamaProduk,
			Harga:        product.Harga,
			Kategori:     product.Kategori,
			Deskripsi:    product.Deskripsi,
			Stok:         product.Stok,
			LastModified: product.LastModified,
		})
	}
	return productsResponse
}

func (ps *ProductServiceImpl) UpdatePatch(ctx context.Context, productPatch web.ProductUpdatePatchWebRequest) web.ProductResponse {
	err := ps.Validator.Struct(productPatch)
	if err != nil {
		panic(err) // 401 Bad Request || Validation Error
	}
	tx, err := ps.DB.Begin()
	if err != nil {
		panic("DB Error : " + err.Error()) //500 Internal Server Error
	}
	defer helper.CommitOrRollBack(tx)

	productData, err := ps.Repository.FindById(ctx, tx, productPatch.Id)
	if err != nil {
		panic(err) // 404 Not Found
	}

	//Copy dari productPatch ke productData, jadi Data yang telah diupdate ada di variable productData
	copier.CopyWithOption(&productData, &productPatch, copier.Option{
		IgnoreEmpty: true,
	})

	/**
		productData.Id = productPatch.Id
		productData.IdUser = productPatch.IdUser
		productData.LastModified = time.Now().UTC().UnixMilli()
		productResponse, err := ps.Repository.Update(ctx, tx, productData)
	CARA INI BISA ATAU **/

	productResponse, err := ps.Repository.Update(ctx, tx, domain.Product{
		Id:           productPatch.Id,     //yang ini tetap dari request user
		IdUser:       productPatch.IdUser, //yang ini tetap dari request User
		NamaProduk:   productData.NamaProduk,
		Harga:        productData.Harga,
		Kategori:     productData.Kategori,
		Deskripsi:    productData.Deskripsi,
		Stok:         productData.Stok,
		LastModified: time.Now().UTC().UnixMilli(),
	})

	if err != nil {
		panic(err) //Tergantung dari Repository Errornya apa
	}

	return web.ProductResponse{
		Id:           productResponse.Id,
		NamaProduk:   productResponse.NamaProduk,
		Harga:        productResponse.Harga,
		Kategori:     productResponse.Kategori,
		Deskripsi:    productResponse.Deskripsi,
		Stok:         productResponse.Stok,
		LastModified: productResponse.LastModified,
	}
}
