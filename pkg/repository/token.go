package repository

import (
	"locket-clone/backend/pkg/service/auth"
	"time"
)

type MemTokenRepo struct {
	mem map[string]time.Time
}

type ErrTokenNotFound struct {
	Err string
}

func (e *ErrTokenNotFound) Error() string {
	return e.Err
}

func (m *MemTokenRepo) SaveToken(token string, exp time.Time) error {
	if len(token) == 0 {
		return &ErrTokenNotFound{
			Err: "emtpy token",
		}
	}
	if exp.IsZero() || exp.Before(time.Now()) {
		return nil
	}

	// Token already in the blacklist, shouldn't reach here but just in case
	// we update the expiration time
	if later, exist := m.mem[token]; exist {
		if exp.After(later) {
			later = exp
		}
		m.mem[token] = later
		return nil
	}

	m.mem[token] = exp
	return nil
}

func (m *MemTokenRepo) CheckTokenExists(token string) (bool, error) {
	if len(token) == 0 {
		return false, &ErrTokenNotFound{
			Err: "emtpy token",
		}
	}
	exp, ok := m.mem[token]
	if !ok {
		return false, nil
	}
	if exp.IsZero() || exp.Before(time.Now()) {
		delete(m.mem, token)
		return false, nil
	}
	return true, nil
}

func NewMemTokenRepo() auth.TokenRepo {
	return &MemTokenRepo{
		mem: make(map[string]time.Time),
	}
}
