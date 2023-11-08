package streaming_cctv

type StreamingCctvInput struct {
	Tid       string `uri:"tid" binding:"required"`
	IpAddress string `json:"ip_address"` // ini input dari service aja
}
