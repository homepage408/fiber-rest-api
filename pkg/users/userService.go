package users

type UserService interface {
	FindAllUser() ([]User, error)
	FindUserById(id int) (User, error)
	SignUp(user UserRequest) (User, error)
	SignIn(email string) (User, error)
}

type serviceUser struct {
	userRepository UserRepository
}

func NewServiceUser(userRepository UserRepository) *serviceUser {
	return &serviceUser{userRepository: userRepository}
}

func (s *serviceUser) FindAllUser() ([]User, error) {
	users, err := s.userRepository.FindAllUser()

	return users, err
}

func (s *serviceUser) FindUserById(id int) (User, error) {
	user, err := s.userRepository.FindUserById(id)
	return user, err
}

func (s *serviceUser) SignUp(userRequest UserRequest) (User, error) {

	userIn := User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Username:  userRequest.Username,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		Role:      userRequest.Role,
	}

	user, err := s.userRepository.SignUp(userIn)
	return user, err
}

func (s *serviceUser) SignIn(email string) (User, error) {
	user, err := s.userRepository.SignIn(email)
	return user, err
}
