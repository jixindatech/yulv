package service

import "admin/server/models"

type DBUser struct {
	ID uint

	Name     string
	Username string
	Password string

	Remark string

	Page     int
	PageSize int

	Dbs []uint
}

func (d *DBUser) Get() (*models.DBUser, error) {
	dbUser, err := models.GetDBUser(d.ID)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}

func (d *DBUser) GetList() ([]*models.DBUser, uint, error) {
	query := make(map[string]interface{})
	query["name"] = d.Name

	pageSize := d.PageSize
	page := d.Page

	return models.GetDBUsers(query, page, pageSize)
}

func (d *DBUser) Save() error {
	data := map[string]interface{}{
		"name":     d.Name,
		"username": d.Username,
		"password": d.Password,
		"remark":   d.Remark,
	}

	var err error
	if d.ID > 0 {
		err = models.UpdateDBUser(d.ID, data)
	} else {
		err = models.AddDBUser(data)
	}

	return err
}

func (d *DBUser) Delete() error {
	return models.DeleteDBUser(d.ID)
}

func (d *DBUser) UpdateUserDB() error {
	data := make(map[string]interface{})
	data["user"] = d.ID
	data["dbs"] = d.Dbs
	return models.UpdateDBUserDB(data)
}
