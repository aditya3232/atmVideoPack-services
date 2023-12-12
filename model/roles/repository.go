package roles

import (
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Roles, helper.Pagination, error)
	GetOne(id int) (Roles, error)
	Create(Roles) (Roles, error)
	Update(Roles) (Roles, error)
	Delete(id int) error
	GetRoleName(roleName string) (Roles, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Roles, helper.Pagination, error) {
	var roles []Roles
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&roles), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return roles, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&roles).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(roles)

	return roles, pagination, nil
}

func (r *repository) GetOne(id int) (Roles, error) {
	var role Roles

	err := r.db.Where("id = ?", id).First(&role).Error
	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *repository) Create(role Roles) (Roles, error) {
	role = Roles{
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
	}

	err := r.db.Model(&role).Create(&role).Error
	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *repository) Update(role Roles) (Roles, error) {
	role = Roles{
		ID:        role.ID,
		Name:      role.Name,
		UpdatedAt: role.UpdatedAt,
	}

	err := r.db.Model(&role).Where("id = ?", role.ID).Updates(&role).Error
	if err != nil {
		return role, err
	}

	return role, nil
}

func (r *repository) Delete(id int) error {
	var role Roles

	err := r.db.Where("id = ?", id).Delete(&role).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetRoleName(roleName string) (Roles, error) {
	var role Roles

	err := r.db.Where("name = ?", roleName).First(&role).Error
	if err != nil {
		return role, err
	}

	return role, nil
}
