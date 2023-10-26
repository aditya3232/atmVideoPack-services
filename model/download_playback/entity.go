package download_playback

// sementara ini gk dipake karena balikan pakai map[string]interface{}

type ServiceDownloadPlayback struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

// untuk decode dari balikan servis mas yahya, harus masuk ke videos dlu soalnya
type ServiceDownloadPlaybackResponse struct {
	Videos []ServiceDownloadPlayback `json:"videos"`
}
