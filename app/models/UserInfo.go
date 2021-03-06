package models

import (
	_ "github.com/google/uuid"
)

type UserInfo struct {
	Id            string `json:"id" xorm:"pk varchar(64) 'id'"`
	LoginName     string `json:"loginName" xorm:"varchar(64) notnull 'login_name'"`
	LoginPassword string `json:"loginPassword" xorm:"varchar(128) notnull 'login_password'"`
	Salt          string `json:"salt" xorm:"varchar(128) notnull 'salt'"`
	Email         string `json:"email" xorm:"varchar(64) notnull 'email'"`

	UserName string `json:"userName" xorm:"varchar(64) notnull 'user_name'"`
	Img      string `json:"img" xorm:"varchar(256) default '' 'img'"`

	Valid int `json:"valid" xorm:"int default 0 'valid'"` //0 有效 1 无效  2 wait active

	CreatedAt int64 `json:"created" xorm:"bigint default 0 'created'"`
	UpdatedAt int64 `json:"updated" xorm:"bigint default 0 'updated'"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

var DefaultUserInfo = UserInfo{
	Id:            "init",
	LoginName:     "anyone",
	LoginPassword: "9fafdfffa1222faddec96c954ed67618aa78808842c8753d6d788e77258ada82",
	Salt:          "init",
	UserName:      "anyone",
	Email:         "anyone@cali.io",
}

var DefaultAdminUserInfo = UserInfo{
	Id:            "admin",
	LoginName:     "admin",
	LoginPassword: "3ef7b7e37a0fe84e4c5cdcd1934db0852f608c5c925751b9ef6cf872eb6eeaca",
	Salt:          "init",
	UserName:      "admin",
	Email:         "admin@cali.io",
}
