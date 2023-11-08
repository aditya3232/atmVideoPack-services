package download_playback

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/config"
	"gorm.io/gorm"
)

type Repository interface {
	DownloadPlayback(input ServiceDownloadPlaybackInput) (*bytes.Buffer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) DownloadPlayback(input ServiceDownloadPlaybackInput) (*bytes.Buffer, error) {
	thirdPartyURL := fmt.Sprintf("http://%s:%s/%s/%s/%s/%s", input.IpAddress, config.CONFIG.SERVICE_DOWNLOAD_PLAYBACK_PORT, input.Videos, input.Tid, input.Yyyymmdd, input.Filename)

	// Ambil video dari URL pihak ketiga
	thirdPartyResp, err := http.Get(thirdPartyURL)
	if err != nil {
		return nil, err
	}

	defer thirdPartyResp.Body.Close()

	// Buat buffer untuk menyimpan file download
	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, thirdPartyResp.Body)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
