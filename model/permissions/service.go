package permissions

import (
	"errors"
	"time"

	"github.com/aditya3232/atmVideoPack-services.git/helper"
)

type Service interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Permissions, helper.Pagination, error)
	GetOne(input PermissionsGetOneByIdInput) (Permissions, error)
	Create(input PermissionsInput) (Permissions, error)
	Update(input PermissionsUpdateInput) (Permissions, error)
	Delete(input PermissionsGetOneByIdInput) error
}

type service struct {
	permissionsRepository Repository
}

func NewService(permissionsRepository Repository) *service {
	return &service{permissionsRepository}
}

func (s *service) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Permissions, helper.Pagination, error) {
	permissions, pagination, err := s.permissionsRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return permissions, pagination, nil
}

func (s *service) GetOne(input PermissionsGetOneByIdInput) (Permissions, error) {
	permission, err := s.permissionsRepository.GetOne(input.ID)
	if err != nil {
		return permission, err
	}

	if permission.ID == 0 {
		return permission, nil
	}

	return permission, nil
}

func (s *service) Create(input PermissionsInput) (Permissions, error) {
	_, err := s.permissionsRepository.GetPermissionName(input.Name)
	if err == nil {
		return Permissions{}, errors.New("permission name must unique")
	}

	now := time.Now()

	permission := Permissions{
		Name:      input.Name,
		CreatedAt: &now,
	}

	newRoles, err := s.permissionsRepository.Create(permission)
	if err != nil {
		return newRoles, err
	}

	return newRoles, nil
}

func (s *service) Update(input PermissionsUpdateInput) (Permissions, error) {
	_, err := s.permissionsRepository.GetOne(input.ID)
	if err != nil {
		return Permissions{}, err
	}

	_, err = s.permissionsRepository.GetPermissionName(input.Name)
	if err == nil {
		return Permissions{}, errors.New("permission name must unique")
	}

	now := time.Now()

	permission := Permissions{
		ID:        input.ID,
		Name:      input.Name,
		UpdatedAt: &now,
	}

	newPermission, err := s.permissionsRepository.Update(permission)
	if err != nil {
		return newPermission, err
	}

	return newPermission, nil
}

func (s *service) Delete(input PermissionsGetOneByIdInput) error {
	_, err := s.permissionsRepository.GetOne(input.ID)
	if err != nil {
		return err
	}

	err = s.permissionsRepository.Delete(input.ID)
	if err != nil {
		return err
	}

	return nil
}
