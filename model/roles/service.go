package roles

import (
	"errors"
	"time"

	"github.com/aditya3232/atmVideoPack-services.git/helper"
)

type Service interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Roles, helper.Pagination, error)
	GetOne(input RolesGetOneByIdInput) (Roles, error)
	Create(input RolesInput) (Roles, error)
	Update(input RolesUpdateInput) (Roles, error)
	Delete(input RolesGetOneByIdInput) error
}

type service struct {
	rolesRepository Repository
}

func NewService(rolesRepository Repository) *service {
	return &service{rolesRepository}
}

func (s *service) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Roles, helper.Pagination, error) {
	roles, pagination, err := s.rolesRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return roles, pagination, nil
}

func (s *service) GetOne(input RolesGetOneByIdInput) (Roles, error) {
	role, err := s.rolesRepository.GetOne(input.ID)
	if err != nil {
		return role, err
	}

	if role.ID == 0 {
		return role, nil
	}

	return role, nil
}

func (s *service) Create(input RolesInput) (Roles, error) {
	_, err := s.rolesRepository.GetRoleName(input.Name)
	if err == nil {
		return Roles{}, errors.New("role name must unique")
	}

	now := time.Now()

	role := Roles{
		Name:      input.Name,
		CreatedAt: &now,
	}

	newRoles, err := s.rolesRepository.Create(role)
	if err != nil {
		return newRoles, err
	}

	return newRoles, nil
}

func (s *service) Update(input RolesUpdateInput) (Roles, error) {
	_, err := s.rolesRepository.GetOne(input.ID)
	if err != nil {
		return Roles{}, err
	}

	now := time.Now()

	role := Roles{
		ID:        input.ID,
		Name:      input.Name,
		UpdatedAt: &now,
	}

	newRoles, err := s.rolesRepository.Update(role)
	if err != nil {
		return newRoles, err
	}

	return newRoles, nil
}

func (s *service) Delete(input RolesGetOneByIdInput) error {
	_, err := s.rolesRepository.GetOne(input.ID)
	if err != nil {
		return err
	}

	err = s.rolesRepository.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}
