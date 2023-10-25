package del_old_log_status_mc_detection_from_elastic

type Service interface {
	DelOneMonthOldStatusMcDetectionLogs() error
}

type service struct {
	delStatusMcDetectionFromElasticRepository Repository
}

func NewService(delStatusMcDetectionFromElasticRepository Repository) *service {
	return &service{delStatusMcDetectionFromElasticRepository}
}

func (s *service) DelOneMonthOldStatusMcDetectionLogs() error {
	err := s.delStatusMcDetectionFromElasticRepository.DelOneMonthOldStatusMcDetectionLogs()
	if err != nil {
		return err
	}

	return nil
}
