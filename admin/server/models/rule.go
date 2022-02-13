package models

import (
	"github.com/jinzhu/gorm"
)

type Rule struct {
	Model

	Name     string `json:"name" gorm:"column:name;unique;comment:'规则名称'"`
	Type     int    `json:"type" gorm:"column:type;comment:'规则类型'"`
	IP       string `json:"ip" gorm:"column:ip;comment:'客户端IP'"`
	User     string `json:"user" gorm:"column:user;comment:'数据库用户'"`
	Database string `json:"database" gorm:"column:database;comment:'数据库名称'"`
	Sql      string `json:"sql" gorm:"column:sql;comment:'规则'"`

	Remark string `json:"remark" gorm:"column:remark;comment:'备注'"`
}

func AddRule(data map[string]interface{}) error {
	rule := Rule{
		Name:     data["name"].(string),
		Type:     data["type"].(int),
		IP:       data["ip"].(string),
		User:     data["user"].(string),
		Database: data["database"].(string),
		Sql:      data["sql"].(string),
		Remark:   data["remark"].(string),
	}
	return db.Create(&rule).Error
}

func DeleteRule(id uint) error {
	var rule Rule
	rule.Model.ID = id
	return db.Where("id = ?", id).Delete(Rule{}).Error
}

func UpdateRule(id uint, data map[string]interface{}) error {
	return db.Model(&Rule{}).Where("id = ?", id).Update(data).Error
}

func GetRule(id uint) (*Rule, error) {
	var rule Rule
	err := db.Where("id = ?", id).First(&rule).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &rule, nil
}

func GetRules(query map[string]interface{}, page int, pageSize int) ([]*Rule, uint, error) {
	var rules []*Rule
	var count uint
	var err error

	pageNum := (page - 1) * pageSize

	var name string
	if query["name"] != nil {
		name = query["name"].(string)
	}

	search := make(map[string]interface{})

	typ := query["type"].(int)
	if typ > 0 {
		search["type"] = typ
	}

	if len(name) > 0 {
		name = "%" + name + "%"
		err = db.Where("name like ?", name).Where(search).Offset(pageNum).Limit(pageSize).Find(&rules).Count(&count).Error
	} else {
		err = db.Where(search).Offset(pageNum).Limit(pageSize).Find(&rules).Count(&count).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}

	return rules, count, nil
}
