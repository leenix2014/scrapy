package entity

import (
	"time"
)

type TPdf struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	UserMail   string    `xorm:"comment('用户邮箱') unique(user_mail) VARCHAR(64)"`
	Root       string    `xorm:"comment('pdf根url') VARCHAR(512)"`
	Url        string    `xorm:"comment('pdf下载url') unique(user_mail) VARCHAR(512)"`
	Visited    int       `xorm:"not null default 0 comment('是否已访问') TINYINT(1)"`
	Createtime time.Time `xorm:"created comment('创建时间') DATETIME"`
	Updatetime time.Time `xorm:"updated comment('更新时间') DATETIME"`
}
