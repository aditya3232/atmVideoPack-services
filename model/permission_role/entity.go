package permission_role

type PermissionRole struct {
	RoleID       int `json:"role_id"`
	PermissionID int `json:"permission_id"`
}

func (m *PermissionRole) TableName() string {
	return "permission_role"
}
