package main

import (
	_ "0922/db_mysql"
	_ "0922/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	config:=beego.AppConfig
	appname:=config.String("appname")
	fmt.Println("应用名称：",appname)
	port,err:=config.Int("httpport")
	if err != nil {
		panic("项目配置文件解析失败，请检查配置文件")
	}
	fmt.Println("端口：",port)
	beego.Run()
}

