package repository

import (
	"context"
	"fmt"

	"github.com/hexhoc/product-service/internal/entity"
	"github.com/hexhoc/product-service/pkg/datasource/postgres"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

type ProductImageInterface interface {
	FindAllByProductId(ctx context.Context, id uint32) ([]*entity.ProductImage, error)
	Save(ctx context.Context, product *entity.ProductImage) error
	DeleteByNameAndProductId(ctx context.Context, imageName string, productId uint32) error
}

type ProductImageRepository struct {
	db *postgres.Postgres
}

func NewProductImageRepository(db *postgres.Postgres) *ProductImageRepository {
	return &ProductImageRepository{db: db}
}

func (r *ProductImageRepository) FindAllByProductId(ctx context.Context, id uint32) ([]*entity.ProductImage, error) {
	query := "SELECT * FROM product_images p WHERE p.id = $1"

	rows, err := r.db.Pool.Query(ctx, query, id)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	var products []*entity.ProductImage
	for rows.Next() {
		var productImage entity.ProductImage
		r.rowMapper(rows, &productImage)
		products = append(products, &productImage)
	}
	defer rows.Close()

	return products, nil
}

func (r *ProductImageRepository) Save(ctx context.Context, productImage *entity.ProductImage) error {
	query := `
	INSERT INTO product_images(product_id, name, image_bytes) 
	VALUES ($1,$2,$3)
	`

	ct, err := r.db.Pool.Exec(
		ctx, query,
		productImage.ProductId,
		productImage.Name,
		productImage.ImageBytes,
	)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(fmt.Sprintf("Product image update row affected %d", ct.RowsAffected()))

	return nil
}

func (r *ProductImageRepository) DeleteByNameAndProductId(ctx context.Context, imageName string, productId uint32) error {

	query := `DELETE FROM product_images p where p.name = $1 AND p.product_id = $2`

	ct, err := r.db.Pool.Exec(
		ctx, query,
		imageName,
		productId,
	)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(fmt.Sprintf("Product image delete row affected %d", ct.RowsAffected()))

	return nil
}

func (r *ProductImageRepository) rowMapper(rows pgx.Rows, productImage *entity.ProductImage) {
	err := rows.Scan(
		&productImage.Id,
		&productImage.ProductId,
		&productImage.Name,
		&productImage.ImageBytes,
	)

	if err != nil {
		log.Error(fmt.Errorf("error while iterating dataset %w", err))
	}

}
