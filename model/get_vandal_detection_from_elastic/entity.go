package get_vandal_detection_from_elastic

type ElasticVandalDetection struct {
	ID                             string `json:"id"`
	Tid                            string `json:"tid"`
	DateTime                       string `json:"date_time"`
	Person                         string `json:"person"`
	FileNameCaptureVandalDetection string `json:"file_name_capture_vandal_detection"`
}
