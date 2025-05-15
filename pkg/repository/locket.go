package repository

import (
	"errors"
	"locket-clone/backend/pkg/service/adding"
	"locket-clone/backend/pkg/service/listing"
	"time"

	"gorm.io/gorm"
)

func (rp *LocketRepo) AddLocket(locket adding.LocketRecord) (listing.Locket, error) {
	locketRecord := Locket{
		Type:     locket.Type,
		Model:    gorm.Model{},
		ImageUrl: locket.ImageUrl,
		Caption:  locket.Caption,
	}
	result := rp.Db.Create(&locketRecord)
	if result.Error != nil {
		return locketRecord.toListingLocket(), result.Error
	}
	return locketRecord.toListingLocket(), nil
}

func (rp *LocketRepo) GetLocket(id uint) (listing.Locket, error) {
	locket := Locket{}
	result := rp.Db.First(&locket, id)
	if result.Error != nil {
		return listing.Locket{}, errors.New("error fetching locket")
	}
	return locket.toListingLocket(), nil
}

func (rp *LocketRepo) ListLocketsByUserIdsTime(userIds []uint, startTime time.Time, limit uint) ([]listing.Locket, error) {
	if len(userIds) == 0 {
		return []listing.Locket{}, nil
	}
	var lockets []listing.Locket
	result := rp.Db.Where("UserID IN ? AND CreateAt <= ?", userIds, startTime).Order("CreatedAt DESC").Limit(int(limit)).Find(&lockets)
	if result.Error != nil {
		return []listing.Locket{}, result.Error
	}
	return lockets, nil
}

func (rp *LocketRepo) ListLatestLockets(userIds []uint, limit uint) ([]listing.Locket, error) {
	if len(userIds) == 0 {
		return []listing.Locket{}, nil
	}

	var lockets []listing.Locket
	result := rp.Db.Where("UserID IN ?", userIds).Order("CreatedAt DESC").Limit(int(limit)).Find(&lockets)
	if result.Error != nil {
		return lockets, result.Error
	}
	return lockets, nil
}

func (rp *LocketRepo) ListUserLocketsByUsername(username string, limit uint) ([]listing.Locket, error) {
	var lockets []Locket
	locket := Locket{
		User: User{
			UserRecord: adding.UserRecord{
				Username: username,
			},
		},
	}
	result := rp.Db.Model(&locket).Order("CreatedAt DESC").Limit(int(limit)).Find(&lockets)
	if result.Error != nil {
		return []listing.Locket{}, errors.New("error fetching locket")
	}
	locketList := make([]listing.Locket, len(lockets))
	for i, l := range lockets {
		locketList[i] = l.toListingLocket()
	}
	return locketList, nil
}

func (rp *LocketRepo) ListUserLocketsByUsernameTime(username string, startTime time.Time, limit uint) ([]listing.Locket, error) {
	var lockets []Locket
	locket := Locket{
		User: User{
			UserRecord: adding.UserRecord{
				Username: username,
			},
		},
	}
	result := rp.Db.Model(&locket).Order("CreatedAt DESC").Where("CreatedAt <= ?", startTime).Limit(int(limit)).Find(&lockets)
	if result.Error != nil {
		return []listing.Locket{}, errors.New("error fetching locket")
	}
	locketList := make([]listing.Locket, len(lockets))
	for i, l := range lockets {
		locketList[i] = l.toListingLocket()
	}
	return locketList, nil
}
