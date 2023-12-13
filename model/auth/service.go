package auth

import (
	"time"

	jwt "github.com/aditya3232/atmVideoPack-services.git/library/jwt"
	"github.com/aditya3232/atmVideoPack-services.git/model/users"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(input LoginInput) (users.Users, error)
	// Logout(input LogoutInput) error
}

type service struct {
	usersRepository users.Repository
}

func NewService(usersRepository users.Repository) *service {
	return &service{usersRepository}
}

func (s *service) Login(input LoginInput) (users.Users, error) {
	var entityUsers users.Users

	user, err := s.usersRepository.GetUsername(input.Username)
	if err != nil {
		return entityUsers, err
	}

	if user.ID == 0 {
		return entityUsers, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return entityUsers, err
	}

	token, err := jwt.GenerateToken(user.ID, 30)
	if err != nil {
		return entityUsers, err
	}

	now := time.Now()

	entityUsers = users.Users{
		ID:            user.ID,
		RememberToken: token,
		UpdatedAt:     &now,
	}

	loginUser, err := s.usersRepository.Update(entityUsers)
	if err != nil {
		return loginUser, err
	}

	return loginUser, nil
}
