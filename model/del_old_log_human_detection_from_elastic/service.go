package del_old_log_human_detection_from_elastic

type Service interface {
	DelOneMonthOldHumanDetectionLogs() error
}

type service struct {
	delHumanDetectionFromElasticRepository Repository
}

func NewService(delHumanDetectionFromElasticRepository Repository) *service {
	return &service{delHumanDetectionFromElasticRepository}
}

func (s *service) DelOneMonthOldHumanDetectionLogs() error {
	err := s.delHumanDetectionFromElasticRepository.DelOneMonthOldHumanDetectionLogs()
	if err != nil {
		return err
	}

	return nil
}
