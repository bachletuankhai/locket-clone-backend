package adding

type LocketType string

var validLocketTypes = [3]LocketType{"image/jpeg", "image/png", "video/mp4"}

type LocketPayload struct {
	Type     LocketType `json:"type"`
	Image    []byte     `json:"-"`
	Caption  string     `json:"caption"`
	Username string     `json:"username"`
}

type LocketRecord struct {
	Type     LocketType `json:"type"`
	ImageUrl string     `json:"imageUrl"`
	Caption  string     `json:"caption"`
	Username string     `json:"username"`
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
	AddLocket(LocketPayload) error
}

type LocketRepo interface {
	AddLocket(LocketRecord) error
}

type locketService struct {
	rp          LocketRepo
	blobStorage ImageBlobStorage
}

func (l *LocketPayload) Validate() error {
	for _, locketType := range validLocketTypes {
		if locketType == l.Type {
			return nil
		}
	}
	return &InvalidPayloadError{message: "invalid locket type"}
}

func (s *locketService) AddLocket(locket LocketPayload) error {
	err := locket.Validate()
	if err != nil {
		return err
	}

	url, err := s.blobStorage.UploadFile(locket.Image, string(locket.Type))
	if err != nil {
		return err
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
