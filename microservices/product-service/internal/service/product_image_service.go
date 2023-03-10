package service

import (
	"bytes"
	"compress/gzip"
	"context"
	"github.com/hexhoc/product-service/internal/entity"
	"github.com/hexhoc/product-service/internal/repository"
)

type ProductImageInterface interface {
	GetAllByProductId(ctx context.Context, id uint32) ([]*entity.ProductImage, error)
	UploadImageToProduct(ctx context.Context, file []byte, productId uint32) error
	DeleteByNameAndProductId(ctx context.Context, imageName string, productId uint32) error
	compressBytes(data []byte) []byte
	decompressBytes(data []byte) []byte
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

func (s *ProductImageService) UploadImageToProduct(ctx context.Context, file []byte, productId uint32) error {
	// TODO: Add filename
	var productImage *entity.ProductImage = &entity.ProductImage{
		ProductId:  productId,
		Name:       "",
		ImageBytes: s.compressBytes(file),
	}
	return s.productImageRepository.Save(ctx, productImage)
}

func (s *ProductImageService) DeleteByNameAndProductId(ctx context.Context, imageName string, productId uint32) error {
	return s.productImageRepository.DeleteByNameAndProductId(ctx, imageName, productId)
}

func (s *ProductImageService) compressBytes(data []byte) []byte {
	// Initialize gzip
	buf := new(bytes.Buffer)
	gzWriter := gzip.NewWriter(buf)
	gzWriter.Write(data)
	gzWriter.Close()
	// Convert buffer to
	return buf.Bytes()
}

func (s *ProductImageService) decompressBytes(data []byte) []byte {
	// Initialize gzip
	buf := new(bytes.Buffer)
	gzWriter, _ := gzip.NewReader(buf)
	gzWriter.Read(data)
	gzWriter.Close()
	// Convert buffer to
	return buf.Bytes()
}
