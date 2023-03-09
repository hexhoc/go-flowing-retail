package service

import (
	"context"
	"github.com/hexhoc/product-service/internal/entity"
	"github.com/hexhoc/product-service/internal/repository"
)

type ProductImageInterface interface {
	GetAllByProductId(ctx context.Context, id uint32) ([]*entity.ProductImage, error)
	UploadImageToProduct(ctx context.Context, product *entity.ProductImage) error
	DeleteByNameAndProductId(ctx context.Context, imageName string, productId uint32) error
}

type ProductImageService struct {
	productImageRepository repository.ProductImageInterface
}

func NewProductImageService(r repository.ProductImageInterface) *ProductImageService {
	return &ProductImageService{productImageRepository: r}
}

func (s *ProductImageService) GetAllByProductId(ctx context.Context, id uint32) ([]*entity.ProductImage, error) {
	return s.productImageRepository.FindAllByProductId(ctx, id)
}

func (s *ProductImageService) UploadImageToProduct(ctx context.Context, entity *entity.ProductImage) error {
	return s.productImageRepository.Save(ctx, entity)
}

func (s *ProductImageService) DeleteByNameAndProductId(ctx context.Context, imageName string, productId uint32) error {
	return s.productImageRepository.DeleteByNameAndProductId(ctx, imageName, productId)
}
