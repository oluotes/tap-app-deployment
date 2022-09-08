package db

import "gorm.io/gorm"

type DBClient struct {
	Db *gorm.DB
}
