package datamodels

import "time"

type StoreProduct struct {
	ProductId    int64     `xorm:"not null pk autoincr comment('商品id') INT(10)"`
	ProductName  string    `xorm:"not null default '' comment('商品名称') VARCHAR(255)"`
	ProductNum   int64     `xorm:"not null default 0 comment('当前库存数量') INT(11)"`
	ProductImage string    `xorm:"not null default '' comment('商品图片地址') VARCHAR(255)"`
	ProductUrl   string    `xorm:"not null default '' comment('商品链接地址') VARCHAR(255)"`
	CreateTime   time.Time `xorm:"not null default 0 comment('创建时间') INT(11)"`
}
