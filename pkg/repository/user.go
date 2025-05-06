package repository

import (
	"locket-clone/backend/pkg/service/adding"

	"gorm.io/gorm"
)

func (rp *UserRepo) AddUser(user adding.UserRecord) error {
	userRecord := User{
		Model:      gorm.Model{},
		UserRecord: user,
		Friends:    []User{},
	}
	result := rp.db.Create(&userRecord)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
