package del_old_log_vandal_detection_from_elastic

type Service interface {
	DelOneMonthOldVandalDetectionLogs() error
}

type service struct {
	delVandalDetectionFromElasticRepository Repository
}

func NewService(delVandalDetectionFromElasticRepository Repository) *service {
	return &service{delVandalDetectionFromElasticRepository}
}

func (s *service) DelOneMonthOldVandalDetectionLogs() error {
	err := s.delVandalDetectionFromElasticRepository.DelOneMonthOldVandalDetectionLogs()
	if err != nil {
		return err
	}

	return nil
}
