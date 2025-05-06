package listing

import "locket-clone/backend/pkg/model"

type LocketRecord struct {
	model.Locket
}

type LocketRepo interface {
	GetLocket(uint) (LocketRecord, error)
	ListUserLockets(uint) ([]LocketRecord, error)
	ListUserLocketsByUsername(string) ([]LocketRecord, error)
}

type LocketService interface {
	GetLocket(uint) (LocketRecord, error)
	ListUserLocketsByUsername(string) ([]LocketRecord, error)
	ListUserLockets(uint) ([]LocketRecord, error)
}

type locketService struct {
	rp LocketRepo
}

func (s *locketService) GetLocket(id uint) (LocketRecord, error) {
	return s.rp.GetLocket(id)
}

func (s *locketService) ListUserLocketsByUsername(username string) ([]LocketRecord, error) {
	return s.rp.ListUserLocketsByUsername(username)
}

func (s *locketService) ListUserLockets(id uint) ([]LocketRecord, error) {
	return s.rp.ListUserLockets(id)
}

func NewLocketService(rp LocketRepo) LocketService {
	return &locketService{
		rp: rp,
	}
}
