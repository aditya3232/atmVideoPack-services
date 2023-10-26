package download_playback

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/aditya3232/atmVideoPack-services.git/config"
	"gorm.io/gorm"
)

type Repository interface {
	DownloadPlayback(serviceDownloadPlaybackInput ServiceDownloadPlaybackInput) ([]ServiceDownloadPlayback, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) DownloadPlayback(serviceDownloadPlaybackInput ServiceDownloadPlaybackInput) ([]ServiceDownloadPlayback, error) {
	var response ServiceDownloadPlaybackResponse

	// Create the URL
	baseURL := fmt.Sprintf("http://%s:%s/%s", serviceDownloadPlaybackInput.IpAddress, config.CONFIG.SERVICE_DOWNLOAD_PLAYBACK_PORT, config.CONFIG.SERVICE_DOWNLOAD_PLAYBACK_HOST)

	// Encode the data as a query string
	data := url.Values{}
	data.Set("tid", serviceDownloadPlaybackInput.Tid)
	data.Set("folder_date", serviceDownloadPlaybackInput.FolderDate)
	data.Set("starthour_date", serviceDownloadPlaybackInput.StarthourDate)
	data.Set("endhour", serviceDownloadPlaybackInput.Endhour)
	body := strings.NewReader(data.Encode())

	// Create the HTTP request with the encoded data as the body
	req, err := http.NewRequest("GET", baseURL, body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req.Header.Add("x-api-key", config.CONFIG.SERVICE_DOWNLOAD_PLAYBACK_API_KEY)
	req.Header.Add("Content-Type", config.CONFIG.SERVICE_DOWNLOAD_PLAYBACK_CONTENT_TYPE)

	// Membuat klien HTTP
	client := &http.Client{}

	// Melakukan permintaan HTTP
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	// Decode the response JSON into your data structure
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	return response.Videos, nil
}
