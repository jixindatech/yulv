package models

import (
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
)

type IP struct {
	Model

	Name string         `json:"name" gorm:"column:name;unique;comment:'IP名称'"`
	Type int            `json:"type" gorm:"column:type;not null;comment:'Type'"`
	IP   datatypes.JSON `json:"ip" gorm:"column:ip;type:VARBINARY(1024);not null;comment:'IP'"`

	Remark string `json:"remark" gorm:"column:remark;comment:'备注'"`
}

func AddIP(data map[string]interface{}) error {
	ip := IP{
		Name:   data["name"].(string),
		Type:   data["type"].(int),
		IP:     data["ip"].([]byte),
		Remark: data["remark"].(string),
	}
	return db.Create(&ip).Error
}

func DeleteIP(id uint) error {
	return db.Where("id = ?", id).Delete(IP{}).Error
}

func UpdateIP(id uint, data map[string]interface{}) error {
	return db.Model(&IP{}).Where("id = ?", id).Update(data).Error
}

func GetIP(id uint) (*IP, error) {
	var ip IP
	err := db.Where("id = ?", id).First(&ip).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &ip, nil
}

func GetIPs(query map[string]interface{}, page int, pageSize int) ([]*IP, uint, error) {
	var ips []*IP
	var count uint
	var err error

	if page > 0 {
		pageNum := (page - 1) * pageSize

		search := make(map[string]interface{})
		name := query["name"].(string)

		ipType := query["type"].(int)
		if ipType > 0 {
			search["type"] = ipType
		}

		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Where(search).Where("name LIKE ?", name).Offset(pageNum).Limit(pageSize).Find(&ips).Count(&count).Error
		} else {
			err = db.Where(search).Offset(pageNum).Limit(pageSize).Find(&ips).Count(&count).Error
		}
	} else {
		err = db.Find(&ips).Count(&count).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}

	return ips, count, nil
}
