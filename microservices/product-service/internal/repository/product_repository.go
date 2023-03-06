package repository

import (
	"context"
	"log"

	"github.com/hexhoc/product-service/internal/models"
	"github.com/hexhoc/product-service/pkg/datasource/postgres"
)

type ProductRepository struct {
	db *postgres.Postgres
}

func NewProductRepository(db *postgres.Postgres) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindAll() []*models.Product {
	query := "SELECT * FROM products"

	rows, err := r.db.Pool.Query(context.Background(), query)

	if err != nil {
		log.Fatalf("(u *UserInfoRepository) findById(userId int64): %s", err)
	}

	var product models.Product
	var products []*models.Product
	for rows.Next() {
		err = rows.Scan(&product)
		if err != nil {
			log.Println("error while iterating dataset")
		}
		products = append(products, &product)
	}
	defer rows.Close()

	return products
}
