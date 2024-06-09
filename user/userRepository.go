package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]User, error)
	FindUserByID(ID int) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindUserByID(ID int) (User, error) {
	var user User
	err := r.db.First(&user, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, errors.New("User not found")
	}
	return user, err
}

func (r *repository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.Save(&user).Error
	return user, err
}
