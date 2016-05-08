package main

import (
	"github.com/astaxie/beego"
	_ "github.com/yakun0622/goblog/models"
	"github.com/yakun0622/goblog/models/class"
	_ "github.com/yakun0622/goblog/routers"
)

func main() {
	class.TestORM()
	beego.Run()
}
