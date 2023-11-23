package get_status_mc_detection_from_elastic

type Service interface {
	FindAll(findAllElasticStatusMcDetectionInput FindAllElasticStatusMcDetectionInput) ([]ElasticStatusMcDetection, error)
	FindDeviceUpDown() ([]ElasticStatusMcDetectionOnOrOff, error)
}

type service struct {
	elasticStatusMcDetectionRepository Repository
}

func NewService(elasticStatusMcDetectionRepository Repository) *service {
	return &service{elasticStatusMcDetectionRepository}
}

func (s *service) FindAll(findAllElasticStatusMcDetectionInput FindAllElasticStatusMcDetectionInput) ([]ElasticStatusMcDetection, error) {

	elasticStatusMcDetections, err := s.elasticStatusMcDetectionRepository.FindAll(findAllElasticStatusMcDetectionInput)
	if err != nil {
		return nil, err
	}

	return elasticStatusMcDetections, nil
}

func (s *service) FindDeviceUpDown() ([]ElasticStatusMcDetectionOnOrOff, error) {

	elasticStatusMcDetectionOnOrOffs, err := s.elasticStatusMcDetectionRepository.FindDeviceUpDown()
	if err != nil {
		return nil, err
	}

	return elasticStatusMcDetectionOnOrOffs, nil
}
