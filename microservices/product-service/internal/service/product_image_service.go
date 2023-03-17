package service

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"

	"github.com/hexhoc/product-service/internal/entity"
	"github.com/hexhoc/product-service/internal/pb"
	"github.com/hexhoc/product-service/internal/repository"
	log "github.com/sirupsen/logrus"
)

type ProductImageInterface interface {
	GetAllByProductId(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error)
	UploadImageToProduct(ctx context.Context, request *pb.UploadImageRequest) (*pb.UploadImageResponse, error)
	DeleteByNameAndProductId(ctx context.Context, request *pb.DeleteImageRequest) (*pb.DeleteImageResponse, error)
}

type ProductImageService struct {
	productImageRepository repository.ProductImageInterface
}

func NewProductImageService(r repository.ProductImageInterface) *ProductImageService {
	return &ProductImageService{productImageRepository: r}
}

func (s *ProductImageService) GetAllByProductId(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	images, err := s.productImageRepository.FindAllByProductId(ctx, request.Id)
	if err != nil {
		return &pb.GetAllResponse{
			Products: []*pb.ProductImageDto{},
			Error:    err.Error(),
		}, err
	}

	var imagesDto []*pb.ProductImageDto
	for i := 0; i < len(images); i++ {
		imagesDto = append(imagesDto, s.mapperToDto(images[0]))
	}

	if len(images) == 0 {
		err := errors.New("nothing found")
		log.Error(err)
		return nil, err
	}

	return &pb.GetAllResponse{
		Products: imagesDto,
		Error:    "",
	}, nil
}

func (s *ProductImageService) UploadImageToProduct(ctx context.Context, request *pb.UploadImageRequest) (*pb.UploadImageResponse, error) {
	var productImage *entity.ProductImage = &entity.ProductImage{
		ProductId:  request.ProductId,
		Name:       request.Filename,
		ImageBytes: s.compressBytes(request.File),
	}

	err := s.productImageRepository.Save(ctx, productImage)
	if err != nil {
		return &pb.UploadImageResponse{
			Status: "NOT OK",
			Error:  err.Error(),
		}, err
	}

	return &pb.UploadImageResponse{
		Status: "OK",
		Error:  "",
	}, nil
}

func (s *ProductImageService) DeleteByNameAndProductId(ctx context.Context, request *pb.DeleteImageRequest) (*pb.DeleteImageResponse, error) {
	err := s.productImageRepository.DeleteByNameAndProductId(ctx, request.ImageName, request.ProductId)
	if err != nil {
		return &pb.DeleteImageResponse{
			Status: "NOT OK",
			Error:  err.Error(),
		}, err
	}

	return &pb.DeleteImageResponse{
		Status: "OK",
		Error:  "",
	}, nil
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

// TODO: use decompress
func (s *ProductImageService) decompressBytes(data []byte) []byte {
	// Initialize gzip
	buf := new(bytes.Buffer)
	gzWriter, _ := gzip.NewReader(buf)
	gzWriter.Read(data)
	gzWriter.Close()
	// Convert buffer to
	return buf.Bytes()
}

func (s *ProductImageService) mapperToDto(e *entity.ProductImage) *pb.ProductImageDto {
	return &pb.ProductImageDto{
		Id:        e.Id,
		Name:      e.Name,
		ProductId: e.ProductId,
	}
}
