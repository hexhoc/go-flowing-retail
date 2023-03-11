package service

import (
	"context"

	"github.com/hexhoc/product-service/internal/entity"

	"github.com/hexhoc/product-service/internal/pb"
	"github.com/hexhoc/product-service/internal/repository"
)

type ProductInterface interface {
	FindAll(ctx context.Context, productRequest *pb.FindAllRequest) (*pb.FindAllResponse, error)
	//FindById(ctx context.Context, id uint32) (*entity.Product, error)
	//Save(ctx context.Context, product *entity.Product) error
	//SaveAll(ctx context.Context, products []*entity.Product) error
	//Update(ctx context.Context, id uint32, product *entity.Product) error
	//Delete(ctx context.Context, id uint32) error
}

// ProductService implements product_grpc interface ProductServiceServer
type ProductService struct {
	productRepository repository.ProductInterface
}

func NewProductService(r repository.ProductInterface) *ProductService {
	return &ProductService{productRepository: r}
}

func (s *ProductService) FindAll(ctx context.Context, ProductRequest *pb.ProductRequest) (*pb.ProductResponse, error) {

	list, err := s.productRepository.FindAll(ctx)
	if err != nil {
		return &pb.ProductResponse{Products: nil, Error: err.Error()}, err
	}

	var productsDto []*pb.ProductDto
	for i := 0; i < len(list); i++ {
		productsDto = append(productsDto, s.mapper(list[i]))
	}

	return &pb.ProductResponse{Products: productsDto, Error: ""}, nil
}

//func (s *ProductService) FindById(ctx context.Context, id uint32) (*entity.Product, error) {
//	return s.productRepository.FindById(ctx, id)
//}
//
//func (s *ProductService) Save(ctx context.Context, product *entity.Product) error {
//	return s.productRepository.Save(ctx, product)
//}
//
//func (s *ProductService) SaveAll(ctx context.Context, products []*entity.Product) error {
//	return s.productRepository.SaveAll(ctx, products)
//}
//
//func (s *ProductService) Update(ctx context.Context, id uint32, product *entity.Product) error {
//	return s.productRepository.Update(ctx, id, product)
//}
//
//func (s *ProductService) Delete(ctx context.Context, id uint32) error {
//	return s.productRepository.Delete(ctx, id)
//}

func (s *ProductService) mapper(e *entity.Product) *pb.ProductDto {
	dto := &pb.ProductDto{
		Id:            e.Id,
		Name:          e.Name,
		Intro:         e.Intro,
		Description:   e.Description,
		CategoryId:    e.CategoryId,
		Category:      e.Category.Name,
		OriginalPrice: e.OriginalPrice,
		SellingPrice:  e.SellingPrice,
		IsSale:        e.IsSale,
		IsDeleted:     e.IsDeleted,
		CreatedAt:     e.CreatedAt.Unix(),
		UpdatedAt:     e.UpdatedAt.Unix(),
	}

	return dto
}
