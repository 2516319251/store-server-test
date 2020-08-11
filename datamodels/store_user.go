package datamodels

import "time"

type StoreUser struct {
	UserId     int64     `xorm:"not null pk autoincr comment('用户id') INT(11)"`
	Nickname   string    `xorm:"not null default '' comment('昵称') VARCHAR(255)"`
	Username   string    `xorm:"not null default '' comment('用户名') VARCHAR(255)"`
	Password   string    `xorm:"not null default '' comment('密码') VARCHAR(255)"`
	CreateTime time.Time `xorm:"not null default 0 comment('创建时间') INT(11)"`
}
