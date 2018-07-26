package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(User{})
	fmt.Println("Auto Migration has beed processed")
}
