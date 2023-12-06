package users

import (
	"time"

	"github.com/aditya3232/atmVideoPack-services.git/helper"
)

type Service interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Users, helper.Pagination, error)
	GetOne(input UsersGetOneByIdInput) (Users, error)
	Create(input UsersInput) (Users, error)
	Update(input UsersUpdateInput) (Users, error)
	Delete(input UsersGetOneByIdInput) error
}

type service struct {
	userRepository Repository
}

func NewService(userRepository Repository) *service {
	return &service{userRepository}
}

func (s *service) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Users, helper.Pagination, error) {
	users, pagination, err := s.userRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return users, pagination, nil
}

func (s *service) GetOne(input UsersGetOneByIdInput) (Users, error) {
	user, err := s.userRepository.GetOne(input.ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, nil
	}

	return user, nil
}

func (s *service) Create(input UsersInput) (Users, error) {
	user := Users{
		RoleId:     input.RoleId,
		Name:       input.Name,
		Username:   input.Username,
		Password:   input.Password,
		FotoProfil: input.FotoProfil,
	}

	newUser, err := s.userRepository.Create(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Update(input UsersUpdateInput) (Users, error) {
	_, err := s.userRepository.GetOne(input.ID)
	if err != nil {
		return Users{}, err
	}

	now := time.Now()

	if input.RoleId == nil || *input.RoleId == 0 {
		roleId := 14
		input.RoleId = &roleId
	}

	user := Users{
		ID:         input.ID,
		RoleId:     input.RoleId,
		Name:       input.Name,
		Username:   input.Username,
		Password:   input.Password,
		FotoProfil: input.FotoProfil,
		UpdatedAt:  &now,
	}

	newUser, err := s.userRepository.Update(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Delete(input UsersGetOneByIdInput) error {
	err := s.userRepository.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}
