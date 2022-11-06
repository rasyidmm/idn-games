package model

import (
	"time"
)

// BaseModels struct to generate model id
type BaseModels struct {
	Id       int64  `gorm:"primaryKey;unique;autoIncrement"`
	CreateBy string `gorm:"column:create_by"`
	UpdateBy string `gorm:"column:update_by"`
}

// BaseCUModels struct to generate CreatedAt, UpdatedAt
type BaseCUModels struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
