package streaming_cctv

type StreamingCctvInput struct {
	ID        int    `uri:"id" binding:"required"`
	IpAddress string `json:"ip_address"` // ini input dari service aja
}
