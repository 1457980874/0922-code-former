package controllers

import (
	"0922/db_mysql"
	"0922/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
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
func (r *RegisterController) Post(){
	fmt.Println("already go run")
	bodyBytes,err:=ioutil.ReadAll(r.Ctx.Request.Body)
	if err != nil {
		r.Ctx.WriteString("数据接收错误")
		return
	}
	var user models.User
	err =json.Unmarshal(bodyBytes,&user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("数据解析错误")
		return
	}
	id,err:=db_mysql.InsertUser(user)
	if err != nil {
		fmt.Println(err.Error())
		r.Ctx.WriteString("用户保存失败")
		return
	}
	fmt.Println(id)
	result :=models.ResponseResult{
		Code: 0,
		Message: "保存成功",
		Data: nil,
	}
	r.Data["json"]=&result
	r.ServeJSON()
}
/*
*post方法
*/
func (c *MainController) Post(){
	fmt.Println("hello world")
	name := c.Ctx.Request.FormValue("Name")
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

