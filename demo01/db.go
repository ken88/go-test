package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Admin struct {
	Id       int
	Username string
	Password string
}

func main() {
	a := 7
	if a > 1 && a < 5 || a > 5 && a < 11 {
		fmt.Println("满足条件")
	}
}
func selectOne() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Panicln(err)
	}
	// 执行完毕后 关闭数据库链接
	defer db.Close()

	//mod := &Admin{}

}

func insertInfo() {

	// 系统启动的时候就加载(并不是用户访问的时候加载) 因为连接池 只是用户使用的时候就会建立连接 用完返回
	/*                 用户名:密码@tcp(链接地址:端口)/数据库名?parseTime=true")
	sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?parseTime=true")
	*/
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?parseTime=true")
	if err != nil {
		log.Panicln(err)
	}
	// 执行完毕后 关闭数据库链接
	defer db.Close()

	rs, err := db.Exec("INSERT INTO test.admin(username,password) VALUES ('zhangsan','123456')")
	if err != nil {
		log.Fatalln(err)
	}
	rowCount, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("inserted %d rows", rowCount)

}

func ConnectDb() {
	fmt.Println("链接数据库")
}
