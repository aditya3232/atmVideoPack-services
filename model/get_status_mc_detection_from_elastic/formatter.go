package get_status_mc_detection_from_elastic

type ElasticStatusMcDetectionGetFormatter struct {
	ID            string `json:"id"`
	TidID         int    `json:"tid_id"`
	DateTime      string `json:"date_time"`
	StatusSignal  string `json:"status_signal"`
	StatusStorage string `json:"status_storage"`
	StatusRam     string `json:"status_ram"`
	StatusCpu     string `json:"status_cpu"`
}

func ElasticStatusMcDetectionGetFormat(elasticStatusMcDetection ElasticStatusMcDetection) ElasticStatusMcDetectionGetFormatter {
	var formatter ElasticStatusMcDetectionGetFormatter

	formatter.ID = elasticStatusMcDetection.ID
	formatter.TidID = elasticStatusMcDetection.TidID
	formatter.DateTime = elasticStatusMcDetection.DateTime
	formatter.StatusSignal = elasticStatusMcDetection.StatusSignal
	formatter.StatusStorage = elasticStatusMcDetection.StatusStorage
	formatter.StatusRam = elasticStatusMcDetection.StatusRam
	formatter.StatusCpu = elasticStatusMcDetection.StatusCpu

	return formatter
}

func ElasticStatusMcDetectionGetAllFormat(elasticStatusMcDetections []ElasticStatusMcDetection) []ElasticStatusMcDetectionGetFormatter {
	elasticStatusMcDetectionsFormatter := []ElasticStatusMcDetectionGetFormatter{}

	for _, elasticStatusMcDetection := range elasticStatusMcDetections {
		elasticStatusMcDetectionFormatter := ElasticStatusMcDetectionGetFormat(elasticStatusMcDetection)                   // format data satu persatu
		elasticStatusMcDetectionsFormatter = append(elasticStatusMcDetectionsFormatter, elasticStatusMcDetectionFormatter) // append data formatter ke slice formatter
	}

	return elasticStatusMcDetectionsFormatter
}

type ElasticStatusMcDetectionOnOrOffFormatter struct {
	ID            string `json:"id"`
	TidID         int    `json:"tid_id"`
	DateTime      string `json:"date_time"`
	StatusSignal  string `json:"status_signal"`
	StatusStorage string `json:"status_storage"`
	StatusRam     string `json:"status_ram"`
	StatusCpu     string `json:"status_cpu"`
	StatusMc      string `json:"status_mc"`
}

func ElasticStatusMcDetectionOnOrOffGetFormat(elasticStatusMcDetectionOnOrOff ElasticStatusMcDetectionOnOrOff) ElasticStatusMcDetectionOnOrOffFormatter {
	var formatter ElasticStatusMcDetectionOnOrOffFormatter

	formatter.ID = elasticStatusMcDetectionOnOrOff.ID
	formatter.TidID = elasticStatusMcDetectionOnOrOff.TidID
	formatter.DateTime = elasticStatusMcDetectionOnOrOff.DateTime
	formatter.StatusSignal = elasticStatusMcDetectionOnOrOff.StatusSignal
	formatter.StatusStorage = elasticStatusMcDetectionOnOrOff.StatusStorage
	formatter.StatusRam = elasticStatusMcDetectionOnOrOff.StatusRam
	formatter.StatusCpu = elasticStatusMcDetectionOnOrOff.StatusCpu
	formatter.StatusMc = elasticStatusMcDetectionOnOrOff.StatusMc

	return formatter
}

func ElasticStatusMcDetectionOnOrOffGetAllFormat(elasticStatusMcDetectionOnOrOffs []ElasticStatusMcDetectionOnOrOff) []ElasticStatusMcDetectionOnOrOffFormatter {
	elasticStatusMcDetectionOnOrOffsFormatter := []ElasticStatusMcDetectionOnOrOffFormatter{}

	for _, elasticStatusMcDetectionOnOrOff := range elasticStatusMcDetectionOnOrOffs {
		elasticStatusMcDetectionOnOrOffFormatter := ElasticStatusMcDetectionOnOrOffGetFormat(elasticStatusMcDetectionOnOrOff)                   // format data satu persatu
		elasticStatusMcDetectionOnOrOffsFormatter = append(elasticStatusMcDetectionOnOrOffsFormatter, elasticStatusMcDetectionOnOrOffFormatter) // append data formatter ke slice formatter
	}

	return elasticStatusMcDetectionOnOrOffsFormatter
}
