package service

import "admin/server/models"

type DB struct {
	ID uint

	Name     string
	User     string
	Password string
	Host     string
	Port     uint32

	Remark string

	Page     int
	PageSize int
}

func (d *DB) Get() (*models.DB, error) {
	dbItem, err := models.GetDB(d.ID)
	if err != nil {
		return nil, err
	}
	return dbItem, nil
}

func (d *DB) GetList() ([]*models.DB, uint, error) {
	query := make(map[string]interface{})
	query["name"] = d.Name

	pageSize := d.PageSize
	page := d.Page

	return models.GetDBs(query, page, pageSize)
}

func (d *DB) Save() error {
	data := map[string]interface{}{
		"name":     d.Name,
		"user":     d.User,
		"password": d.Password,
		"status":   d.Host,
		"port":     d.Port,
		"remark":   d.Remark,
	}

	var err error
	if d.ID > 0 {
		err = models.UpdateDB(d.ID, data)
	} else {
		err = models.AddDB(data)
	}

	return err
}

func (d *DB) Delete() error {
	return models.DeleteDB(d.ID)
}
