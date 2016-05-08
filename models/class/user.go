package class

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id    string `orm:"pk"`
	Nick  string
	Email string
}

func TestORM() {
	o := orm.NewOrm()

	u := User{"wyk", "yakun", "123@qq.com"}
	o.Insert(&u)

	u1 := User{Id: "wyk"}
	o.Read(&u1)
	fmt.Println(u1)

	u1.Nick = "lily"
	o.Update(&u1)

	u2 := User{Id: "wyk"}
	o.Read(&u2)
	fmt.Println(u2)

	o.Delete(&u2)
	o.Read(&u2)
}
