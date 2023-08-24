package biz

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type  Customer struct {
	CustomerWork
	CustomerToken
	gorm.Model

}
type CustomerWork struct {
	Telephone string `gorm:"type:varchar(15);uniqueIndex;"json:"telephone"`
	Name string`gorm:"type:varchar(255);",json:"name"`
	Wechat string`gorm:"type:varchar(255);",json:"wechat"`
	Email string`gorm:"type:varchar(255);",json:"email"`
}

type CustomerToken struct {
	Token string `gorm:"varchar(4095);"json:"token"`
	TokenCreateAt sql.NullTime `json:"token_create_at"`
}

const Duration = 30 * 60 * 60 * 24 * time.Second
const CustomerSecret = "my_secret_key"