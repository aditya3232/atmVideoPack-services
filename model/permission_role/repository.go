package permission_role

import "gorm.io/gorm"

type Repository interface {
	GetPermissionId(roleId int) ([]PermissionRole, error)
	// Create(PermissionRole) (PermissionRole, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetPermissionId(roleId int) ([]PermissionRole, error) {
	var permission []PermissionRole

	err := r.db.Where("role_id = ?", roleId).Find(&permission).Error
	if err != nil {
		return permission, err
	}

	return permission, nil
}
