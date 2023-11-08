package get_vandal_detection_from_elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"time"

	esv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Repository interface {
	FindAll(id string, tid string, date_time string, start_date string, end_date string, person string, file_name_capture_vandal_detection string) ([]ElasticVandalDetection, error)
}

type repository struct {
	elasticsearch *esv7.Client
}

func NewRepository(elasticsearch *esv7.Client) *repository {
	return &repository{elasticsearch}
}

func (r *repository) FindAll(id string, tid string, date_time string, start_date string, end_date string, person string, file_name_capture_vandal_detection string) ([]ElasticVandalDetection, error) {
	var (
		err   error
		query map[string]interface{}
		res   *esapi.Response
		rdb   map[string]interface{}
		hits  []interface{}
		edh   ElasticVandalDetection
		edhs  []ElasticVandalDetection
	)

	if r.elasticsearch == nil {
		return []ElasticVandalDetection{}, errors.New("elasticsearch not initialized")
	}

	if id != "" || tid != "" || date_time != "" || start_date != "" || end_date != "" || person != "" || file_name_capture_vandal_detection != "" {
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

	if tid != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"term": map[string]interface{}{
				"tid": tid,
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

	// range date time
	if start_date != "" && end_date != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"range": map[string]interface{}{
				"date_time.keyword": map[string]interface{}{
					"gte": start_date,
					"lte": end_date,
				},
			},
		})
	}

	if start_date != "" && end_date == "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"range": map[string]interface{}{
				"date_time.keyword": map[string]interface{}{
					"gte": start_date,
					"lte": time.Now().Format("2006-01-02T15:04:05"),
				},
			},
		})
	}

	if start_date == "" && end_date != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"range": map[string]interface{}{
				"date_time.keyword": map[string]interface{}{
					"gte": "2000-01-01T00:00:00",
					"lte": end_date,
				},
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
		Perhatikan penggunaan .keyword di bidang "file_name_capture_vandal_detection". Ini mengacu pada bidang yang tidak dianalisis dan memungkinkan pencocokan tepat dengan nilai yang diberikan.
	*/
	if file_name_capture_vandal_detection != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"term": map[string]interface{}{
				"file_name_capture_vandal_detection.keyword": file_name_capture_vandal_detection,
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
				"order": "desc",
			},
		},
	}

	// done=========================================================================================================================================
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return []ElasticVandalDetection{}, err
	}

	res, err = r.elasticsearch.Search(
		r.elasticsearch.Search.WithContext(context.Background()),
		r.elasticsearch.Search.WithIndex("vandal_detection_index"),
		r.elasticsearch.Search.WithBody(bytes.NewReader(queryBytes)),
		r.elasticsearch.Search.WithTrackTotalHits(true),
		r.elasticsearch.Search.WithPretty(),
	)

	if err != nil {
		return []ElasticVandalDetection{}, err
	}

	defer res.Body.Close()

	if res.IsError() {
		return []ElasticVandalDetection{}, err
	}

	if err := json.NewDecoder(res.Body).Decode(&rdb); err != nil {
		return []ElasticVandalDetection{}, err
	}

	hits, _ = rdb["hits"].(map[string]interface{})["hits"].([]interface{})

	for _, hit := range hits {
		edh.ID = hit.(map[string]interface{})["_id"].(string)

		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			continue // Skip this iteration if _source is not found in the hit
		}

		edh.Tid = source["tid"].(string)
		edh.DateTime = source["date_time"].(string)
		edh.Person = source["person"].(string)
		edh.FileNameCaptureVandalDetection = source["file_name_capture_vandal_detection"].(string)

		edhs = append(edhs, edh)
	}

	return edhs, nil
}
