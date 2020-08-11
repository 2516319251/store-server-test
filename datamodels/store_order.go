package datamodels

import "time"

type StoreOrder struct {
	OrderId    int64     `xorm:"not null pk autoincr comment('用户id') INT(11)"`
	UserId     int64     `xorm:"not null default 0 comment('用户id') INT(11)"`
	ProductId  int64     `xorm:"not null default 0 comment('商品id') INT(11)"`
	Status     int64     `xorm:"not null default 0 comment('订单是否已完成(0未完成 1已完成)') TINYINT(3)"`
	CreateTime time.Time `xorm:"not null default 0 comment('创建时间') INT(11)"`
}

type OrderInformation struct {
	OrderId     int64     `xorm:"not null pk autoincr comment('用户id') INT(11)"`
	ProductName string    `xorm:"not null default '' comment('商品名称') VARCHAR(255)"`
	Nickname    string    `xorm:"not null default '' comment('昵称') VARCHAR(255)"`
	Status      int64     `xorm:"not null default 0 comment('订单是否已完成(0未完成 1已完成)') TINYINT(3)"`
	CreateTime  time.Time `xorm:"not null default 0 comment('创建时间') INT(11)"`
}
