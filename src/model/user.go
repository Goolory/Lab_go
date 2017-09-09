package model

import "github.com/jinzhu/gorm"

type User struct {
	Id uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Username string `gorm:"size:64" json:"username"`
	Password string `gorm:"size:225" json:"password"`
}

func (User) TableName() string {
	return "user"
}

func initUser(db *gorm.DB) error {
	var err error

	if db.HasTable(&User{}) {
		err = db.AutoMigrate(&User{}).Error
	} else {
		err = db.CreateTable(&User{}).Error
	}
	return err
}

func dropUser(db *gorm.DB) {
	db.DropTableIfExists(&User{})
}