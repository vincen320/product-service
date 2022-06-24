package repository

import (
	"context"
	"database/sql"

	"github.com/vincen320/product-service/exception"
	"github.com/vincen320/product-service/model/domain"
)

type ProductRepositoryImpl struct {
}

func NewProductRespository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (pr *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := `INSERT INTO product (id_user, nama_produk, harga, kategori, deskripsi, stok, last_modified)
		VALUES (?,?,?,?,?,?,?)`
	result, err := tx.ExecContext(ctx, SQL, product.IdUser, product.NamaProduk, product.Harga, product.Kategori,
		product.Deskripsi, product.Stok, product.LastModified)
	if err != nil {
		panic("query error : " + err.Error()) // 500 Internal Server Error
	}
	productId, err := result.LastInsertId()
	if err != nil {
		panic("cannot get last id : " + err.Error()) // 500 Internal Server Error
	}
	product.Id = int(productId)
	return product
}

func (pr *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) (domain.Product, error) {
	SQL := `UPDATE PRODUCT SET nama_produk=?, harga=?, kategori=?, deskripsi=?, stok=?, last_modified=? WHERE id=? AND id_user=?`
	result, err := tx.ExecContext(ctx, SQL, product.NamaProduk, product.Harga, product.Kategori,
		product.Deskripsi, product.Stok, product.LastModified, product.Id, product.IdUser)
	if err != nil {
		panic("query error : " + err.Error()) // 500 Internal Server Error
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic("driver not support rows affected : " + err.Error()) // 500 Internal Server Error
	}

	if rowsAffected == 0 {
		return product, exception.NewUnauthorizedError("can't update product that's not yours") //401 Unauthorized
	}
	if rowsAffected > 1 {
		return product, exception.NewUnauthorizedError("can't update more than 1 items in one time") // 401 Unauthorized
	}
	return product, nil
}

func (pr *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId, userId int) error {
	SQL := `DELETE FROM PRODUCT WHERE ID=? AND ID_USER=?`
	result, err := tx.ExecContext(ctx, SQL, productId, userId)
	if err != nil {
		panic("query error : " + err.Error()) // 500 Internal Server Error
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic("driver not support rows affected : " + err.Error()) // 500 Internal Server Error
	}
	if rowsAffected == 0 {
		return exception.NewUnauthorizedError("can't delete product that's not yours") //401 Unauthorized
	}
	if rowsAffected > 1 {
		return exception.NewUnauthorizedError("can't delete more than 1 items in one time") // 401 Unauthorized
	}
	return nil
}

func (pr *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := `SELECT id, id_user, nama_produk, harga, kategori, deskripsi, stok, last_modified FROM PRODUCT WHERE id=?`

	result, err := tx.QueryContext(ctx, SQL, productId)
	if err != nil {
		panic("query error : " + err.Error()) // 500 Internal Server Error
	}
	defer result.Close()
	var product domain.Product

	if result.Next() {
		result.Scan(&product.Id, &product.IdUser, &product.NamaProduk, &product.Harga, &product.Kategori, &product.Deskripsi,
			&product.Stok, &product.LastModified)
		return product, nil
	} else {
		return product, exception.NewNotFoundError("product not found") // 404 Not Found
	}
}

func (pr *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := `SELECT id, id_user, nama_produk, harga, kategori, deskripsi, stok, last_modified FROM PRODUCT`
	result, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic("query error : " + err.Error()) // 500 Internal Server Error
	}
	defer result.Close()
	var products []domain.Product

	for result.Next() {
		var product domain.Product
		result.Scan(&product.Id, &product.IdUser, &product.NamaProduk, &product.Harga, &product.Kategori, &product.Deskripsi,
			&product.Stok, &product.LastModified)
		products = append(products, product)
	}
	return products
}
