package db_mysql

import (
	"0922/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)
var Db *sql.DB
func init(){
	fmt.Println("链接mysql数据库")
	config := beego.AppConfig
	dbDriver:=config.String("driverName")
	dbUser:=config.String("db_user")
	dbPassword:=config.String("db_password")
	dbIp :=config.String("db_ip")
	dbName:=config.String("db_name")
	connUrl:=dbUser+":"+dbPassword+"@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	fmt.Println(connUrl)
	db,err:=sql.Open(dbDriver,connUrl)
	if err != nil {
		fmt.Println(err.Error())
		panic("数据连接错误，请检查错误")
	}
	Db=db

}

func InsertUser(user models.User)(int64,error){
	hashMD5 :=md5.New()
	hashMD5.Write([]byte(user.Password))
	bytes :=hashMD5.Sum(nil)
	user.Password=hex.EncodeToString(bytes)
	fmt.Println("保存的用户名：",user.Nick,"密码：",user.Password)
	result,err:=Db.Exec("insert into user(nick,password) value(?,?)",user.Nick,user.Password)
	if err != nil {
		return -1,err
	}
	id,err:=result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return id,nil
}

func QueryUser(){
	Db.QueryRow("select * from ")
}
