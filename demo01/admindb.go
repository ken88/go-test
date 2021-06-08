package main

import (
	"demo01/mysql"
	"encoding/json"
	"fmt"
)

func insert() {
	M := mysql.NewModel("admin")
	data := make(map[string]interface{})
	data["username"] = "a2"
	data["password"] = "123456"
	res := M.Add(data)
	fmt.Println(res)

	var res1 mysql.ApiRes
	err := json.Unmarshal([]byte(res.(string)), &res1)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("%#v %T\n", res1, res1)

	// fmt.Printf("返回key:%v 返回值：%v  返回值类型：%T\n", res1.Code, res1.Code, res1.Code)

	if res1.Code == 1 {
		fmt.Println("录入成功！")
	} else {
		fmt.Println("录入失败！")
	}
}

func update() {
	M := mysql.NewModel("admin")
	data := make(map[string]interface{})
	data["username"] = "张三3"
	data["password"] = "87654321"
	res := M.Where("id=41").Update(data)
	var res1 mysql.ApiRes
	err := json.Unmarshal([]byte(res.(string)), &res1)
	if err != nil {
		fmt.Println(err)
	}
	if res1.Code == 1 {
		fmt.Println("修改成功！")
	} else {
		fmt.Println("修改失败！")
	}
}

func delete() {
	M := mysql.NewModel("admin")
	res := M.Where("id>=41").Delete()
	var res1 mysql.ApiRes
	err := json.Unmarshal([]byte(res.(string)), &res1)
	if err != nil {
		fmt.Println(err)
	}
	if res1.Code == 1 {
		fmt.Println("删除成功！")
	} else {
		fmt.Println("删除失败！")
	}
}

func find() {
	M := mysql.NewModel("admin")
	// M.Find() //查询所有字段 select * from admin limit 1
	// M.Field("id,username").Find() // select id,username from admin limit 1

	var res1 mysql.ApiRes
	res := M.Field("id,username").Where("id>45").Find()
	// fmt.Println(res)
	err := json.Unmarshal([]byte(res.(string)), &res1)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range res1.Result {
		info, _ := v.(map[string]interface{})
		for key, val := range info {
			fmt.Printf("key:%v val:%v\n", key, val)
		}
	}
}

func all() {
	M := mysql.NewModel("admin")
	var res1 mysql.ApiRes
	res := M.Field("id,username").Where("id>45").All()
	err := json.Unmarshal([]byte(res.(string)), &res1)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range res1.Result {
		info, _ := v.(map[string]interface{})
		for key, val := range info {
			fmt.Printf("key:%v val:%v\n", key, val)
		}
	}
}
func main() {
	// insert()
	// delete()
	// update()
	// find()
	// find()
	all()

}
