package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type DBUser struct {
	Model
	Name     string `json:"name" gorm:"column:name;unique;comment:'数据库用户名称'"`
	Username string `json:"username" gorm:"column:username;unique;comment:'数据库用户'"`
	Password string `json:"password" gorm:"column:password;comment:'数据库密码'"`

	Remark    string `json:"remark" gorm:"column:remark;comment:'备注'"`
	Timestamp int64  `json:"-" gorm:"column:timestamp;comment:'时间戳'"`

	Dbs []DB `json:"db" gorm:"many2many:user_db;"`
}

func (r *DBUser) AfterSave(tx *gorm.DB) (err error) {
	fmt.Println("save  ???")
	return nil
}

func AddDBUser(data map[string]interface{}) error {
	dbUser := DBUser{
		Name:      data["name"].(string),
		Username:  data["username"].(string),
		Password:  data["password"].(string),
		Remark:    data["remark"].(string),
		Timestamp: time.Now().Unix(),
	}
	return db.Create(&dbUser).Error
}

func DeleteDBUser(id uint) error {
	return db.Where("id = ?", id).Delete(DBUser{}).Error
}

func UpdateDBUser(id uint, data map[string]interface{}) error {
	data["timestamp"] = time.Now().Unix()
	return db.Model(&DBUser{}).Where("id = ?", id).Update(data).Error
}

func GetDBUser(id uint) (*DBUser, error) {
	var user DBUser
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func GetDBUsers(query map[string]interface{}, page int, pageSize int) ([]*DBUser, uint, error) {
	var users []*DBUser
	var count uint
	var err error

	pageNum := (page - 1) * pageSize

	var name string
	if query["name"] != nil {
		name = query["name"].(string)
	}

	if page > 0 {
		if len(name) > 0 {
			name = "%" + name + "%"
			err = db.Preload("Dbs").Where("name like ?", name).Offset(pageNum).Limit(pageSize).Find(&users).Count(&count).Error
		} else {
			err = db.Preload("Dbs").Offset(pageNum).Limit(pageSize).Find(&users).Count(&count).Error
		}
	} else {
		err = db.Preload("Dbs").Find(&users).Count(&count).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}

	return users, count, nil
}

func UpdateDBUserDB(userId uint, data map[string]interface{}) error {
	dbUser := DBUser{}
	dbUser.Model.ID = userId
	dbIds := data["dbs"].([]uint)
	var dbs []*DB
	for _, id := range dbIds {
		temp := DB{}
		temp.Model.ID = id
		dbs = append(dbs, &temp)
	}

	err := db.Model(&dbUser).Association("Dbs").Replace(dbs).Error
	if err != nil {
		return err
	}

	return db.Debug().Model(&DBUser{}).Where("id = ?", userId).Update(map[string]interface{}{
		"timestamp": time.Now().Unix(),
	}).Error
}
