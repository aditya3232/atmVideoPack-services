package get_download_playback_from_elastic

type ElasticDownloadPlayback struct {
	ID              string `json:"id"`
	Tid             string `json:"tid"`
	DateModified    string `json:"date_modified"`
	DurationMinutes string `json:"duration_minutes"`
	FileSizeBytes   string `json:"file_size_bytes"`
	Filename        string `json:"filename"`
	Url             string `json:"url"`
}
