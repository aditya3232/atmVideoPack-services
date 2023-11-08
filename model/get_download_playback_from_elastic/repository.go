package get_download_playback_from_elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/aditya3232/atmVideoPack-services.git/config"
	esv7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Repository interface {
	FindAll(tid, date_time, start_date, end_date string) ([]ElasticDownloadPlayback, error)
}

type repository struct {
	elasticsearch *esv7.Client
}

func NewRepository(elasticsearch *esv7.Client) *repository {
	return &repository{elasticsearch}
}

func (r *repository) FindAll(tid, date_time, start_date, end_date string) ([]ElasticDownloadPlayback, error) {
	var (
		err   error
		query map[string]interface{}
		res   *esapi.Response
		rdb   map[string]interface{}
		hits  []interface{}
		edh   ElasticDownloadPlayback
		edhs  []ElasticDownloadPlayback
	)

	if r.elasticsearch == nil {
		return []ElasticDownloadPlayback{}, errors.New("elasticsearch not initialized")
	}

	if tid != "" || date_time != "" || start_date != "" || end_date != "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{},
				},
			},
			"size": 100,
		}
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
				"date_modified.keyword": date_time,
			},
		})
	}

	// range date time
	if start_date != "" && end_date != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"range": map[string]interface{}{
				"date_modified.keyword": map[string]interface{}{
					"gte": start_date,
					"lte": end_date,
				},
			},
		})
	}

	if start_date != "" && end_date == "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"range": map[string]interface{}{
				"date_modified.keyword": map[string]interface{}{
					"gte": start_date,
					"lte": time.Now().Format("2006-01-02T15:04:05"),
				},
			},
		})
	}

	if start_date == "" && end_date != "" {
		query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]map[string]interface{}), map[string]interface{}{
			"range": map[string]interface{}{
				"date_modified.keyword": map[string]interface{}{
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
			"size": 100,
		}
	}

	// Tambahkan sorting berdasarkan field date_time secara descending
	query["sort"] = []map[string]interface{}{
		{
			"date_modified.keyword": map[string]interface{}{
				"order": "desc",
			},
		},
	}

	// done=========================================================================================================================================
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return []ElasticDownloadPlayback{}, err
	}

	res, err = r.elasticsearch.Search(
		r.elasticsearch.Search.WithContext(context.Background()),
		r.elasticsearch.Search.WithIndex("download_playback_index"),
		r.elasticsearch.Search.WithBody(bytes.NewReader(queryBytes)),
		r.elasticsearch.Search.WithTrackTotalHits(true),
		r.elasticsearch.Search.WithPretty(),
	)

	if err != nil {
		return []ElasticDownloadPlayback{}, err
	}

	defer res.Body.Close()

	if res.IsError() {
		return []ElasticDownloadPlayback{}, err
	}

	if err := json.NewDecoder(res.Body).Decode(&rdb); err != nil {
		return []ElasticDownloadPlayback{}, err
	}

	hits, _ = rdb["hits"].(map[string]interface{})["hits"].([]interface{})

	for _, hit := range hits {
		edh.ID = hit.(map[string]interface{})["_id"].(string)

		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			continue // Skip this iteration if _source is not found in the hit
		}
		edh.Tid = source["tid"].(string)
		edh.DateModified = source["date_modified"].(string)
		edh.DurationMinutes = source["duration_minutes"].(string)
		edh.FileSizeBytes = source["file_size_bytes"].(string)
		edh.Filename = source["filename"].(string)
		// edh.Url = source["url"].(string)
		// encrypt url for save link
		// edh.Url = helper.Encrypt(edh.Url)

		/*
			- link asli => http://10.8.0.2:5001/videos/160001/2023-10-30/160001_20231030000107.mp4
			- ubah ke link be, nanti link be itu akan buffer ke link asli
			- berikut adalah link be, http://127.0.0.1:3636/api/atmvideopack/v1/downloadvideoplayback/videos/160001/2023-10-30/160001_20231030000107.mp4
			- nanti fe akan mendapatkan link be
		*/

		Yyyymmdd := source["date_modified"].(string)[:10]

		BeDownloadPlaybackURL := fmt.Sprintf("http://%s:%s/api/atmvideopack/v1/downloadvideoplayback/videos/%s/%s/%s", config.CONFIG.APP_HOST, config.CONFIG.APP_PORT, source["tid"].(string), Yyyymmdd, source["filename"].(string))

		edh.Url = BeDownloadPlaybackURL

		edhs = append(edhs, edh)
	}

	return edhs, nil
}
