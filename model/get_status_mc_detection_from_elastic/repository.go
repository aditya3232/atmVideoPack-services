package get_status_mc_detection_from_elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	esv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Repository interface {
	FindAll(id string, tid string, date_time string, start_date string, end_date string) ([]ElasticStatusMcDetection, error)
	FindDeviceUpDown() ([]ElasticStatusMcDetectionOnOrOff, error)
}

type repository struct {
	elasticsearch *esv7.Client
}

func NewRepository(elasticsearch *esv7.Client) *repository {
	return &repository{elasticsearch}
}

func (r *repository) FindAll(id string, tid string, date_time string, start_date string, end_date string) ([]ElasticStatusMcDetection, error) {
	var (
		err   error
		query map[string]interface{}
		res   *esapi.Response
		rdb   map[string]interface{}
		hits  []interface{}
		edh   ElasticStatusMcDetection
		edhs  []ElasticStatusMcDetection
	)

	if r.elasticsearch == nil {
		return []ElasticStatusMcDetection{}, errors.New("elasticsearch not initialized")
	}

	if id != "" || tid != "" || date_time != "" || start_date != "" || end_date != "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{},
				},
			},
			"size": 1,
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

	// Jika tidak ada kondisi pencarian yang diberikan, maka gunakan "match_all".
	if query == nil {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match_all": map[string]interface{}{},
			},
			"size": 1,
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
		return []ElasticStatusMcDetection{}, err
	}

	res, err = r.elasticsearch.Search(
		r.elasticsearch.Search.WithContext(context.Background()),
		r.elasticsearch.Search.WithIndex("status_mc_detection_index"),
		r.elasticsearch.Search.WithBody(bytes.NewReader(queryBytes)),
		r.elasticsearch.Search.WithTrackTotalHits(true),
		r.elasticsearch.Search.WithPretty(),
	)

	if err != nil {
		return []ElasticStatusMcDetection{}, err
	}

	defer res.Body.Close()

	// if res.IsError() {
	// 	return []ElasticStatusMcDetection{}, err
	// }

	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch error: %s", res.String())
	}

	if err := json.NewDecoder(res.Body).Decode(&rdb); err != nil {
		return []ElasticStatusMcDetection{}, err
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
		edh.StatusSignal = source["status_signal"].(string)
		edh.StatusStorage = source["status_storage"].(string)
		edh.StatusRam = source["status_ram"].(string)
		edh.StatusCpu = source["status_cpu"].(string)

		edhs = append(edhs, edh)
	}

	return edhs, nil
}

func (r *repository) FindDeviceUpDown() ([]ElasticStatusMcDetectionOnOrOff, error) {
	var (
		err   error
		query map[string]interface{}
		res   *esapi.Response
		rdb   map[string]interface{}
		hits  []interface{}
		edh   ElasticStatusMcDetectionOnOrOff
		edhs  []ElasticStatusMcDetectionOnOrOff
	)

	if r.elasticsearch == nil {
		return []ElasticStatusMcDetectionOnOrOff{}, errors.New("elasticsearch not initialized")
	}

	/*
		- ambil data unik setiap tid_id, lalu ambil satu data terbaru berdasarkan date_time
		- lanjut, kita beri kondisi di golang, jika melebihi 2 jam maka offline, dan jika kurang dari atau sama dengan 2 jam maka online
	*/

	/*
		{
		"size": 0,
		"aggs": {
			"unique_tid_ids": {
			"terms": {
				"field": "tid_id",
				"size": 10
			},
			"aggs": {
				"latest_date_time": {
				"top_hits": {
					"size": 1,
					"sort": [
					{
						"date_time": {
						"order": "desc"
						}
					}
					]
				}
				}
			}
			}
		}
		}
	*/

	query = map[string]interface{}{
		"size": 0,
		"aggs": map[string]interface{}{
			"unique_tids": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "tid.keyword",
					"size":  1000000,
				},
				"aggs": map[string]interface{}{
					"latest_date_time": map[string]interface{}{
						"top_hits": map[string]interface{}{
							"size": 1,
							"sort": []map[string]interface{}{
								{
									"date_time.keyword": map[string]interface{}{
										"order": "desc",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// proses menampilkan datanya
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return []ElasticStatusMcDetectionOnOrOff{}, err
	}

	res, err = r.elasticsearch.Search(
		r.elasticsearch.Search.WithContext(context.Background()),
		r.elasticsearch.Search.WithIndex("status_mc_detection_index"),
		r.elasticsearch.Search.WithBody(bytes.NewReader(queryBytes)),
		r.elasticsearch.Search.WithTrackTotalHits(true),
		r.elasticsearch.Search.WithPretty(),
	)

	if err != nil {
		return []ElasticStatusMcDetectionOnOrOff{}, err
	}

	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch error: %s", res.String())
	}

	if err := json.NewDecoder(res.Body).Decode(&rdb); err != nil {
		return []ElasticStatusMcDetectionOnOrOff{}, err
	}

	/*
			{
			"took": 10,
			"timed_out": false,
			"_shards": {
				"total": 1,
				"successful": 1,
				"skipped": 0,
				"failed": 0
			},
			"hits": {
				"total": {
					"value": 57,
					"relation": "eq"
				},
				"max_score": null,
				"hits": []
			},
			"aggregations": {
				"unique_tid_ids": {
					"doc_count_error_upper_bound": 0,
					"sum_other_doc_count": 0,
					"buckets": [
						{
							"key": "1",
							"doc_count": 54,
							"latest_date_time": {
								"hits": {
									"total": {
										"value": 54,
										"relation": "eq"
									},
									"max_score": null,
									"hits": [
										{
											"_index": "status_mc_detection_index",
											"_type": "_doc",
											"_id": "Wk28k4sB_mgHndj7i4Rl",
											"_score": null,
											"_source": {
												"id": "2023-11-03-12-51-29-122",
												"tid_id": "1",
												"date_time": "2023-11-03 12:51:28",
												"status_signal": "Download Speed: 3.46 Mbps, Upload Speed: 6.45 Mbps",
												"status_storage": "Used 37.96 GB , Not Used 404.05 GB",
												"status_ram": "2.85 GB, 17.90 %",
												"status_cpu": "59.60%"
											},
											"sort": [
												"2023-11-03 12:51:28"
											]
										}
									]
								}
							}
						},
						{
							"key": "2",
							"doc_count": 3,
							"latest_date_time": {
								"hits": {
									"total": {
										"value": 3,
										"relation": "eq"
									},
									"max_score": null,
									"hits": [
										{
											"_index": "status_mc_detection_index",
											"_type": "_doc",
											"_id": "5E2VkIsB_mgHndj7loN8",
											"_score": null,
											"_source": {
												"id": "2023-11-02-22-10-04-334",
												"tid_id": "2",
												"date_time": "2023-10-09 05:12:39",
												"status_signal": "good",
												"status_storage": "good",
												"status_ram": "good",
												"status_cpu": "good"
											},
											"sort": [
												"2023-10-09 05:12:39"
											]
										}
									]
								}
							}
						}
					]
				}
			}
		}

	*/

	hits, _ = rdb["aggregations"].(map[string]interface{})["unique_tids"].(map[string]interface{})["buckets"].([]interface{})
	for _, hit := range hits {
		edh.ID = hit.(map[string]interface{})["latest_date_time"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_id"].(string)
		edh.Tid = hit.(map[string]interface{})["latest_date_time"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["tid"].(string)
		edh.DateTime = hit.(map[string]interface{})["latest_date_time"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["date_time"].(string)
		edh.StatusSignal = hit.(map[string]interface{})["latest_date_time"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["status_signal"].(string)
		edh.StatusStorage = hit.(map[string]interface{})["latest_date_time"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["status_storage"].(string)
		edh.StatusRam = hit.(map[string]interface{})["latest_date_time"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["status_ram"].(string)
		edh.StatusCpu = hit.(map[string]interface{})["latest_date_time"].(map[string]interface{})["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})["status_cpu"].(string)

		// if date_time > 2 hours then edh.StatusMc = "offline", if date_time <= 2 hours then edh.StatusMc = "online"
		date_time, _ := time.Parse("2006-01-02 15:04:05", edh.DateTime)
		if time.Since(date_time).Hours() > 2 {
			edh.StatusMc = "offline"
		} else {
			edh.StatusMc = "online"
		}

		edhs = append(edhs, edh)

	}

	return edhs, nil

}
