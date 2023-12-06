package users

import (
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]Users, helper.Pagination, error)
	GetOne(id int) (Users, error)
	Create(Users) (Users, error)
	Update(Users) (Users, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]Users, helper.Pagination, error) {
	var users []Users
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&users), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return users, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&users).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(users)

	return users, pagination, nil
}

func (r *repository) GetOne(id int) (Users, error) {
	var user Users

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Create(user Users) (Users, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user Users) (Users, error) {
	if user.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		if err != nil {
			return user, err
		}

		user.Password = string(password)
	}

	user = Users{
		ID:         user.ID,
		RoleId:     user.RoleId,
		Name:       user.Name,
		Password:   user.Password,
		FotoProfil: user.FotoProfil,
		UpdatedAt:  user.UpdatedAt,
	}

	err := r.db.Model(&user).Where("id = ?", user.ID).Updates(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Delete(id int) error {
	var user Users

	err := r.db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
