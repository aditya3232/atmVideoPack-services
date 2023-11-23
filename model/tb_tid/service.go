package tb_tid

import (
	"errors"

	"github.com/aditya3232/atmVideoPack-services.git/helper"
	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
)

type Service interface {
	Create(tbTidInput TbTidCreateInput) (TbTid, error)
	GetOneByID(input GetOneByIDInput) (TbTid, error)
	GetOneByTid(tid string) (TbTid, error)
	GetAll(filter map[string]string, pagination helper.Pagination, sort helper.Sort) ([]TbTid, helper.Pagination, error)
	CheckUniqueTidInput(tid string) bool
}

type service struct {
	tbTidRepository Repository
}

/*
tbTidRepository := tb_tid.NewRepository(connection.DatabaseMysql())
tbTidService := tb_tid.NewService(tbTidRepository)
tbTidHandler := handler.NewTbTidHandler(tbTidService)

configureTbTidRoutes(tbTidRoutes, tbTidHandler)

	func configureTbTidRoutes(group *gin.RouterGroup, handler *handler.TbTidHandler) {
		group.POST("/create", handler.CreateTbTid)
		group.POST("/getonebyid", handler.GetOneByID)
		group.GET("/getall", handler.GetAllTbEntries)
	}
*/

/*
- fungsi NewRepository, NewService, NewHandler
- adalah contoh penggunaan konsep dependency injection (DI)
- keuntungannya :
- Pemahaman Ketergantungan: Membuat struktur dan ketergantungan antar komponen lebih jelas dan mudah dipahami.
- Penerapan Dependency Injection dapat membantu mencapai desain yang lebih bersih, terstruktur, dan mudah diubah pada aplikasi Anda.
- membuat sistem lebih fleksibel, mudah diuji, dan mudah dipelihara.
*/
func NewService(tbTidRepository Repository) *service {
	return &service{tbTidRepository}
}

/*
- (s *service)
- ini adalah struct method
- dengan adanya struct method, maka kita dapat memanggil structnya
- berikut adalah pemanggilan structnya :
- unique := s.tbTidRepository.CheckUniqueTidInput(tbTidInput.Tid)
- alasan kenapa di struct method menggunakan asterik(pointer) adalah :
- Secara umum, Anda harus menggunakan penerima pointer saat Anda perlu memodifikasi keadaan struct yang dipanggil metode. Jika Anda tidak perlu memodifikasi keadaan struct, Anda dapat menggunakan penerima nilai.
*/
func (s *service) Create(tbTidInput TbTidCreateInput) (TbTid, error) {
	var tbTid TbTid

	// check unique tid
	unique := s.tbTidRepository.CheckUniqueTidInput(tbTidInput.Tid)
	if !unique {
		log_function.Error("add device error, tid is not unique")
		return tbTid, errors.New("tid is not unique")
	}

	tbTid = TbTid{
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

func (s *service) GetOneByTid(tid string) (TbTid, error) {
	tbTid, err := s.tbTidRepository.GetOneByTid(tid)
	if err != nil {
		return tbTid, err
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

func (s *service) CheckUniqueTidInput(tid string) bool {
	unique := s.tbTidRepository.CheckUniqueTidInput(tid)
	return unique
}
