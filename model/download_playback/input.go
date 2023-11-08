package download_playback

type ServiceDownloadPlaybackInput struct {
	Videos    string `uri:"videos" binding:"required"`
	Tid       string `uri:"tid" binding:"required"`
	Yyyymmdd  string `uri:"yyyymmdd" binding:"required"`
	Filename  string `uri:"filename" binding:"required"`
	IpAddress string `json:"ip_address"` // untuk menampung input dari service
}
