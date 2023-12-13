package profile

import (
	"github.com/aditya3232/atmVideoPack-services.git/model/permission_role"
	"github.com/aditya3232/atmVideoPack-services.git/model/permissions"
	"github.com/aditya3232/atmVideoPack-services.git/model/roles"
	"github.com/aditya3232/atmVideoPack-services.git/model/users"
)

type Service interface {
	GetProfile(userID int) (Profile, error)
}

type service struct {
	userRepository           users.Repository
	roleRepository           roles.Repository
	permissionRepository     permissions.Repository
	permissionRoleRepository permission_role.Repository
}

func NewService(userRepository users.Repository, roleRepository roles.Repository, permissionRepository permissions.Repository, permissionRoleRepository permission_role.Repository) *service {
	return &service{userRepository, roleRepository, permissionRepository, permissionRoleRepository}
}

func (s *service) GetProfile(userID int) (Profile, error) {
	var entityProfile Profile

	user, err := s.userRepository.GetOne(userID)
	if err != nil {
		return entityProfile, err
	}

	role, err := s.roleRepository.GetOne(*user.RoleId)
	if err != nil {
		return entityProfile, err
	}

	/*
		- disini ada banyak data permission_id
	*/
	permissionRole, err := s.permissionRoleRepository.GetPermissionId(role.ID)
	if err != nil {
		return entityProfile, err
	}

	/*
		- GetPermissionNameById return slice of permissions
		- disini ada banyak data permission_data, berdasarkan id masing2
		- maka dari itu disini ada perluangan
		- nanti ada variable untuk menampung data didalam slice
		- lalu data dari perluangan harus dimasukkan kedalam variable tersebut dengan cara append
		- nah ada kasus kenapa permission..., karena GetPermissionNameById mengembalikasi slice, bukan satu data
	*/
	var permissions []permissions.Permissions // Slice to store multiple permissions
	for _, pr := range permissionRole {
		permission, err := s.permissionRepository.GetPermissionNameById(pr.PermissionID)
		if err != nil {
			return entityProfile, err
		}

		permissions = append(permissions, permission)
	}

	entityProfile = Profile{
		User:        user,
		Role:        role,
		Permissions: permissions,
	}

	return entityProfile, nil
}
