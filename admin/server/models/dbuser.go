package models

import "github.com/jinzhu/gorm"

type DBUser struct {
	Model

	Username string `json:"username" gorm:"column:username;unique;comment:'数据库用户'"`
	Password string `json:"password" gorm:"column:password;comment:'数据库密码'"`

	Remark string `json:"remark" gorm:"column:remark;comment:'备注'"`

	Db []DB `json:"db" gorm:"many2many:user_db;"`
}

func AddDBUser(data map[string]interface{}) error {
	dbUser := DBUser{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Remark:   data["remark"].(string),
	}
	return db.Create(&dbUser).Error
}

func DeleteDBUser(id uint) error {
	return db.Where("id = ?", id).Delete(DBUser{}).Error
}

func UpdateDBUser(id uint, data map[string]interface{}) error {
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

	var username string
	if query["username"] != nil {
		username = query["username"].(string)
	}

	if len(username) > 0 {
		username = "%" + username + "%"
		err = db.Where("username like ?", username).Offset(pageNum).Limit(pageSize).Find(&users).Count(&count).Error
	} else {
		err = db.Offset(pageNum).Limit(pageSize).Find(&users).Count(&count).Error
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, count, err
	}

	return users, count, nil
}

func UpdateDBUserByDB(data map[string]interface{}) error {
	dbUser := DBUser{}
	dbUser.Model.ID = data["user_id"].(uint)
	dbIds := data["db_ids"].([]uint)
	var dbs []*DB
	for _, id := range dbIds {
		temp := DB{}
		temp.Model.ID = id
		dbs = append(dbs, &temp)
	}

	err := db.Model(&dbUser).Association("db").Replace(dbs).Error
	if err != nil {
		return err
	}
	return nil
}
