package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {

	// 数据库链接
	driverName := "mysql"
	user := "root"
	pwd := "123456"
	host := "127.0.0.1"
	port := "3306"
	dbname := "acl"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	DB, err = gorm.Open(driverName, dbConn)
	DB.SingularTable(true) // 让grom转义struct名字的时候不用加上s
	// DB.LogMode(true)       // 开启sql打印信息

	if err != nil {
		fmt.Println("mysql链接失败!")
	} else {
		fmt.Println("mysql链接成功")
	}
}
