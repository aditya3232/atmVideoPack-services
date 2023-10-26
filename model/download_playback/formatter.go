package download_playback

type ServiceDownloadPlaybackFormatter struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

func ServiceDownloadPlaybackFormat(serviceDownloadPlayback ServiceDownloadPlayback) ServiceDownloadPlaybackFormatter {
	var formatter ServiceDownloadPlaybackFormatter

	formatter.Filename = serviceDownloadPlayback.Filename
	formatter.Url = serviceDownloadPlayback.Url

	return formatter
}

func ServiceDownloadPlaybackFormatMany(serviceDownloadPlaybacks []ServiceDownloadPlayback) []ServiceDownloadPlaybackFormatter {
	serviceDownloadPlaybacksFormatter := []ServiceDownloadPlaybackFormatter{}

	if len(serviceDownloadPlaybacks) == 0 {
		return serviceDownloadPlaybacksFormatter
	}

	for _, serviceDownloadPlayback := range serviceDownloadPlaybacks {
		serviceDownloadPlaybackFormatter := ServiceDownloadPlaybackFormat(serviceDownloadPlayback)
		serviceDownloadPlaybacksFormatter = append(serviceDownloadPlaybacksFormatter, serviceDownloadPlaybackFormatter)
	}

	return serviceDownloadPlaybacksFormatter
}
