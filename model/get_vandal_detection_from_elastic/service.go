package get_vandal_detection_from_elastic

type Service interface {
	FindAll(findAllElasticVandalDetectionInput FindAllElasticVandalDetectionInput) ([]ElasticVandalDetection, error)
}

type service struct {
	elasticVandalDetectionRepository Repository
}

func NewService(elasticVandalDetectionRepository Repository) *service {
	return &service{elasticVandalDetectionRepository}
}

func (s *service) FindAll(findAllElasticVandalDetectionInput FindAllElasticVandalDetectionInput) ([]ElasticVandalDetection, error) {

	elasticVandalDetections, err := s.elasticVandalDetectionRepository.FindAll(
		findAllElasticVandalDetectionInput.ID,
		findAllElasticVandalDetectionInput.Tid,
		findAllElasticVandalDetectionInput.DateTime,
		findAllElasticVandalDetectionInput.StartDate,
		findAllElasticVandalDetectionInput.EndDate,
		findAllElasticVandalDetectionInput.Person,
		findAllElasticVandalDetectionInput.FileNameCaptureVandalDetection,
	)
	if err != nil {
		return nil, err
	}

	return elasticVandalDetections, nil
}
