package service

import (
	"admin/server/cache"
	"admin/server/models"
	"encoding/json"
	"time"
)

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
	data["dbs"] = d.Dbs
	return models.UpdateDBUserDB(d.ID, data)
}

func (d *DBUser) Distribute() error {
	users, _, err := models.GetDBUsers(nil, 0, 0)
	if err != nil {
		return err
	}

	var items []cacheHead
	for _, user := range users {
		var item dbUser
		item.User = user.Username
		item.Password = user.Password
		for _, db := range user.Dbs {
			var dbItem database
			dbItem.Name = db.Name
			dbItem.User = db.User
			dbItem.Password = db.Password
			dbItem.Host = db.Host
			dbItem.Port = db.Port
			dbItem.Charset = db.Charset
			dbItem.Collation = db.Collation

			item.Database = append(item.Database, dbItem)
		}
		var cacheItem cacheHead
		cacheItem.ID = user.ID
		cacheItem.Timestamp = user.Timestamp
		if len(item.Database) == 0 {
			item.Database = []database{}
		}
		cacheItem.Config = item
		items = append(items, cacheItem)
	}

	var data interface{}
	if len(items) == 0 {
		data = [][]struct{}{}
	} else {
		data = items
	}

	config := make(map[string]interface{})
	config["values"] = data
	config["timestamp"] = time.Now().Unix()

	str, err := json.Marshal(config)
	if err != nil {
		return err
	}

	err = cache.Set("redis", cacheUser, string(str), 0)
	if err != nil {
		return err
	}

	return nil
}
