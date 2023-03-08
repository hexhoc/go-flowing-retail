package repository

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hexhoc/product-service/internal/models"
	"github.com/hexhoc/product-service/pkg/datasource/postgres"
	"github.com/jackc/pgx/v4"
)

var l = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

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
		l.Fatal(err)
	}

	var products []*models.Product
	for rows.Next() {
		var product models.Product
		mapModel(rows, &product)
		products = append(products, &product)

	}
	defer rows.Close()

	return products
}

func (r *ProductRepository) FindById(id uint32) *models.Product {
	query := "SELECT * FROM products p WHERE p.id = $1"

	rows, err := r.db.Pool.Query(context.Background(), query, id)

	if err != nil {
		l.Fatal(err)
	}

	var product models.Product
	for rows.Next() {
		mapModel(rows, &product)
	}

	defer rows.Close()

	return &product
}

func (r *ProductRepository) Save(product *models.Product) uint32 {
	query := `
	UPDATE "products" SET 
		"name" = $1
		"intro" = $2
		"description" = $3
		"category_id" = $4
		"original_price" = $5
		"selling_price" = $6
		"is_sale" = $7
		"is_deleted" = $8
		"created_at" = $9
		"updated_at" = $10
	WHERE p.id = $11`

	ct, err := r.db.Pool.Exec(
		context.Background(), query,
		product.Name,
		product.Intro,
		product.Description,
		product.CategoryId,
		product.OriginalPrice,
		product.SellingPrice,
		product.IsSale,
		product.IsDeleted,
		product.CreatedAt,
		product.UpdatedAt,
	)

	if err != nil {
		l.Fatal(err)
	}

	fmt.Println("Product update row affected ", ct.RowsAffected())

	return product.Id
}

func (r *ProductRepository) SaveAll(products []*models.Product) {
	// r.productRepository.SaveAll(products)
}

func (r *ProductRepository) Update(id uint32, product *models.Product) {
	// r.productRepository.Update(id, product)
}

func (r *ProductRepository) Delete(id uint32) {
	// r.productRepository.Delete(id)
}

func mapModel(r pgx.Rows, product *models.Product) {
	err := r.Scan(
		&product.Id,
		&product.Name,
		&product.Intro,
		&product.Description,
		&product.CategoryId,
		&product.OriginalPrice,
		&product.SellingPrice,
		&product.IsSale,
		&product.IsDeleted,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		l.Println("error while iterating dataset ", err)
	}

}
