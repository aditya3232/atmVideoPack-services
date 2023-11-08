package get_status_mc_detection_from_elastic

type FindAllElasticStatusMcDetectionInput struct {
	ID        string `form:"id"`
	Tid       string `form:"tid"`
	DateTime  string `form:"date_time"`
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
}
