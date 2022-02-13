package service

import "admin/server/models"

type Rule struct {
	ID uint

	Name     string
	Type     int
	IP       string
	User     string
	Database string
	Sql      string

	Remark string

	Page     int
	PageSize int
}

func (d *Rule) Get() (*models.Rule, error) {
	dbItem, err := models.GetRule(d.ID)
	if err != nil {
		return nil, err
	}
	return dbItem, nil
}

func (d *Rule) GetList() ([]*models.Rule, uint, error) {
	query := make(map[string]interface{})
	query["name"] = d.Name
	query["type"] = d.Type
	pageSize := d.PageSize
	page := d.Page

	return models.GetRules(query, page, pageSize)
}

func (d *Rule) Save() error {
	data := map[string]interface{}{
		"name":     d.Name,
		"type":     d.Type,
		"user":     d.User,
		"ip":       d.IP,
		"database": d.Database,
		"sql":      d.Sql,
		"remark":   d.Remark,
	}

	var err error
	if d.ID > 0 {
		err = models.UpdateRule(d.ID, data)
	} else {
		err = models.AddRule(data)
	}

	return err
}

func (d *Rule) Delete() error {
	return models.DeleteRule(d.ID)
}
