package get_human_detection_from_elastic

type ElasticHumanDetectionGetFormatter struct {
	ID                            string `json:"id"`
	TidID                         *int   `json:"tid_id"`
	DateTime                      string `json:"date_time"`
	Person                        string `json:"person"`
	FileNameCaptureHumanDetection string `json:"file_name_capture_human_detection"`
}

func ElasticHumanDetectionGetFormat(elasticHumanDetection ElasticHumanDetection) ElasticHumanDetectionGetFormatter {
	var formatter ElasticHumanDetectionGetFormatter

	formatter.ID = elasticHumanDetection.ID
	formatter.TidID = elasticHumanDetection.TidID
	formatter.DateTime = elasticHumanDetection.DateTime
	formatter.Person = elasticHumanDetection.Person
	formatter.FileNameCaptureHumanDetection = elasticHumanDetection.FileNameCaptureHumanDetection

	return formatter
}

func ElasticHumanDetectionGetAllFormat(elasticHumanDetections []ElasticHumanDetection) []ElasticHumanDetectionGetFormatter {
	elasticHumanDetectionsFormatter := []ElasticHumanDetectionGetFormatter{}

	for _, elasticHumanDetection := range elasticHumanDetections {
		elasticHumanDetectionFormatter := ElasticHumanDetectionGetFormat(elasticHumanDetection)                   // format data satu persatu
		elasticHumanDetectionsFormatter = append(elasticHumanDetectionsFormatter, elasticHumanDetectionFormatter) // append data formatter ke slice formatter
	}

	return elasticHumanDetectionsFormatter
}
