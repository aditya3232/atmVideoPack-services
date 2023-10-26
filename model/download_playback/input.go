package download_playback

type ServiceDownloadPlaybackInput struct {
	Tid           string `form:"tid"`
	FolderDate    string `form:"folder_date"`
	StarthourDate string `form:"starthour_date"`
	Endhour       string `form:"endhour"`
	IpAddress     string `json:"ip_address"` // untuk balikan ip address from getOnByTid
}
