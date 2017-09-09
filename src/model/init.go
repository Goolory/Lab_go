package model

import "github.com/jinzhu/gorm"

func InitDBModel(db *gorm.DB) {
	var err error

	err = initUser(db)
	if err != nil {
		print("Init db user failed", err)
		return
	}
}

func RebuildDbModel(db *gorm.DB) {
	dropUser(db)
}
