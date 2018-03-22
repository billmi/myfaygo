package model

import (
	"time"

	"github.com/henrylee2cn/faygo/ext/db/xorm"
)

//国家
type Country struct {
	Id       int64     `json:"id"`
	Nickname string    `xorm:"'nickname' notnull default ''" json:"nickname"` //昵称
	Name     string    `xorm:"'name' notnull default ''" json:"name"`         //中文名称
	NameEn   string    `xorm:"'name_en' notnull default ''" json:"name_en"`   //英文名称
	Created  time.Time `xorm:"created" json:"created"`
	Updated  time.Time `xorm:"updated" json:"updated"`
	Deleted  time.Time `xorm:"deleted" json:"deleted"`
}

func init() {
	xorm.MustDB().Sync2(new(Country))
}

func (Country) TableName() string {
	return "country"
}
