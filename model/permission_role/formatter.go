package permission_role

type PermissionRoleGetFormatter struct {
	PermissionID int `json:"permission_id"`
	RoleID       int `json:"role_id"`
}

type PermissionRoleCreateFormatter struct {
	PermissionID int `json:"permission_id"`
	RoleID       int `json:"role_id"`
}

func PermissionRoleGetFormat(permissionRole PermissionRole) PermissionRoleGetFormatter {
	var formatter PermissionRoleGetFormatter

	formatter.PermissionID = permissionRole.PermissionID
	formatter.RoleID = permissionRole.RoleID

	return formatter
}

func PermissionRoleCreateFormat(permissionRole PermissionRole) PermissionRoleCreateFormatter {
	var formatter PermissionRoleCreateFormatter

	formatter.PermissionID = permissionRole.PermissionID
	formatter.RoleID = permissionRole.RoleID

	return formatter
}

func PermissionRoleGetAllFormat(permissionRole []PermissionRole) []PermissionRoleGetFormatter {
	formatter := []PermissionRoleGetFormatter{}

	for _, permissionRole := range permissionRole {
		permissionRoleGetFormatter := PermissionRoleGetFormat(permissionRole)
		formatter = append(formatter, permissionRoleGetFormatter)
	}

	return formatter
}
