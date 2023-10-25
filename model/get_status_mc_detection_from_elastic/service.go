package get_status_mc_detection_from_elastic

type Service interface {
	FindAll(findAllElasticStatusMcDetectionInput FindAllElasticStatusMcDetectionInput) ([]ElasticStatusMcDetection, error)
}

type service struct {
	elasticStatusMcDetectionRepository Repository
}

func NewService(elasticStatusMcDetectionRepository Repository) *service {
	return &service{elasticStatusMcDetectionRepository}
}

func (s *service) FindAll(findAllElasticStatusMcDetectionInput FindAllElasticStatusMcDetectionInput) ([]ElasticStatusMcDetection, error) {

	elasticStatusMcDetections, err := s.elasticStatusMcDetectionRepository.FindAll(
		findAllElasticStatusMcDetectionInput.ID,
		findAllElasticStatusMcDetectionInput.TidID,
		findAllElasticStatusMcDetectionInput.DateTime,
		findAllElasticStatusMcDetectionInput.StartDate,
		findAllElasticStatusMcDetectionInput.EndDate,
	)
	if err != nil {
		return nil, err
	}

	return elasticStatusMcDetections, nil
}
