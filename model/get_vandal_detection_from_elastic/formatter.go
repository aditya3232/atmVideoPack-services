package get_vandal_detection_from_elastic

type ElasticVandalDetectionGetFormatter struct {
	ID                             string `json:"id"`
	TidID                          int    `json:"tid_id"`
	DateTime                       string `json:"date_time"`
	Person                         string `json:"person"`
	FileNameCaptureVandalDetection string `json:"file_name_capture_vandal_detection"`
}

func ElasticVandalDetectionGetFormat(elasticVandalDetection ElasticVandalDetection) ElasticVandalDetectionGetFormatter {
	var formatter ElasticVandalDetectionGetFormatter

	formatter.ID = elasticVandalDetection.ID
	formatter.TidID = elasticVandalDetection.TidID
	formatter.DateTime = elasticVandalDetection.DateTime
	formatter.Person = elasticVandalDetection.Person
	formatter.FileNameCaptureVandalDetection = elasticVandalDetection.FileNameCaptureVandalDetection

	return formatter
}

func ElasticVandalDetectionGetAllFormat(elasticVandalDetections []ElasticVandalDetection) []ElasticVandalDetectionGetFormatter {
	elasticVandalDetectionsFormatter := []ElasticVandalDetectionGetFormatter{}

	for _, elasticVandalDetection := range elasticVandalDetections {
		elasticVandalDetectionFormatter := ElasticVandalDetectionGetFormat(elasticVandalDetection)                   // format data satu persatu
		elasticVandalDetectionsFormatter = append(elasticVandalDetectionsFormatter, elasticVandalDetectionFormatter) // append data formatter ke slice formatter
	}

	return elasticVandalDetectionsFormatter
}
