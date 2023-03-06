package models

type ProductImage struct {
	Id         uint32
	ProductId  uint32
	Name       string
	ImageBytes []byte
}
