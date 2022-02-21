package service

import "admin/server/models"

type AccessEvent struct {
	ID uint

	Start int64
	End   int64

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
