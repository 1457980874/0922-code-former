package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type RegisterController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *RegisterController) Post(){
	fmt.Println("already go run")
	name:=c.Ctx.Request.PostFormValue("name")
	birthday:=c.Ctx.Request.PostFormValue("birthday")
	address:=c.Ctx.Request.PostFormValue("address")
	nick:=c.Ctx.Request.PostFormValue("nick")
	fmt.Println(name,birthday,address,nick)
	if name !="chen" {
		c.Ctx.WriteString("数据解析失败")
		return
	}
	c.Ctx.WriteString("数据解析成功")
}
/*
*post方法
*/
func (c *MainController) Post(){
	fmt.Println("hello world")
	name := c.Ctx.Request.PostFormValue("Name")
	age := c.Ctx.Request.PostFormValue("Age")
	sex := c.Ctx.Request.PostFormValue("Sex")
	fmt.Println(name,age,sex)
	if name!="chen" && age!="20" && sex!="man" {
		c.Ctx.WriteString("数据校验失败")
		return
	}
	c.Ctx.WriteString("数据校验成功")
}


//func (c *MainController)Post()  {
//	var person models.Person
//	database,err:=ioutil.ReadAll(c.Ctx.Request.Body)
//	if err != nil {
//		c.Ctx.WriteString("数据接受错误")
//		return
//	}
//	err =json.Unmarshal(database,&person)
//	if err != nil {
//		c.Ctx.WriteString("数据解析错误")
//		return
//	}
//	fmt.Println("姓名：",person.Name)
//	fmt.Println("年龄：",person.Age)
//	fmt.Println("性别：",person.Sex)
//	c.Ctx.WriteString("数据解析成功")
//}

