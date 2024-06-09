package user

import "fmt"

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
		Role:               "USER",
		UserName:           userRequest.UserName,
		Email:              userRequest.Email,
		Password:           userRequest.Password,
		UserWhatsappNumber: userRequest.UserWhatsappNumber,
		UserGender:         userRequest.UserGender,
		UserIsDeleted:      false,
	}

	user, err := s.userRepository.CreateUser(userData)
	if err != nil {
		return user, err
	}

	user.UserCode = fmt.Sprintf("USER%d", user.ID)

	newUser, err := s.userRepository.UpdateUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, err
}

func (s *service) UpdateUser(ID int, userRequest UserUpdateRequest) (User, error) {
	user, _ := s.userRepository.FindUserByID(ID)

	user.UserName = userRequest.UserName
	user.Email = userRequest.Email
	user.Password = userRequest.Password
	user.UserWhatsappNumber = userRequest.UserWhatsappNumber
	user.UserGender = userRequest.UserGender

	return s.userRepository.UpdateUser(user)
}
