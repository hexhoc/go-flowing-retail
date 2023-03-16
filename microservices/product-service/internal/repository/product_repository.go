package repository

import (
	"context"
	"fmt"

	"github.com/hexhoc/product-service/internal/entity"
	"github.com/hexhoc/product-service/pkg/datasource/postgres"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

type ProductInterface interface {
	FindAll(ctx context.Context, limit uint32, offset uint32) ([]*entity.Product, error)
	FindById(ctx context.Context, id uint32) (*entity.Product, error)
	Save(ctx context.Context, product *entity.Product) error
	SaveAll(ctx context.Context, products []*entity.Product) error
	Update(ctx context.Context, id uint32, product *entity.Product) error
	Delete(ctx context.Context, id uint32) error
}

type ProductRepository struct {
	db *postgres.Postgres
}

func NewProductRepository(db *postgres.Postgres) *ProductRepository {
	return &ProductRepository{db: db}
}

// TODO: унифицировать здесь все методы как в order-service
// TODO: Добавить транзакции БД
func (r *ProductRepository) FindAll(ctx context.Context, limit uint32, offset uint32) ([]*entity.Product, error) {
	query := "SELECT * FROM products LIMIT $1 OFFSET $2"

	rows, err := r.db.Pool.Query(ctx, query, limit, offset)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var list []*entity.Product
	for rows.Next() {
		var item *entity.Product
		item = r.rowMapper(rows)
		list = append(list, item)

	}

	return list, nil
}

func (r *ProductRepository) FindById(ctx context.Context, id uint32) (*entity.Product, error) {
	query := "SELECT * FROM products p WHERE p.id = $1"

	rows, err := r.db.Pool.Query(ctx, query, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer rows.Close()

	var item *entity.Product
	for rows.Next() {
		item = r.rowMapper(rows)
	}

	return item, nil
}

func (r *ProductRepository) Save(ctx context.Context, item *entity.Product) error {
	query := `
	INSERT INTO products(name, intro, description, category_id, original_price, selling_price, is_sale, is_deleted, created_at, updated_at) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	`

	ct, err := r.db.Pool.Exec(
		ctx, query,
		item.Name,
		item.Intro,
		item.Description,
		item.CategoryId,
		item.OriginalPrice,
		item.SellingPrice,
		item.IsSale,
		item.IsDeleted,
		item.CreatedAt,
		item.UpdatedAt,
	)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(fmt.Sprintf("Product update row affected %d", ct.RowsAffected()))

	return nil
}

func (r *ProductRepository) SaveAll(ctx context.Context, list []*entity.Product) error {
	//TODO: Implement
	//r.ProductRepository.SaveAll(products)
	return nil
}

// TODO: TEST
func (r *ProductRepository) Update(ctx context.Context, id uint32, item *entity.Product) error {
	query := `
	UPDATE products p SET 
		name = $1,
		intro = $2,
		description = $3,
		category_id = $4,
		original_price = $5,
		selling_price = $6,
		is_sale = $7,
		is_deleted = $8,
		updated_at = $9
	WHERE p.id = $10
	`

	ct, err := r.db.Pool.Exec(
		ctx, query,
		item.Name,
		item.Intro,
		item.Description,
		item.CategoryId,
		item.OriginalPrice,
		item.SellingPrice,
		item.IsSale,
		item.IsDeleted,
		item.UpdatedAt,
		id,
	)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(fmt.Sprintf("Product update row affected %d", ct.RowsAffected()))

	return nil
}

func (r *ProductRepository) Delete(ctx context.Context, id uint32) error {

	query := `DELETE FROM products where products.id = $1`

	ct, err := r.db.Pool.Exec(
		ctx, query,
		id,
	)

	if err != nil {
		log.Error(err)
		return err
	}

	log.Info(fmt.Sprintf("Product delete row affected %d", ct.RowsAffected()))

	return nil
}

func (r *ProductRepository) rowMapper(rows pgx.Rows) *entity.Product {
	var item entity.Product
	err := rows.Scan(
		&item.Id,
		&item.Name,
		&item.Intro,
		&item.Description,
		&item.CategoryId,
		&item.OriginalPrice,
		&item.SellingPrice,
		&item.IsSale,
		&item.IsDeleted,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		log.Error(fmt.Errorf("error while iterating dataset %w", err))
	}

	return &item
}
