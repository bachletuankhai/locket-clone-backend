package listing

import "locket-clone/backend/pkg/model"

type Locket struct {
	model.Locket
}

type LocketRepo interface {
	GetLocket(uint) (Locket, error)
	ListUserLockets(uint) ([]Locket, error)
	ListUserLocketsByUsername(string) ([]Locket, error)
}

type LocketService interface {
	GetLocket(uint) (Locket, error)
	ListUserLocketsByUsername(string) ([]Locket, error)
	ListUserLockets(uint) ([]Locket, error)
}

type locketService struct {
	rp LocketRepo
}

func (s *locketService) GetLocket(id uint) (Locket, error) {
	return s.rp.GetLocket(id)
}

func (s *locketService) ListUserLocketsByUsername(username string) ([]Locket, error) {
	return s.rp.ListUserLocketsByUsername(username)
}

func (s *locketService) ListUserLockets(id uint) ([]Locket, error) {
	return s.rp.ListUserLockets(id)
}

func NewLocketService(rp LocketRepo) LocketService {
	return &locketService{
		rp: rp,
	}
}
