package get_human_detection_from_elastic

type FindAllElasticHumanDetectionInput struct {
	ID                            string `form:"id"`
	TidID                         int    `form:"tid_id"`
	DateTime                      string `form:"date_time"`
	StartDate                     string `form:"start_date"`
	EndDate                       string `form:"end_date"`
	Person                        string `form:"person"`
	FileNameCaptureHumanDetection string `form:"file_name_capture_human_detection"`
	// FileNameCaptureHumanDetection string `form:"file_name_capture_human_detection" binding:"required"`
}
