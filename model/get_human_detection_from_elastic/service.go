package get_human_detection_from_elastic

type Service interface {
	FindAll(findAllElasticHumanDetectionInput FindAllElasticHumanDetectionInput) ([]ElasticHumanDetection, error)
}

type service struct {
	elasticHumanDetectionRepository Repository
}

func NewService(elasticHumanDetectionRepository Repository) *service {
	return &service{elasticHumanDetectionRepository}
}

func (s *service) FindAll(findAllElasticHumanDetectionInput FindAllElasticHumanDetectionInput) ([]ElasticHumanDetection, error) {

	elasticHumanDetections, err := s.elasticHumanDetectionRepository.FindAll(
		findAllElasticHumanDetectionInput.ID,
		findAllElasticHumanDetectionInput.TidID,
		findAllElasticHumanDetectionInput.DateTime,
		findAllElasticHumanDetectionInput.StartDate,
		findAllElasticHumanDetectionInput.EndDate,
		findAllElasticHumanDetectionInput.Person,
		findAllElasticHumanDetectionInput.FileNameCaptureHumanDetection,
	)
	if err != nil {
		return nil, err
	}

	return elasticHumanDetections, nil
}
