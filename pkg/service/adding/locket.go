package adding

import (
	"locket-clone/backend/pkg/model"
	"locket-clone/backend/pkg/service/listing"
)

type LocketPayload struct {
	Type     model.LocketType `json:"type" form:"type" binding:"required"`
	Image    []byte           `json:"-" form:"image" binding:"required"`
	Caption  string           `json:"caption" form:"caption" binding:"required"`
	Username string           `json:"username" form:"-" binding:"-"`
}

type LocketRecord struct {
	Type     model.LocketType `json:"type"`
	ImageUrl string           `json:"imageUrl"`
	Caption  string           `json:"caption"`
	Username string           `json:"username"`
}

type InvalidPayloadError struct {
	message string
}

func (e *InvalidPayloadError) Error() string {
	return e.message
}

type ImageBlobStorage interface {
	UploadFile([]byte, string) (string, error)
}

type LocketService interface {
	AddLocket(LocketPayload) (listing.Locket, error)
}

type LocketRepo interface {
	AddLocket(LocketRecord) (listing.Locket, error)
}

type locketService struct {
	rp          LocketRepo
	blobStorage ImageBlobStorage
}

func (l *LocketPayload) Validate() error {
	for _, locketType := range model.ValidLocketTypes {
		if locketType == l.Type {
			return nil
		}
	}
	return &InvalidPayloadError{message: "invalid locket type"}
}

func (s *locketService) AddLocket(locket LocketPayload) (listing.Locket, error) {
	err := locket.Validate()
	if err != nil {
		return listing.Locket{}, err
	}

	url, err := s.blobStorage.UploadFile(locket.Image, string(locket.Type))
	if err != nil {
		return listing.Locket{}, err
	}
	return s.rp.AddLocket(LocketRecord{
		Type:     locket.Type,
		ImageUrl: url,
		Caption:  locket.Caption,
		Username: locket.Username,
	})
}

func NewLocketService(rp LocketRepo, storage ImageBlobStorage) LocketService {
	return &locketService{
		rp:          rp,
		blobStorage: storage,
	}
}
