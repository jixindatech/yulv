package service

import (
	"admin/server/models"
	"math"
)

type AccessEvent struct {
	ID uint

	Database string
	Start    int64
	End      int64

	Page     int
	PageSize int
}

func (a *AccessEvent) GetList() (map[string]interface{}, error) {
	start := a.Start / 1000
	end := a.End / 1000
	filter := []map[string]interface{}{
		{
			"range": map[string]interface{}{
				"timestamp": map[string]interface{}{
					"gte": start,
					"lte": end,
				},
			},
		},
	}
	query := map[string]interface{}{
		"sort": map[string]interface{}{
			"timestamp": map[string]interface{}{
				"order": "desc",
			},
		},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": filter,
			},
		},
	}

	data, err := models.EsSearchList("access", query, a.Page, a.PageSize)
	if err != nil {
		return nil, err
	}

	res := make(map[string]interface{})
	res["count"] = data["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"]
	res["data"] = data["hits"].(map[string]interface{})["hits"]

	return res, nil
}

func (a *AccessEvent) GetInfo() (map[string]interface{}, error) {
	start := a.Start / 1000
	end := a.End / 1000

	var query map[string]interface{}
	if len(a.Database) > 0 {
		query = map[string]interface{}{
			"size": 0,
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{
							"term": map[string]interface{}{
								"database": a.Database,
							},
						},
						{
							"range": map[string]interface{}{
								"timestamp": map[string]interface{}{
									"gte": start,
									"lte": end,
								},
							},
						},
					},
				},
			},
			"aggs": map[string]interface{}{
				"group_by_fingerprint": map[string]interface{}{
					"terms": map[string]interface{}{
						"field": "fingerprint",
						"size":  math.MaxInt32,
					},
				},
			},
		}
	} else {
		query = map[string]interface{}{
			"size": 0,
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{
							"range": map[string]interface{}{
								"timestamp": map[string]interface{}{
									"gte": start,
									"lte": end,
								},
							},
						},
					},
				},
			},
			"aggs": map[string]interface{}{
				"group_by_fingerprint": map[string]interface{}{
					"terms": map[string]interface{}{
						"field": "fingerprint",
						"size":  math.MaxInt32,
					},
				},
			},
		}
	}

	data, err := models.EsSearch("access", query)
	if err != nil {
		return nil, err
	}

	buckets := data["aggregations"].(map[string]interface{})["group_by_fingerprint"].(map[string]interface{})["buckets"]
	keys, count := getKeyCount(buckets.([]interface{}))
	if len(keys) == 0 {
		keys = []string{}
	}
	if len(count) == 0 {
		count = []int{}
	}

	return map[string]interface{}{
		"item": keys,
		"num":  count,
	}, nil
}
