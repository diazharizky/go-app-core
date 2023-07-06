package rdb

import (
	"gorm.io/gorm"
)

type IDatabase interface {
	GetDB() (*gorm.DB, error)
}

func GetDB(dbType DBType) (*gorm.DB, error) {
	return getDBType(dbType).GetDB()
}

func getDBType(dbType DBType) (db IDatabase) {
	switch dbType {
	case DBTypePostgres:
		db = NewPostgresClient()
	}

	return
}
