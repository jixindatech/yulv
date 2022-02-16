package service

import (
	"admin/server/cache"
	"admin/server/models"
	"encoding/json"
	"time"
)

type IP struct {
	ID uint

	Name   string
	Type   int
	IP     []byte
	Remark string

	Page     int
	PageSize int
}

func (d *IP) Get() (*models.IP, error) {
	dbItem, err := models.GetIP(d.ID)
	if err != nil {
		return nil, err
	}
	return dbItem, nil
}

func (d *IP) GetList() ([]*models.IP, uint, error) {
	query := make(map[string]interface{})
	query["name"] = d.Name
	query["type"] = d.Type

	pageSize := d.PageSize
	page := d.Page

	return models.GetIPs(query, page, pageSize)
}

func (d *IP) Save() error {
	data := map[string]interface{}{
		"name":   d.Name,
		"type":   d.Type,
		"ip":     d.IP,
		"remark": d.Remark,
	}

	var err error
	if d.ID > 0 {
		err = models.UpdateIP(d.ID, data)
	} else {
		err = models.AddIP(data)
	}

	return err
}

func (d *IP) Delete() error {
	return models.DeleteIP(d.ID)
}

func (d *IP) Distribute() error {
	ips, _, err := models.GetIPs(nil, 0, 0)
	if err != nil {
		return err
	}

	var deny []string
	var allow []string
	for _, ipItem := range ips {
		bytes, err := ipItem.IP.MarshalJSON()
		if err != nil {
			return err
		}

		var iparray []string
		err = json.Unmarshal(bytes, &iparray)
		if err != nil {
			return err
		}

		if ipItem.Type == 1 {
			allow = append(allow, iparray...)
		} else if ipItem.Type == 2 {
			deny = append(deny, iparray...)
		}
	}

	if len(allow) == 0 {
		allow = []string{}
	}
	if len(deny) == 0 {
		deny = []string{}
	}

	var ipallow cacheHead
	var ipdeny cacheHead
	ipallow.ID = 1
	ipallow.Timestamp = time.Now().Unix()
	ipallow.Config = map[string]interface{}{
		"type": "allow",
		"data": allow,
	}
	ipdeny.ID = 1
	ipdeny.Timestamp = time.Now().Unix()
	ipdeny.Config = map[string]interface{}{
		"type": "deny",
		"data": deny,
	}

	var items []cacheHead
	items = append(items, ipallow)
	items = append(items, ipdeny)

	config := make(map[string]interface{})
	config["values"] = items
	config["timestamp"] = time.Now().Unix()

	str, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = cache.Set("redis", cacheIP, string(str), 0)
	if err != nil {
		return err
	}

	return nil
}
