package tb_tid

import "gorm.io/gorm"

type Repository interface {
	Create(tbTid TbTid) (TbTid, error)
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
