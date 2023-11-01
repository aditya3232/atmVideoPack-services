package del_old_log_download_playback_from_elastic

type Service interface {
	DelOneMonthOldDownloadPlaybackLogs() error
}

type service struct {
	delDownloadPlaybackFromElasticRepository Repository
}

func NewService(delDownloadPlaybackFromElasticRepository Repository) *service {
	return &service{delDownloadPlaybackFromElasticRepository}
}

func (s *service) DelOneMonthOldDownloadPlaybackLogs() error {
	err := s.delDownloadPlaybackFromElasticRepository.DelOneMonthOldDownloadPlaybackLogs()
	if err != nil {
		return err
	}

	return nil
}
