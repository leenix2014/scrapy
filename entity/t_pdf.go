package entity

import (
	"time"
)

type TPdf struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	UserMail   string    `xorm:"comment('用户邮箱') VARCHAR(256)"`
	Root       string    `xorm:"comment('pdf根url') VARCHAR(1024)"`
	Url        string    `xorm:"comment('pdf下载url') VARCHAR(1024)"`
	Visited    int       `xorm:"default 0 comment('是否已访问') TINYINT(1)"`
	Createtime time.Time `xorm:"comment('创建时间') DATETIME"`
	Updatetime time.Time `xorm:"comment('更新时间') DATETIME"`
}
