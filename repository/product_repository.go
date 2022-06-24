package repository

import (
	"context"
	"database/sql"

	"github.com/vincen320/product-service/model/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, tx *sql.Tx, productId, userId int) error
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
