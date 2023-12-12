package permissions

import "time"

type PermissionsGetFormatter struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type PermissionsCreateFormatter struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
}

type PermissionsUpdateFormatter struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func PermissionsCreateFormat(permissions Permissions) PermissionsCreateFormatter {
	var formatter PermissionsCreateFormatter

	formatter.ID = permissions.ID
	formatter.Name = permissions.Name
	formatter.CreatedAt = permissions.CreatedAt

	return formatter
}

func PermissionsUpdateFormat(permissions Permissions) PermissionsUpdateFormatter {
	var formatter PermissionsUpdateFormatter

	formatter.ID = permissions.ID
	formatter.Name = permissions.Name
	formatter.UpdatedAt = permissions.UpdatedAt

	return formatter
}

func PermissionsGetFormat(permissions Permissions) PermissionsGetFormatter {
	var formatter PermissionsGetFormatter

	formatter.ID = permissions.ID
	formatter.Name = permissions.Name
	formatter.CreatedAt = permissions.CreatedAt
	formatter.UpdatedAt = permissions.UpdatedAt

	return formatter
}

func PermissionsGetAllFormat(permissions []Permissions) []PermissionsGetFormatter {
	formatter := []PermissionsGetFormatter{}

	for _, permission := range permissions {
		permissionGetFormatter := PermissionsGetFormat(permission)
		formatter = append(formatter, permissionGetFormatter)
	}

	return formatter
}
