package get_human_detection_from_elastic

type Service interface {
	FindAll(id, date_time, person, file_name_capture_human_detection *string) ([]ElasticHumanDetection, error)
}

type service struct {
	elasticHumanDetectionRepository Repository
}

func NewService(elasticHumanDetectionRepository Repository) *service {
	return &service{elasticHumanDetectionRepository}
}

func (s *service) FindAll(id, date_time, person, file_name_capture_human_detection *string) ([]ElasticHumanDetection, error) {
	elasticHumanDetections, err := s.elasticHumanDetectionRepository.FindAll(id, date_time, person, file_name_capture_human_detection)
	if err != nil {
		return nil, err
	}

	return elasticHumanDetections, nil
}
