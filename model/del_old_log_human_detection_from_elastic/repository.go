package del_old_log_human_detection_from_elastic

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	log_function "github.com/aditya3232/atmVideoPack-services.git/log"
	esv7 "github.com/elastic/go-elasticsearch/v7"
)

type Repository interface {
	DelOneMonthOldHumanDetectionLogs() error
}

type repository struct {
	elasticsearch *esv7.Client
}

func NewRepository(elasticsearch *esv7.Client) *repository {
	return &repository{elasticsearch}
}

func (r *repository) DelOneMonthOldHumanDetectionLogs() error {
	oneMonthAgo := time.Now().AddDate(0, -1, 0)
	oneMonthAgoStr := oneMonthAgo.Format("15:04:05 02-01-2006")

	for {
		// Prepare the Elasticsearch query as a map
		query := map[string]interface{}{
			"query": map[string]interface{}{
				"range": map[string]interface{}{
					"timestamp": map[string]interface{}{
						"lte": oneMonthAgoStr,
					},
				},
			},
		}

		// Convert the query map to JSON
		queryJSON, err := json.Marshal(query)
		if err != nil {
			log.Println("Error marshaling JSON:", err)
			continue
		}

		// Delete documents using DeleteByQuery
		_, err = r.elasticsearch.DeleteByQuery([]string{"atm_video_pack_log"}, strings.NewReader(string(queryJSON)))
		if err != nil {
			return err
		}

		log_function.Info("delete old human detection from elastic success")
		// sleep for 10 minute
		time.Sleep(10 * time.Minute)

	}
}