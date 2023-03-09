package entity

import "time"

type Category struct {
	Id        uint32    `json:"id" example:"1"`
	Level     uint16    `json:"id" example:"1"`
	ParentId  uint32    `json:"id" example:"1"`
	Name      string    `json:"id" example:"1"`
	Rank      uint32    `json:"id" example:"1"`
	IsDeleted bool      `json:"id" example:"1"`
	CreatedAt time.Time `json:"id" example:"1"`
	UpdatedAt time.Time `json:"id" example:"1"`
}
