package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type DB struct {
	Model

	Name      string `json:"name" gorm:"column:name;unique;comment:'数据库名称'"`
	User      string `json:"user" gorm:"column:user;comment:'数据库用户'"`
	Password  string `json:"password" gorm:"column:password;comment:'数据库用户'"`
	Host      string `json:"host" gorm:"column:host;comment:'数据库服务器IP'"`
	Port      uint32 `json:"port" gorm:"column:port;comment:'数据库服务器Port'"`
	Charset   string `json:"charset" gorm:"column:charset;comment:'数据库字符集'"`
	Collation string `json:"collation" gorm:"column:collation;comment:'数据库字符序'"`
	Remark    string `json:"remark" gorm:"column:remark;comment:'备注'"`

	DBUser []DBUser `json:"dbuser" gorm:"many2many:user_db;"`
}

func AddDB(data map[string]interface{}) error {
	dbUser := DB{
		Name:      data["name"].(string),
		User:      data["user"].(string),
		Password:  data["password"].(string),
		Host:      data["host"].(string),
		Port:      data["port"].(uint32),
		Charset:   data["charset"].(string),
		Collation: data["collation"].(string),

		Remark: data["remark"].(string),
	}
	return db.Create(&dbUser).Error
}

func DeleteDB(id uint) error {
	var dbItem DB
	dbItem.Model.ID = id
	count := db.Model(&dbItem).Association("db_user").Count()
	if count > 0 {
		return fmt.Errorf("%s", "db exist link with dbuser")
	}
	return db.Where("id = ?", id).Delete(DB{}).Error
}

func UpdateDB(id uint, data map[string]interface{}) error {
	return db.Model(&DB{}).Where("id = ?", id).Update(data).Error
}

func GetDB(id uint) (*DB, error) {
	var dbItem DB
	err := db.Where("id = ?", id).First(&dbItem).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &dbItem, nil
}

func GetDBs(query map[string]interface{}, page int, pageSize int) ([]*DB, uint, error) {
	var dbItems []*DB
	var count uint
	var err error
	if page > 0 {
		pageNum := (page - 1) * pageSize

		var name string
		if query["name"] != nil {
			name = query["name"].(string)
		}

		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Where("name like ?", name).Offset(pageNum).Limit(pageSize).Find(&dbItems).Count(&count).Error
		} else {
			err = db.Offset(pageNum).Limit(pageSize).Find(&dbItems).Count(&count).Error
		}
	} else {
		err = db.Model(DB{}).Find(&dbItems).Count(&count).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}

	return dbItems, count, nil
}
