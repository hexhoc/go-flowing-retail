package service

import (
	"context"
	"github.com/hexhoc/product-service/internal/entity"
	"github.com/hexhoc/product-service/internal/repository"
)

type ProductInterface interface {
	FindAll(ctx context.Context) ([]*entity.Product, error)
	FindById(ctx context.Context, id uint32) (*entity.Product, error)
	Save(ctx context.Context, product *entity.Product) error
	SaveAll(ctx context.Context, products []*entity.Product) error
	Update(ctx context.Context, id uint32, product *entity.Product) error
	Delete(ctx context.Context, id uint32) error
}

type ProductService struct {
	productRepository repository.ProductInterface
}

func NewProductService(r repository.ProductInterface) *ProductService {
	return &ProductService{productRepository: r}
}

func (s *ProductService) FindAll(ctx context.Context) ([]*entity.Product, error) {
	return s.productRepository.FindAll(ctx)
}

func (s *ProductService) FindById(ctx context.Context, id uint32) (*entity.Product, error) {
	return s.productRepository.FindById(ctx, id)
}

func (s *ProductService) Save(ctx context.Context, product *entity.Product) error {
	return s.productRepository.Save(ctx, product)
}

func (s *ProductService) SaveAll(ctx context.Context, products []*entity.Product) error {
	return s.productRepository.SaveAll(ctx, products)
}

func (s *ProductService) Update(ctx context.Context, id uint32, product *entity.Product) error {
	return s.productRepository.Update(ctx, id, product)
}

func (s *ProductService) Delete(ctx context.Context, id uint32) error {
	return s.productRepository.Delete(ctx, id)
}
