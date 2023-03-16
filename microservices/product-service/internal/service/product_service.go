package service

import (
	"context"
	"time"

	"github.com/hexhoc/product-service/internal/entity"

	"github.com/hexhoc/product-service/internal/pb"
	"github.com/hexhoc/product-service/internal/repository"
)

type ProductInterface interface {
	FindAll(ctx context.Context, request *pb.FindAllRequest) (*pb.FindAllResponse, error)
	FindById(ctx context.Context, request *pb.FindByIdRequest) (*pb.FindByIdResponse, error)
	Save(ctx context.Context, request *pb.SaveRequest) (*pb.SaveResponse, error)
	SaveAll(ctx context.Context, request *pb.SaveAllRequest) (*pb.SaveAllResponse, error)
	Update(ctx context.Context, request *pb.UpdateRequest) (*pb.StatusResponse, error)
	Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.StatusResponse, error)
}

// ProductService implements product_grpc interface ProductServiceServer
type ProductService struct {
	productRepository repository.ProductInterface
}

func NewProductService(r repository.ProductInterface) *ProductService {
	return &ProductService{productRepository: r}
}

func (s *ProductService) FindAll(ctx context.Context, request *pb.FindAllRequest) (*pb.FindAllResponse, error) {

	// TODO: add limit and offset
	list, err := s.productRepository.FindAll(ctx, request.Limit, request.Offset)
	if err != nil {
		return &pb.FindAllResponse{Products: nil, Error: err.Error()}, err
	}

	var productsDto []*pb.ProductDto
	for i := 0; i < len(list); i++ {
		productsDto = append(productsDto, s.mapperToDto(list[i]))
	}

	return &pb.FindAllResponse{Products: productsDto, Error: ""}, nil
}

func (s *ProductService) FindById(ctx context.Context, request *pb.FindByIdRequest) (*pb.FindByIdResponse, error) {
	e, err := s.productRepository.FindById(ctx, request.GetId())
	if err != nil {
		return &pb.FindByIdResponse{Product: nil, Error: err.Error()}, err
	}

	return &pb.FindByIdResponse{Product: s.mapperToDto(e), Error: ""}, nil
}

func (s *ProductService) Save(ctx context.Context, request *pb.SaveRequest) (*pb.SaveResponse, error) {
	//TODO: return id
	err := s.productRepository.Save(ctx, s.mapperToEntity(request.Product))
	if err != nil {
		return &pb.SaveResponse{Status: "Not okay", Error: err.Error()}, err
	} else {
		return &pb.SaveResponse{Status: "OK", Error: ""}, nil
	}
}

func (s *ProductService) SaveAll(ctx context.Context, request *pb.SaveAllRequest) (*pb.SaveAllResponse, error) {
	//TODO: return ids
	var products []*entity.Product
	productsDto := request.Products
	for i := 0; i < len(productsDto); i++ {
		products = append(products, s.mapperToEntity(productsDto[i]))
	}
	err := s.productRepository.SaveAll(ctx, products)
	if err != nil {
		return &pb.SaveAllResponse{Status: "Not okay", Error: err.Error()}, err
	} else {
		return &pb.SaveAllResponse{Status: "OK", Error: ""}, nil
	}
}

func (s *ProductService) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.StatusResponse, error) {
	err := s.productRepository.Update(ctx, request.Id, s.mapperToEntity(request.Product))
	if err != nil {
		return &pb.StatusResponse{
			Status: "NOT OK",
			Error:  err.Error(),
		}, err
	}
	return &pb.StatusResponse{
		Status: "OK",
		Error:  "",
	}, nil
}

func (s *ProductService) Delete(ctx context.Context, request *pb.DeleteRequest) (*pb.StatusResponse, error) {
	err := s.productRepository.Delete(ctx, request.Id)
	if err != nil {
		return &pb.StatusResponse{
			Status: "NOT OK",
			Error:  err.Error(),
		}, err
	}
	return &pb.StatusResponse{
		Status: "OK",
		Error:  "",
	}, nil
}

func (s *ProductService) mapperToDto(e *entity.Product) *pb.ProductDto {
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

func (s *ProductService) mapperToEntity(d *pb.ProductDto) *entity.Product {
	e := &entity.Product{
		Id:            d.Id,
		Name:          d.Name,
		Intro:         d.Intro,
		Description:   d.Description,
		CategoryId:    d.CategoryId,
		OriginalPrice: d.OriginalPrice,
		SellingPrice:  d.SellingPrice,
		IsSale:        d.IsSale,
		IsDeleted:     d.IsDeleted,
		CreatedAt:     time.Unix(d.CreatedAt, 0),
		UpdatedAt:     time.Unix(d.UpdatedAt, 0),
	}

	return e
}
