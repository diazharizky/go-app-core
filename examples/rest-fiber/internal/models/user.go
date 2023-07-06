package models

import "github.com/diazharizky/go-app-core/pkg/rdb"

type User struct {
	ID       int64  `json:"id" gorm:"column:id"`
	Email    string `json:"email" gorm:"column:email;index:unique"`
	FullName string `json:"fullName" gorm:"column:full_name"`

	rdb.Date
}
