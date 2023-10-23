package tb_tid

import "github.com/aditya3232/atmVideoPack-services.git/helper"

type Service interface {
	Create(tbTidInput TbTidCreateInput) (TbTid, error)
	GetOneByID(input GetOneByIDInput) (TbTid, error)
	GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]TbTid, helper.Pagination, error)
}

type service struct {
	tbTidRepository Repository
}

func NewService(tbTidRepository Repository) *service {
	return &service{tbTidRepository}
}

func (s *service) Create(tbTidInput TbTidCreateInput) (TbTid, error) {
	tbTid := TbTid{
		Tid:        tbTidInput.Tid,
		IpAddress:  tbTidInput.IpAddress,
		SnMiniPc:   tbTidInput.SnMiniPc,
		LocationId: tbTidInput.LocationId,
	}

	newTbTid, err := s.tbTidRepository.Create(tbTid)
	if err != nil {
		return newTbTid, err
	}

	return newTbTid, nil
}

func (s *service) GetOneByID(input GetOneByIDInput) (TbTid, error) {
	tbTid, err := s.tbTidRepository.GetOneByID(input.ID)
	if err != nil {
		return tbTid, err
	}
	if tbTid.ID == 0 {
		return tbTid, nil
	}

	return tbTid, nil
}

func (s *service) GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]TbTid, helper.Pagination, error) {
	tbEntries, pagination, err := s.tbTidRepository.GetAll(filter, pagination, sort)
	if err != nil {
		return nil, helper.Pagination{}, err
	}

	return tbEntries, pagination, nil
}
