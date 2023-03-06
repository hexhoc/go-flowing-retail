package models

import "time"

type Category struct {
	Id        uint32
	Level     uint16
	ParentId  uint32
	Name      string
	Rank      uint32
	IsDeleted bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
