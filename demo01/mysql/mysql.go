package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type Model struct {
	link  *sql.DB // 存储连接对象
	field string  // 存储字段
	table string  // 表名
	order string  // 排序
	limit string  // 存储limit条件
	where string  // 存储where条件
}

type ApiRes struct {
	Code   int
	Result map[string]interface{}
	Msg    string
}

/**
 *  构造方法
 * @param string table 初始化数据表
 */
func NewModel(table string) Model {

	var model Model
	model.field = "*"
	// 1. 存储操作的表名
	model.table = table

	// 2. 初始化链接数据库
	model.getConnect()

	return model
}

// 链接数据库
func (model *Model) getConnect() {

	//1.连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/acl?parseTime=true")
	//db, err := sql.Open("mysql", "super-xxx:9HAkMVxWXIkX9HBD1@tcp(rm-t4n1a1m62fyvq1913o.mysql.singapore.rds.aliyuncs.com:3306)/product_center_online?parseTime=true")
	//2.判断连接
	if err != nil {
		fmt.Printf("connect mysql fail! [%s]", err)
	}

	model.link = db
}

/**
 *  设置要查询的字段信息
 * @param string field  要查询的字段
 * @return object 返回自己，保证连贯操作
 */
func (this *Model) Field(field string) *Model {
	this.field = field
	return this
}

/**
 * where条件
 * @param string $where 输入的where条件
 * @return $model 返回自己，保证连贯操作
 */
func (this *Model) Where(where string) *Model {
	this.where = " where " + where
	return this
}

/**
 * 执行并发送SQL语句(增删改)
 * @param string $sql 要执行的SQL语句
 * @return bool|int|string 添加成功则返回新录入的id,删除修改操作则返回true,失败则返回false
 */
func (model *Model) exce(sql string) interface{} {

	res, err := model.link.Exec(sql)

	// 关闭连接
	defer model.link.Close()

	if err != nil {
		returnRes(0, ``, err.Error())
	}
	result, err := res.LastInsertId()
	if err != nil {
		returnRes(0, ``, err.Error())
	}
	return returnRes(1, result, "success")
}

//返回json
func returnRes(code int, res interface{}, msg interface{}) string {
	result := make(map[string]interface{})
	result["code"] = code
	result["result"] = res
	result["msg"] = msg
	// 转化json字符串
	data, _ := json.Marshal(result)
	return string(data)
}

/**
 * 新增数据方法
 * @param map data  要插入的数据
 * @return bool 成功返回true 失败返回false
 */
func (model *Model) Add(data map[string]interface{}) interface{} {

	var key, value string

	for k, v := range data {
		key += "," + k
		value += `,` + `'` + v.(string) + `'`
	}

	// 去掉拼接的左侧,号
	key = strings.TrimLeft(key, ",")

	//将map中的值转化为字符串拼接
	value = strings.TrimLeft(value, ",")

	// 准备sql语句
	sql := `insert into ` + model.table + `(` + key + `) values (` + value + `)`

	// 执行sql
	result := model.exce(sql)
	return result
}

/**
 * 删除数据方法
 * @return bool 成功返回true 失败返回false
 */
func (model *Model) Delete() interface{} {
	sql := "delete from " + model.table + model.where
	// 执行sql
	result := model.exce(sql)
	return result
}

/**
 * 修改数据方法
 * @param map data  要修改的数据
 * @return bool 成功返回true 失败返回false
 */
func (model *Model) Update(data map[string]interface{}) interface{} {
	var str string
	for k, v := range data {
		str += "," + k + `='` + v.(string) + `'`
	}
	//将map中取出的键转为字符串拼接
	str = strings.TrimLeft(str, ",")

	sql := "update " + model.table + " set " + str + model.where

	// 执行sql
	result := model.exce(sql)
	return result
}

// 返回 []map[string]interface{}
func (model *Model) Query(sql string) []map[string]string {

	rows, _ := model.link.Query(sql)

	// fmt.Println(rows) // &{0xc000114000 0x6e3b40 0xc000040140 <nil> <nil> {{0 0} 0 0 0 0} false <nil> []}

	// 返回所有列
	cols, _ := rows.Columns()

	// fmt.Println(cols)

	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))

	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))

	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	//result := make(map[int]map[string]string)
	result := make([]map[string]string, 0)
	for rows.Next() { // 如果有数据 执行

		//填充数据
		rows.Scan(scans...) //将slic地址传入

		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]

			//这里把[]byte数据转成string
			row[key] = string(v)

		}
		//放入结果集
		//result[i] = row
		result = append(result, row)
		i++
	}

	// fmt.Println(result)
	return result

}

// 返回json数组
func (model *Model) Query1(sql string) interface{} {

	rows, err := model.link.Query(sql)
	if err != nil {
		return returnRes(0, ``, err)
	}
	// fmt.Println(rows) // &{0xc000114000 0x6e3b40 0xc000040140 <nil> <nil> {{0 0} 0 0 0 0} false <nil> []}

	// 返回所有列
	cols, err := rows.Columns()
	if err != nil {
		return returnRes(0, ``, err)
	}
	// fmt.Println(cols)

	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))

	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))

	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	i := 0
	result := make(map[int]map[string]string)

	for rows.Next() { // 如果有数据 执行

		//填充数据
		rows.Scan(scans...) //将slic地址传入

		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}

	// fmt.Println(result)
	return returnRes(1, result, "success")

}

func (model *Model) Find() interface{} {
	sql := "select " + model.field + " from " + model.table + model.where + " limit 1"

	//执行并发送sql
	result := model.Query(sql)
	return result
}

func (model *Model) All() interface{} {
	sql := "select " + model.field + " from " + model.table + model.where + model.limit
	//执行并发送sql
	result := model.Query(sql)
	return result
}

func (model *Model) Order(order string) *Model {
	model.order = " order by " + order
	return model
}

func (model *Model) Limit(limit string) *Model {
	model.limit = " limit " + limit
	return model
}
