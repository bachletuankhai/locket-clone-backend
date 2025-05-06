package listing

type Friend struct {
	Name string `json:"name"`
}

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Friends  []Friend `json:"friends"`
}

type UserService interface {
	GetUserByEmail(string) (User, error)
	GetUserByUsername(string) (User, error)
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

func NewUserService(rp UserRepo) UserService {
	return &userService{
		rp: rp,
	}
}
