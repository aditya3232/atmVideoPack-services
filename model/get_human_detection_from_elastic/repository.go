package get_human_detection_from_elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	esv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Repository interface {
	FindAll(id string, tid_id int, date_time string, person string, file_name_capture_human_detection string) ([]ElasticHumanDetection, error)
}

type repository struct {
	elasticsearch *esv7.Client
}

func NewRepository(elasticsearch *esv7.Client) *repository {
	return &repository{elasticsearch}
}

func (r *repository) FindAll(id string, tid_id int, date_time string, person string, file_name_capture_human_detection string) ([]ElasticHumanDetection, error) {
	var (
		err   error
		query map[string]interface{}
		res   *esapi.Response
		rdb   map[string]interface{}
		hits  []interface{}
		edh   ElasticHumanDetection
		edhs  []ElasticHumanDetection
	)

	if r.elasticsearch == nil {
		return []ElasticHumanDetection{}, errors.New("elasticsearch not initialized")
	}

	if id != "" || tid_id != 0 || date_time != "" || person != "" || file_name_capture_human_detection != "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{},
				},
			},
			"size": 100,
		}
	}

	if id != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"term": map[string]interface{}{
				"_id": id,
			},
		})
	}

	if tid_id != 0 {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"term": map[string]interface{}{
				"tid_id": tid_id,
			},
		})
	}

	if date_time != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"match": map[string]interface{}{
				"date_time.keyword": date_time,
			},
		})
	}

	if person != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"term": map[string]interface{}{
				"person": person,
			},
		})
	}

	/*
		Perhatikan penggunaan .keyword di bidang "file_name_capture_human_detection". Ini mengacu pada bidang yang tidak dianalisis dan memungkinkan pencocokan tepat dengan nilai yang diberikan.
	*/
	if file_name_capture_human_detection != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"term": map[string]interface{}{
				"file_name_capture_human_detection.keyword": file_name_capture_human_detection,
			},
		})
	}

	// Jika tidak ada kondisi pencarian yang diberikan, maka gunakan "match_all".
	if query == nil {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
			"size": 100,
		}
	}

	// Tambahkan sorting berdasarkan field date_time secara descending
	query["sort"] = []map[string]interface{}{
		{
			"date_time.keyword": map[string]interface{}{
				"order": "asc",
			},
		},
	}

	// done=========================================================================================================================================
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return []ElasticHumanDetection{}, err
	}

	res, err = r.elasticsearch.Search(
		r.elasticsearch.Search.WithContext(context.Background()),
		r.elasticsearch.Search.WithIndex("human_detection_index"),
		r.elasticsearch.Search.WithBody(bytes.NewReader(queryBytes)),
		r.elasticsearch.Search.WithTrackTotalHits(true),
		r.elasticsearch.Search.WithPretty(),
	)

	if err != nil {
		return []ElasticHumanDetection{}, err
	}

	defer res.Body.Close()

	if res.IsError() {
		return []ElasticHumanDetection{}, err
	}

	if err := json.NewDecoder(res.Body).Decode(&rdb); err != nil {
		return []ElasticHumanDetection{}, err
	}

	hits, _ = rdb["hits"].(map[string]interface{})["hits"].([]interface{})

	for _, hit := range hits {
		edh.ID = hit.(map[string]interface{})["_id"].(string)

		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			continue // Skip this iteration if _source is not found in the hit
		}

		tidID, ok := source["tid_id"]
		if ok {
			tidIDInt, err := strconv.Atoi(fmt.Sprintf("%v", tidID))
			if err != nil {
				// Handle the error if the conversion fails
				return []ElasticHumanDetection{}, err
			}
			edh.TidID = tidIDInt
		} else {
			edh.TidID = 0 // or set it to another default value if needed
		}

		edh.DateTime = source["date_time"].(string)
		edh.Person = source["person"].(string)
		edh.FileNameCaptureHumanDetection = source["file_name_capture_human_detection"].(string)

		edhs = append(edhs, edh)
	}

	return edhs, nil
}
