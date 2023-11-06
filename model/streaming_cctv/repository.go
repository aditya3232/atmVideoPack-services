package streaming_cctv

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/aditya3232/atmVideoPack-services.git/config"
	"gorm.io/gorm"
)

type Repository interface {
	StreamingCctv(input StreamingCctvInput) (*bytes.Buffer, error)
}

// db disini hanya buat syarat aja
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) StreamingCctv(input StreamingCctvInput) (*bytes.Buffer, error) {
	thirdPartyURL := fmt.Sprintf("http://%s:%s", input.IpAddress, config.CONFIG.SERVICE_STREAMING_CCTV_PORT)

	// Ambil video dari URL pihak ketiga
	thirdPartyResp, err := http.Get(thirdPartyURL)
	if err != nil {
		return nil, err
	}

	defer thirdPartyResp.Body.Close()

	// Buat buffer untuk menyimpan video streaming
	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, thirdPartyResp.Body)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
