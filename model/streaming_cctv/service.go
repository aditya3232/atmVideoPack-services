package streaming_cctv

import (
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/model/tb_tid"
)

type Service interface {
	StreamingCctv(input StreamingCctvInput) (*http.Response, error)
}

type service struct {
	streamingCctvRepository Repository
	tbTidRepository         tb_tid.Repository
}

func NewService(streamingCctvRepository Repository, tbTidRepository tb_tid.Repository) *service {
	return &service{streamingCctvRepository, tbTidRepository}
}

func (s *service) StreamingCctv(input StreamingCctvInput) (*http.Response, error) {
	// get ip_address
	tbTid, err := s.tbTidRepository.GetOneByID(input.ID)
	if err != nil {
		return nil, err
	}

	input.IpAddress = tbTid.IpAddress

	response, err := s.streamingCctvRepository.StreamingCctv(input)
	if err != nil {
		return nil, err
	}

	return response, nil
}
