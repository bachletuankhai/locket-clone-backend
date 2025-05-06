package repository

import (
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
