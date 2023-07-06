package rdb

import "time"

type Date struct {
	DateHardDelete

	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"column:deleted_at"`
}

type DateHardDelete struct {
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updated_at"`
}
