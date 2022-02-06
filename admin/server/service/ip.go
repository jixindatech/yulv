package service

import "admin/server/models"

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
