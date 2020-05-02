package models

import (
	"github.com/jinzhu/gorm"
)

// User model represents the schema for users table
type User struct {
	gorm.Model        // inject fields `ID`, `CreatedAt`, `UpdatedAt`, & `DeletedAt`
	GivenName  string `gorm:"size:255; not null" json:"given_name"`
	FamilyName string `gorm:"size:255; not null" json:"family_name"`
	Email      string `gorm:"type:varchar(100); unique_index" json:"email"`
}
