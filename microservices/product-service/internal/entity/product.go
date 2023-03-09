package entity

import "time"

type Product struct {
	Id            uint32    `json:"id"                   example:"1"`
	Name          string    `json:"name"                 example:"macbook"`
	Intro         string    `json:"intro"                example:"intro"`
	Description   string    `json:"description"          example:"description"`
	CategoryId    uint32    `json:"category_id"          example:"1"`
	Category      Category  `json:"category"             example:"notebook"`
	OriginalPrice float64   `json:"original_price"       example:"10000.00"`
	SellingPrice  float64   `json:"selling_price"        example:"10000.00"`
	IsSale        bool      `json:"is_sale"              example:"true"`
	IsDeleted     bool      `json:"is_deleted"           example:"false"`
	CreatedAt     time.Time `json:"created_at"           example:"2022-10-20 14:00:00"`
	UpdatedAt     time.Time `json:"updated_at"           example:"2022-10-20 14:00:00"`
}
