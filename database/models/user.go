package models

import (
	"github.com/jinzhu/gorm"
	"github.com/velopert/gin-rest-api-sample/lib/common"
)

// User data model
type User struct {
	gorm.Model
	Username     string
	DisplayName  string
	PasswordHash string
}

// Serialize serializes user data
func (u User) Serialize() common.JSON {
	return common.JSON{
		"id":           u.ID,
		"username":     u.Username,
		"display_name": u.DisplayName,
	}
}
