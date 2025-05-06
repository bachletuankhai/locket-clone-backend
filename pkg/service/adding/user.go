package adding

import "locket-clone/backend/pkg/service/auth"

type UserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Username string `json:"username"`
}

type UserRecord struct {
	Name         string `json:"name" gorm:"not null;default:null"`
	Email        string `json:"email" gorm:"uniqueIndex;not null;default:null"`
	PasswordHash string `json:"-"`
	Username     string `json:"username" gorm:"uniqueIndex;not null;default:null"`
}

type UserRepo interface {
	AddUser(UserRecord) error
}

type UserService interface {
	AddUser(UserPayload) error
}

type service struct {
	repo UserRepo
}

func (s *service) AddUser(user UserPayload) error {

	passwordHash, err := auth.HashPassword(user.Password)
	if err != nil {
		return err
	}
	return s.repo.AddUser(UserRecord{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: passwordHash,
	})
}

func NewUserService(rp UserRepo) UserService {
	return &service{repo: rp}
}
