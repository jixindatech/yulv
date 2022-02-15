package service

import (
	"admin/server/cache"
	"admin/server/models"
	"encoding/json"
)

type Rule struct {
	ID uint

	Name     string
	Action   int
	IP       string
	User     string
	Database string
	Type     string
	Sql      string
	Match    string
	Pattern  string
	Rows     int

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
	query["action"] = d.Action
	pageSize := d.PageSize
	page := d.Page

	return models.GetRules(query, page, pageSize)
}

func (d *Rule) Save() error {
	data := map[string]interface{}{
		"name":     d.Name,
		"action":   d.Action,
		"user":     d.User,
		"ip":       d.IP,
		"database": d.Database,
		"type":     d.Type,
		"sql":      d.Sql,
		"match":    d.Match,
		"pattern":  d.Pattern,
		"rows":     d.Rows,
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

func (d *Rule) Distribute() error {
	rules, _, err := models.GetRules(nil, 0, 0)
	if err != nil {
		return err
	}
	var reqRules []cacheHead
	var respRules []cacheHead
	for _, item := range rules {
		var head cacheHead
		head.ID = item.ID
		head.Timestamp = item.UpdatedAt.Unix()

		match := make(map[string]interface{})
		if len(item.IP) > 0 {
			match["ip"] = item.IP
		}
		if len(item.User) > 0 {
			match["user"] = item.User
		}
		if len(item.Database) > 0 {
			match["database"] = item.Database
		}
		if len(item.User) > 0 {
			match["user"] = item.User
		}
		if len(item.Match) > 0 {
			match["match"] = item.Match
			match["pattern"] = item.Pattern
		}
		if len(item.Sql) > 0 {
			match["sql"] = item.Sql
		}
		head.Config = map[string]interface{}{
			"matcher": match,
			"action":  item.Action,
		}

		match["rows"] = item.Rows

		if item.Rows == 0 {
			reqRules = append(reqRules, head)
		} else {
			respRules = append(respRules, head)
		}
	}

	if len(reqRules) == 0 {
		respRules = []cacheHead{}
	}
	if len(respRules) == 0 {
		respRules = []cacheHead{}
	}

	str, err := json.Marshal(reqRules)
	if err != nil {
		return err
	}

	err = cache.Set("redis", cacheReqRule, string(str), 0)
	if err != nil {
		return err
	}

	str, err = json.Marshal(respRules)
	if err != nil {
		return err
	}

	err = cache.Set("redis", cacheRespRule, string(str), 0)
	if err != nil {
		return err
	}

	return nil
}

func (d *Rule) Test(sql string) (string, error) {
	return GetFingerprint(sql), nil
}
