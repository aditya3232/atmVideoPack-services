package permission_role

type PermissionRoleInput struct {
	PermissionID int `form:"permission_id" binding:"required"`
	RoleID       int `form:"role_id" binding:"required"`
}
