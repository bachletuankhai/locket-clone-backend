package repository

import (
	"errors"
	"locket-clone/backend/pkg/service/adding"

	"gorm.io/gorm"
)

func (rp *LocketRepo) AddLocket(locket adding.LocketRecord) error {
	locketRecord := Locket{
		Type:     locket.Type,
		Model:    gorm.Model{},
		ImageUrl: locket.ImageUrl,
		Caption:  locket.Caption,
	}
	result := rp.db.Create(&locketRecord)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (rp *LocketRepo) GetLocket(id uint) (Locket, error) {
	locket := Locket{}
	result := rp.db.First(&locket, id)
	if result.Error != nil {
		return locket, errors.New("error fetching locket")
	}
	return locket, nil
}

func (rp *LocketRepo) ListUserLocketsByUsername(username string, offset uint, limit uint) ([]Locket, error) {
	var lockets []Locket
	locket := Locket{
		User: User{
			UserRecord: adding.UserRecord{
				Username: username,
			},
		},
	}
	result := rp.db.Model(&locket).Order("CreatedAt DESC").Offset(int(offset)).Limit(int(limit)).Find(&lockets)
	if result.Error != nil {
		return lockets, errors.New("error fetching locket")
	}
	return lockets, nil
}

func (rp *LocketRepo) ListUserLockets(userId uint, offset uint, limit uint) ([]Locket, error) {
	var lockets []Locket
	result := rp.db.Model(&Locket{
		UserID: userId,
	}).Order("CreatedAt DESC").Offset(int(offset)).Limit(int(limit)).Find(&lockets)
	if result.Error != nil {
		return lockets, errors.New("error fetching locket")
	}
	return lockets, nil
}
