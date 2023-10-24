package get_vandal_detection_from_elastic

type FindAllElasticVandalDetectionInput struct {
	ID                             string `form:"id"`
	TidID                          int    `form:"tid_id"`
	DateTime                       string `form:"date_time"`
	StartDate                      string `form:"start_date"`
	EndDate                        string `form:"end_date"`
	Person                         string `form:"person"`
	FileNameCaptureVandalDetection string `form:"file_name_capture_vandal_detection"`
	// FileNameCaptureVandalDetection string `form:"file_name_capture_vandal_detection" binding:"required"`
}
