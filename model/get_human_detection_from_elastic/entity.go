package get_human_detection_from_elastic

type ElasticHumanDetection struct {
	ID                            string `json:"id"`
	Tid                           string `json:"tid"`
	DateTime                      string `json:"date_time"`
	Person                        string `json:"person"`
	FileNameCaptureHumanDetection string `json:"file_name_capture_human_detection"`
}
