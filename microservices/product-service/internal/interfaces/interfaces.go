package interfaces

import "github.com/hexhoc/product-service/internal/models"

type ProductRepository interface {
	FindAll() []*models.Product
}

type ProductUseCase interface {
	FindAll() []*models.Product
}
