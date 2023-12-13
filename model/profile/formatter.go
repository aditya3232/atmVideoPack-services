package profile

import (
	permissions_model "github.com/aditya3232/atmVideoPack-services.git/model/permissions"
	roles_model "github.com/aditya3232/atmVideoPack-services.git/model/roles"
	users_model "github.com/aditya3232/atmVideoPack-services.git/model/users"
)

type ProfileGetFormatter struct {
	User        users_model.UsersGetFormatter               `json:"user"`
	Role        roles_model.RolesGetFormatter               `json:"role"`
	Permissions []permissions_model.PermissionsGetFormatter `json:"permissions"`
}

func ProfileGetFormat(user users_model.Users, role roles_model.Roles, permissions []permissions_model.Permissions) ProfileGetFormatter {
	var formatter ProfileGetFormatter

	formatter.User = users_model.UsersGetFormat(user)
	formatter.Role = roles_model.RolesGetFormat(role)
	formatter.Permissions = permissions_model.PermissionsGetAllFormat(permissions)

	return formatter
}
