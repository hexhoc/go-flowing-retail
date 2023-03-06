package models

import "time"

type Product struct {
	Id            uint32
	Name          string
	Intro         string
	Description   string
	CategoryId    uint32
	Category      Category
	OriginalPrice float64
	SellingPrice  float64
	IsSale        bool
	IsDeleted     bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
