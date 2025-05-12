package repository

import (
	"locket-clone/backend/pkg/model"
	"locket-clone/backend/pkg/service/adding"
	"time"

	"gorm.io/gorm"
)

type LocketRepo struct {
	db *gorm.DB
}

type Locket struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"index:,sort:desc,type:btree"`
	DeleteAt  time.Time `gorm:"index"`
	Type      model.LocketType
	ImageUrl  string
	Caption   string
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UserRepo struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	adding.UserRecord
	Friends []User `gorm:"foreignkey:UserID"`
}
