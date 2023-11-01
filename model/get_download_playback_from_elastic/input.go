package get_download_playback_from_elastic

type FindAllElasticDownloadPlaybackInput struct {
	Tid       string `form:"tid"`
	DateTime  string `form:"date_time"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}
