package get_download_playback_from_elastic

type Service interface {
	FindAll(findAllElasticDownloadPlaybackInput FindAllElasticDownloadPlaybackInput) ([]ElasticDownloadPlayback, error)
}

type service struct {
	elasticDownloadPlaybackRepository Repository
}

func NewService(elasticDownloadPlaybackRepository Repository) *service {
	return &service{elasticDownloadPlaybackRepository}
}

func (s *service) FindAll(findAllElasticDownloadPlaybackInput FindAllElasticDownloadPlaybackInput) ([]ElasticDownloadPlayback, error) {

	elasticDownloadPlaybacks, err := s.elasticDownloadPlaybackRepository.FindAll(findAllElasticDownloadPlaybackInput)
	if err != nil {
		return nil, err
	}

	return elasticDownloadPlaybacks, nil
}
