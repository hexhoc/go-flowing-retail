package usecase

import (
	"github.com/hexhoc/product-service/internal/interfaces"
	"github.com/hexhoc/product-service/internal/models"
)

type ProductUseCase struct {
	productRepository interfaces.ProductRepository
}

func NewProductUseCase(r interfaces.ProductRepository) *ProductUseCase {
	return &ProductUseCase{productRepository: r}
}

func (u *ProductUseCase) FindAll() []*models.Product {
	return u.productRepository.FindAll()
}
