package service

type bucketIntItem struct {
	DocCount int64 `json:"doc_count"`
	Key      int64 `json:"key"`
}

type bucketFloatItem struct {
	DocCount int     `json:"doc_count"`
	Key      float64 `json:"key"`
}

type bucketStringItem struct {
	DocCount int    `json:"doc_count"`
	Key      string `json:"key"`
}

type bucketIntervalData struct {
	Buckets []bucketFloatItem `json:"buckets"`
}
type bucketIntervalItem struct {
	Key      int                `json:"key"`
	DocCount int                `json:"doc_count"`
	Interval bucketIntervalData `json:"interval_data"`
}

func getKeyCount(items []interface{}) ([]string, []int) {
	var keys []string
	var count []int
	for _, item := range items {
		data := item.(map[string]interface{})
		keys = append(keys, data["key"].(string))
		count = append(count, int(data["doc_count"].(float64)))
	}

	return keys, count
}
