package repository

import (
	"locket-clone/backend/pkg/model"
	"locket-clone/backend/pkg/service/adding"
	"locket-clone/backend/pkg/service/listing"
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
	UserID    uint `gorm:"index:,unique,type:hash"`
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

func (user *User) toListingUser() listing.User {
	friends := make([]listing.Friend, len(user.Friends))
	for i, friend := range user.Friends {
		friends[i] = listing.Friend{
			ID:       friend.ID,
			Email:    friend.Email,
			Username: friend.Username,
			Name:     friend.Name,
		}
	}
	return listing.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Friends:  friends,
	}
}
