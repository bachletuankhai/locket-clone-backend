package repository

import (
	"errors"
	"locket-clone/backend/pkg/service/adding"
	"time"

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

func (rp *LocketRepo) ListLocketsByUserIdsTime(userIds []uint, startTime time.Time, limit uint) ([]Locket, error) {
	if len(userIds) == 0 {
		return []Locket{}, nil
	}
	var lockets []Locket
	result := rp.db.Where("UserID IN ? AND CreateAt <= ?", userIds, startTime).Order("CreatedAt DESC").Limit(int(limit)).Find(&lockets)
	if result.Error != nil {
		return lockets, result.Error
	}
	return lockets, nil
}

func (rp *LocketRepo) ListLatestLockets(userIds []uint, limit uint) ([]Locket, error) {
	if len(userIds) == 0 {
		return []Locket{}, nil
	}

	var lockets []Locket
	result := rp.db.Where("UserID IN ?", userIds).Order("CreatedAt DESC").Limit(int(limit)).Find(&lockets)
	if result.Error != nil {
		return lockets, result.Error
	}
	return lockets, nil
}

func (rp *LocketRepo) ListUserLocketsByUsername(username string, limit uint) ([]Locket, error) {
	var lockets []Locket
	locket := Locket{
		User: User{
			UserRecord: adding.UserRecord{
				Username: username,
			},
		},
	}
	result := rp.db.Model(&locket).Order("CreatedAt DESC").Limit(int(limit)).Find(&lockets)
	if result.Error != nil {
		return lockets, errors.New("error fetching locket")
	}
	return lockets, nil
}

func (rp *LocketRepo) ListUserLocketsByUsernameTime(username string, startTime time.Time, limit uint) ([]Locket, error) {
	var lockets []Locket
	locket := Locket{
		User: User{
			UserRecord: adding.UserRecord{
				Username: username,
			},
		},
	}
	result := rp.db.Model(&locket).Order("CreatedAt DESC").Where("CreatedAt <= ?", startTime).Limit(int(limit)).Find(&lockets)
	if result.Error != nil {
		return lockets, errors.New("error fetching locket")
	}
	return lockets, nil
}
