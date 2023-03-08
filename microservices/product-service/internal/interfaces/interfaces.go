package interfaces

import "github.com/hexhoc/product-service/internal/models"

type ProductRepository interface {
	FindAll() []*models.Product
	FindById(id uint32) *models.Product
	Save(product *models.Product) uint32
	SaveAll(products []*models.Product)
	Update(id uint32, product *models.Product)
	Delete(id uint32)
}

type ProductUseCase interface {
	GetAll() []*models.Product
	GetById(id uint32) *models.Product
	Save(product *models.Product) uint32
	SaveAll(products []*models.Product)
	Update(id uint32, product *models.Product)
	Delete(id uint32)
}
