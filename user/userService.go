package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserService interface {
	FindAll() ([]User, error)
	FindUserByID(ID int) (User, error)
	CreateUser(user UserCreateRequest) (User, error)
	UpdateUser(ID int, user UserUpdateRequest) (User, error)
}

type service struct {
	userRepository UserRepository
}

func NewService(userRepository UserRepository) *service {
	return &service{userRepository}
}

func (s *service) FindAll() ([]User, error) {
	return s.userRepository.FindAll()
}

func (s *service) FindUserByID(ID int) (User, error) {
	return s.userRepository.FindUserByID(ID)
}

func (s *service) CreateUser(userRequest UserCreateRequest) (User, error) {
	userData := User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Whatsapp: userRequest.Whatsapp,
		Gender:   userRequest.Gender,
		Role:     userRequest.Role,
	}

	user, err := s.userRepository.CreateUser(userData)

	return user, err
}

func (s *service) UpdateUser(ID int, userRequest UserUpdateRequest) (User, error) {
	user, err := s.userRepository.FindUserByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, errors.New("User not found")
	}

	if err != nil {
		return User{}, err
	}

	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.Password = userRequest.Password
	user.Whatsapp = userRequest.Whatsapp
	user.Gender = userRequest.Gender
	user.Role = userRequest.Role

	return s.userRepository.UpdateUser(user)
}
