package download_playback

import (
	"bytes"

	"github.com/aditya3232/atmVideoPack-services.git/model/tb_tid"
)

type Service interface {
	DownloadPlayback(input ServiceDownloadPlaybackInput) (*bytes.Buffer, error)
}

type service struct {
	downloadPlaybackRepository Repository
	tbTidRepository            tb_tid.Repository
}

func NewService(downloadPlaybackRepository Repository, tbTidRepository tb_tid.Repository) *service {
	return &service{downloadPlaybackRepository, tbTidRepository}
}

func (s *service) DownloadPlayback(input ServiceDownloadPlaybackInput) (*bytes.Buffer, error) {
	// get ip_address
	tbTid, err := s.tbTidRepository.GetOneByTid(input.Tid)
	if err != nil {
		return nil, err
	}

	input.IpAddress = tbTid.IpAddress

	response, err := s.downloadPlaybackRepository.DownloadPlayback(input)
	if err != nil {
		return nil, err
	}

	return response, nil
}
