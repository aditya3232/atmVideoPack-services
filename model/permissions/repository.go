package permissions

import (
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Permissions, helper.Pagination, error)
	GetOne(id int) (Permissions, error)
	Create(Permissions) (Permissions, error)
	Update(Permissions) (Permissions, error)
	Delete(id int) error
	GetPermissionName(permissionName string) (Permissions, error)
	GetPermissionNameById(id int) (Permissions, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Permissions, helper.Pagination, error) {
	var permissions []Permissions
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&permissions), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return permissions, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&permissions).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(permissions)

	return permissions, pagination, nil
}

func (r *repository) GetOne(id int) (Permissions, error) {
	var permission Permissions

	err := r.db.Where("id = ?", id).First(&permission).Error
	if err != nil {
		return permission, err
	}

	return permission, nil
}

func (r *repository) Create(permission Permissions) (Permissions, error) {
	permission = Permissions{
		Name:      permission.Name,
		CreatedAt: permission.CreatedAt,
	}

	err := r.db.Model(&permission).Create(&permission).Error
	if err != nil {
		return permission, err
	}

	return permission, nil
}

func (r *repository) Update(permission Permissions) (Permissions, error) {
	permission = Permissions{
		ID:        permission.ID,
		Name:      permission.Name,
		UpdatedAt: permission.UpdatedAt,
	}

	err := r.db.Model(&permission).Where("id = ?", permission.ID).Updates(&permission).Error
	if err != nil {
		return permission, err
	}

	return permission, nil
}

func (r *repository) Delete(id int) error {
	var permission Permissions

	err := r.db.Where("id = ?", id).Delete(&permission).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetPermissionName(permissionName string) (Permissions, error) {
	var permission Permissions

	err := r.db.Where("name = ?", permissionName).First(&permission).Error
	if err != nil {
		return permission, err
	}

	return permission, nil
}

func (r *repository) GetPermissionNameById(id int) (Permissions, error) {
	var permissions Permissions

	err := r.db.Where("id = ?", id).First(&permissions).Error
	if err != nil {
		return permissions, err
	}

	return permissions, nil
}
