package get_status_mc_detection_from_elastic

type ElasticStatusMcDetection struct {
	ID            string `json:"id"`
	TidID         int    `json:"tid_id"`
	DateTime      string `json:"date_time"`
	StatusSignal  string `json:"status_signal"`
	StatusStorage string `json:"status_storage"`
	StatusRam     string `json:"status_ram"`
	StatusCpu     string `json:"status_cpu"`
}