package get_human_detection_from_elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	esv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Repository interface {
	FindAll(id, date_time, person, file_name_capture_human_detection *string) ([]ElasticHumanDetection, error)
}

type repository struct {
	elasticsearch *esv7.Client
}

func NewRepository(elasticsearch *esv7.Client) *repository {
	return &repository{elasticsearch}
}

func (r *repository) FindAll(id, date_time, person, file_name_capture_human_detection *string) ([]ElasticHumanDetection, error) {
	var (
		err   error
		query map[string]interface{}
		res   *esapi.Response
		rdb   map[string]interface{}
		hits  []interface{}
		edh   ElasticHumanDetection
	)

	if r.elasticsearch == nil {
		return []ElasticHumanDetection{}, errors.New("elasticsearch not initialized")
	}

	// if file_name_capture_human_detection not null search with wild card
	if file_name_capture_human_detection != nil {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"wildcard": map[string]interface{}{
					"file_name_capture_human_detection": map[string]interface{}{
						"value": fmt.Sprintf("*%s*", *file_name_capture_human_detection),
					},
				},
			},
			"size": 1000,
		}
	}

	// jika tidak ada parameter search all
	query = map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"size": 1000,
	}

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

	var (
		edhs []ElasticHumanDetection
	)

	for _, hit := range hits {

		edh.ID = hit.(map[string]interface{})["_id"].(string)
		tidID, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})["tid_id"].(float64)
		if ok {
			tidIDInt := int(tidID)
			edh.TidID = &tidIDInt
		}
		edh.DateTime = hit.(map[string]interface{})["_source"].(map[string]interface{})["date_time"].(string)
		edh.Person = hit.(map[string]interface{})["_source"].(map[string]interface{})["person"].(string)
		edh.FileNameCaptureHumanDetection = hit.(map[string]interface{})["_source"].(map[string]interface{})["file_name_capture_human_detection"].(string)

		edhs = append(edhs, edh)
	}

	return edhs, nil
}
