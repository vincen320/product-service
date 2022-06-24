package service

import (
	"context"

	"github.com/vincen320/product-service/model/web"
)

type ProductService interface {
	Create(ctx context.Context, productRequest web.ProductCreateRequestWeb) web.ProductResponse
	Update(ctx context.Context, productUpdate web.ProductUpdateWebRequest) web.ProductResponse
	Delete(ctx context.Context, productId, userId int)
	FindById(ctx context.Context, productId int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
	UpdatePatch(ctx context.Context, productPatch web.ProductUpdatePatchWebRequest) web.ProductResponse
}
