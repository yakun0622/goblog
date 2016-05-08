package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Profile() {
	c.Data["userid"] = "wyk"
	c.Data["tag"] = "I am a good person"
	c.Data["email"] = "123@qq.com"
	c.Data["hobby"] = []string{"chess", "code"}
	c.TplName = "user/profile.html"
}
