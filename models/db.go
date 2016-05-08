package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yakun0622/goblog/models/class"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/blog?charset=utf8")
	orm.RegisterModel(new(class.User))

	orm.RunSyncdb("default", false, true)
}
