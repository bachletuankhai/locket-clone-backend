package repository

import (
	"locket-clone/backend/pkg/model"
	"locket-clone/backend/pkg/service/adding"

	"gorm.io/gorm"
)

type LocketRepo struct {
	db *gorm.DB
}

type Locket struct {
	gorm.Model
	Type     model.LocketType
	ImageUrl string
	Caption  string
	UserID   uint
	User     User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserRepo struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	adding.UserRecord
	Friends []User `gorm:"foreignkey:UserID"`
}
