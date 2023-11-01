package get_download_playback_from_elastic

type ElasticDownloadPlaybackGetFormatter struct {
	ID              string `json:"id"`
	Tid             string `json:"tid"`
	DateModified    string `json:"date_modified"`
	DurationMinutes string `json:"duration_minutes"`
	FileSizeBytes   string `json:"file_size_bytes"`
	Filename        string `json:"filename"`
	Url             string `json:"url"`
}

func ElasticDownloadPlaybackGetFormat(elasticDownloadPlayback ElasticDownloadPlayback) ElasticDownloadPlaybackGetFormatter {
	var formatter ElasticDownloadPlaybackGetFormatter

	formatter.ID = elasticDownloadPlayback.ID
	formatter.Tid = elasticDownloadPlayback.Tid
	formatter.DateModified = elasticDownloadPlayback.DateModified
	formatter.DurationMinutes = elasticDownloadPlayback.DurationMinutes
	formatter.FileSizeBytes = elasticDownloadPlayback.FileSizeBytes
	formatter.Filename = elasticDownloadPlayback.Filename
	formatter.Url = elasticDownloadPlayback.Url

	return formatter
}

func ElasticDownloadPlaybackGetAllFormat(elasticDownloadPlaybacks []ElasticDownloadPlayback) []ElasticDownloadPlaybackGetFormatter {
	elasticDownloadPlaybacksFormatter := []ElasticDownloadPlaybackGetFormatter{}

	for _, elasticDownloadPlayback := range elasticDownloadPlaybacks {
		elasticDownloadPlaybackFormatter := ElasticDownloadPlaybackGetFormat(elasticDownloadPlayback)                   // format data satu persatu
		elasticDownloadPlaybacksFormatter = append(elasticDownloadPlaybacksFormatter, elasticDownloadPlaybackFormatter) // append data formatter ke slice formatter
	}

	return elasticDownloadPlaybacksFormatter
}
