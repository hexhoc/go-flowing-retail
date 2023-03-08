package usecase

import (
	"github.com/hexhoc/product-service/internal/models"
	"github.com/hexhoc/product-service/internal/repository"
)

type ProductUseCase struct {
	productRepository *repository.ProductRepository
}

func NewProductUseCase(r *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{productRepository: r}
}

func (u *ProductUseCase) FindAll() []*models.Product {
	return u.productRepository.FindAll()
}

func (u *ProductUseCase) FindById(id uint32) *models.Product {
	return u.productRepository.FindById(id)
}

func (u *ProductUseCase) Save(product *models.Product) uint32 {
	return u.productRepository.Save(product)
}

func (u *ProductUseCase) SaveAll(products []*models.Product) {
	u.productRepository.SaveAll(products)
}

func (u *ProductUseCase) Update(id uint32, product *models.Product) {
	u.productRepository.Update(id, product)
}

func (u *ProductUseCase) Delete(id uint32) {
	u.productRepository.Delete(id)
}
