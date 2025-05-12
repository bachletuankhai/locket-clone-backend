package listing

import "locket-clone/backend/pkg/model"

type Locket struct {
	model.Locket
}

type LocketRepo interface {
	GetLocket(uint) (Locket, error)
	ListUserLockets(uint, uint, uint) ([]Locket, error)
	ListUserLocketsByUsername(string, uint, uint) ([]Locket, error)
}

type LocketService interface {
	GetLocket(uint) (Locket, error)
	ListUserLocketsByUsername(username string, offset uint, limit uint) ([]Locket, error)
	ListUserLockets(userId uint, offset uint, limit uint) ([]Locket, error)
}

type locketService struct {
	rp LocketRepo
}

func (s *locketService) GetLocket(id uint) (Locket, error) {
	return s.rp.GetLocket(id)
}

func (s *locketService) ListUserLocketsByUsername(username string, offset uint, limit uint) ([]Locket, error) {
	return s.rp.ListUserLocketsByUsername(username, offset, limit)
}

func (s *locketService) ListUserLockets(id uint, offset uint, limit uint) ([]Locket, error) {
	return s.rp.ListUserLockets(id, offset, limit)
}

func NewLocketService(rp LocketRepo) LocketService {
	return &locketService{
		rp: rp,
	}
}
