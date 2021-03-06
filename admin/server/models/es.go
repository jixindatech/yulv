package models

import (
	"admin/config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	es "github.com/elastic/go-elasticsearch/v7"
	"io"
)

var esClient *es.Client
var esCfg *config.Elasticsearch

func SetupES(cfg *config.Elasticsearch) error {
	esCfg = cfg
	es7Config := es.Config{
		Addresses: []string{esCfg.Host},
		Username:  esCfg.User,
		Password:  esCfg.Password,
		CloudID:   "",
		APIKey:    "",
	}

	var err error
	esClient, err = es.NewClient(es7Config)
	if err != nil {
		return err
	}

	res, err := esClient.Info()
	if err != nil {
		return fmt.Errorf("es error: %s\n", err)
	}
	defer res.Body.Close()

	return nil
}

func EsSearch(index string, query map[string]interface{}) (map[string]interface{}, error) {
	if index == "access" {
		index = esCfg.AccessIndex
	} else if index == "rule" {
		index = esCfg.RuleIndex
	} else {
		return nil, fmt.Errorf("%s", "incorrect es index")
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("Error encoding query: %s", err)
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)

	if err != nil {
		return nil, fmt.Errorf("Error getting response: %s", err)
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("Error parsing the response body: %s", err)
		} else {
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}

	return r, nil
}

func EsSearchList(index string, query map[string]interface{}, page, pageSize int) (map[string]interface{}, error) {
	if index == "access" {
		index = esCfg.AccessIndex
	} else if index == "rule" {
		index = esCfg.RuleIndex
	} else {
		return nil, fmt.Errorf("%s", "incorrect es index")
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, fmt.Errorf("Error encoding query: %s", err)
	}

	res, err := esClient.Search(
		esClient.Search.WithContext(context.Background()),
		esClient.Search.WithIndex(index),
		esClient.Search.WithBody(&buf),
		esClient.Search.WithFrom((page-1)*pageSize),
		esClient.Search.WithSize(pageSize),
		esClient.Search.WithTrackTotalHits(true),
		esClient.Search.WithPretty(),
	)

	if err != nil {
		return nil, fmt.Errorf("Error getting response: %s", err)
	}

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, fmt.Errorf("Error parsing the response body: %s", err)
		} else {
			return nil, fmt.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("Error parsing the response body: %s", err)
	}

	return r, nil
}
