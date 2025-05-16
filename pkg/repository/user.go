package repository

import (
	"locket-clone/backend/pkg/service/adding"
	"locket-clone/backend/pkg/service/listing"

	"gorm.io/gorm"
)

func (rp *UserRepo) AddUser(user adding.UserRecord) error {
	userRecord := User{
		Model:      gorm.Model{},
		UserRecord: user,
		Friends:    []User{},
	}
	result := rp.Db.Create(&userRecord)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (rp *UserRepo) GetUserByUsername(username string) (listing.User, error) {
	user := User{
		UserRecord: adding.UserRecord{
			Username: username,
		},
	}
	result := rp.Db.First(&user)
	if result.Error != nil {
		return listing.User{}, result.Error
	}
	return user.toListingUser(), nil
}

func (rp *UserRepo) GetUserByEmail(email string) (listing.User, error) {
	user := User{
		UserRecord: adding.UserRecord{
			Email: email,
		},
	}
	result := rp.Db.First(&user)
	if result.Error != nil {
		return listing.User{}, result.Error
	}
	return user.toListingUser(), nil
}

func (rp *UserRepo) GetUserPasswordHashByUsername(username string) (string, error) {
	user := User{
		UserRecord: adding.UserRecord{
			Username: username,
		},
	}
	result := rp.Db.First(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.PasswordHash, nil
}
