package users

import "gorm.io/gorm"

type UserRepository interface {
	FindAllUser() ([]User, error)
	FindUserById(id int) (User, error)
	SignUp(user User) (User, error)
	SignIn(email string) (User, error)
}

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) *userRepository {
	return &userRepository{Db}
}

func (r *userRepository) FindAllUser() ([]User, error) {
	err := r.Db.Find(&User{}).Error
	return []User{}, err
}

func (r *userRepository) FindUserById(id int) (User, error) {
	err := r.Db.First(&User{}, id).Error
	return User{}, err
}

func (r *userRepository) SignUp(user User) (User, error) {
	err := r.Db.Create(&user).Error
	return user, err
}

func (r *userRepository) SignIn(email string) (User, error) {
	err := r.Db.Where("email = ? ", email).First(&User{Email: email}).Error
	return User{}, err
}
