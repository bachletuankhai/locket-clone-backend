package listing

type Friend struct {
	ID       uint   `json:"-"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
}

type User struct {
	ID       uint     `json:"-"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Friends  []Friend `json:"friends"`
}

type UserService interface {
	GetUserByEmail(string) (User, error)
	GetUserByUsername(string) (User, error)
	GetVisibleUserIds(username string) ([]uint, error)
}

type UserRepo interface {
	GetUserByEmail(string) (User, error)
	GetUserByUsername(string) (User, error)
}

type userService struct {
	rp UserRepo
}

func (s *userService) GetUserByEmail(email string) (User, error) {
	return s.rp.GetUserByEmail(email)
}

func (s *userService) GetUserByUsername(username string) (User, error) {
	return s.rp.GetUserByUsername(username)
}

func (s *userService) GetVisibleUserIds(username string) ([]uint, error) {
	user, err := s.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	visibleUserIds := make([]uint, len(user.Friends))
	for i, friend := range user.Friends {
		visibleUserIds[i] = friend.ID
	}
	visibleUserIds = append(visibleUserIds, user.ID)
	return visibleUserIds, nil
}

func NewUserService(rp UserRepo) UserService {
	return &userService{
		rp: rp,
	}
}
