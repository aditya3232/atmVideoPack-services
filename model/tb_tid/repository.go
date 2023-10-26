package tb_tid

import (
	"github.com/aditya3232/atmVideoPack-services.git/helper"
	"gorm.io/gorm"
)

type Repository interface {
	Create(tbTid TbTid) (TbTid, error)
	GetOneByID(id int) (TbTid, error)
	GetAll(map[string]string, helper.Pagination, helper.Sort) ([]TbTid, helper.Pagination, error)
	CheckUniqueTidInput(tid string) bool
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(tbTid TbTid) (TbTid, error) {
	err := r.db.Create(&tbTid).Error
	if err != nil {
		return TbTid{}, err
	}

	return tbTid, nil
}

func (r *repository) GetOneByID(id int) (TbTid, error) {
	var tbTid TbTid

	err := r.db.Where("id = ?", id).First(&tbTid).Error
	if err != nil {
		return tbTid, err
	}

	return tbTid, nil
}

func (r *repository) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]TbTid, helper.Pagination, error) {
	var tbTids []TbTid
	var total int64

	db := helper.ConstructWhereClause(r.db.Model(&tbTids), filter)

	err := db.Count(&total).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	if total == 0 {
		return tbTids, helper.Pagination{}, nil
	}

	db = helper.ConstructPaginationClause(db, pagination)
	db = helper.ConstructOrderClause(db, sort)

	err = db.Find(&tbTids).Error
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	pagination.Total = int(total)
	pagination.TotalFiltered = len(tbTids)

	return tbTids, pagination, nil

}

func (r *repository) CheckUniqueTidInput(tid string) bool {
	var tbTid TbTid

	err := r.db.Where("tid = ?", tid).First(&tbTid).Error
	return err != nil

}
