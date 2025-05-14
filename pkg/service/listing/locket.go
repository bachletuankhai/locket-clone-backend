package listing

import (
	"locket-clone/backend/pkg/model"
	"time"
)

type Locket struct {
	model.Locket
}

type LocketRepo interface {
	GetLocket(uint) (Locket, error)
	ListLocketsByUserIdsTime(userIds []uint, startTime time.Time, limit uint) ([]Locket, error)
	ListLatestLockets(userIds []uint, limit uint) ([]Locket, error)
	ListUserLocketsByUsername(username string, limit uint) ([]Locket, error)
	ListUserLocketsByUsernameTime(username string, startTime time.Time, limit uint) ([]Locket, error)
}

type LocketService interface {
	ListLocketsByUserIdsTime(userIds []uint, startTime time.Time, limit uint) ([]Locket, error)
	ListLatestLockets(userIds []uint, limit uint) ([]Locket, error)
	ListUserLocketsByUsername(username string, limit uint) ([]Locket, error)
	ListUserLocketsByUsernameTime(username string, startTime time.Time, limit uint) ([]Locket, error)
}

type locketService struct {
	rp LocketRepo
}

func (s *locketService) ListLocketsByUserIdsTime(userIds []uint, startTime time.Time, limit uint) ([]Locket, error) {
	return s.rp.ListLocketsByUserIdsTime(userIds, startTime, limit)
}

func (s *locketService) ListLatestLockets(userIds []uint, limit uint) ([]Locket, error) {
	return s.rp.ListLatestLockets(userIds, limit)
}

func (s *locketService) ListUserLocketsByUsername(username string, limit uint) ([]Locket, error) {
	return s.rp.ListUserLocketsByUsername(username, limit)
}

func (s *locketService) ListUserLocketsByUsernameTime(username string, startTime time.Time, limit uint) ([]Locket, error) {
	return s.rp.ListUserLocketsByUsernameTime(username, startTime, limit)
}

func NewLocketService(rp LocketRepo) LocketService {
	return &locketService{
		rp: rp,
	}
}
