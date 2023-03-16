package entity

import "time"

type Category struct {
	Id        uint32    `json:"id" example:"1"`
	Level     uint16    `json:"level" example:"1"`
	ParentId  uint32    `json:"parent_id" example:"1"`
	Name      string    `json:"name" example:"1"`
	Rank      uint32    `json:"rank" example:"1"`
	IsDeleted bool      `json:"is_deleted" example:"1"`
	CreatedAt time.Time `json:"created_at" example:"1"`
	UpdatedAt time.Time `json:"updated_at" example:"1"`
}
