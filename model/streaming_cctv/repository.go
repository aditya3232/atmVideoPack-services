package streaming_cctv

import (
	"fmt"
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/config"
	"gorm.io/gorm"
)

type Repository interface {
	StreamingCctv(input StreamingCctvInput) (*http.Response, error)
}

// db disini hanya buat syarat aja
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) StreamingCctv(input StreamingCctvInput) (*http.Response, error) {
	streamingURL := fmt.Sprintf("http://%s:%s", input.IpAddress, config.CONFIG.SERVICE_STREAMING_CCTV_PORT)

	// Ambil video dari URL pihak ketiga
	response, err := http.Get(streamingURL)
	if err != nil {
		return nil, err
	}

	return response, nil
}
