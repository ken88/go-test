package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"gotest/demo01/gorm"
	"strconv"
)

var client *elastic.Client

func init() {

	var err error
	client, err = elastic.NewClient(
		elastic.SetURL("http://localhost:9200"), // 服务器连接地址 ，多个用逗号分开
		// elastic.SetErrorLog(log.New(os.Stderr, "esApp ", log.LstdFlags)),
		elastic.SetSniff(false), // 关闭Sniff  服务器用的是docker 不关闭访问不到
		elastic.SetHealthcheck(false),
	)
	if err != nil {
		fmt.Println("链接错误:", err.Error())
	}
	info, code, err := client.Ping("http://localhost:9200").Do(context.Background())
	if err != nil {
		fmt.Println("ping错误:", err.Error())
	}
	fmt.Printf("elasticsearch返回版本%s code:%d\n ", info.Version.Number, code)
}

// CreateIndex 创建索引 mapping
func CreateIndex() (bool, error) {

	// 1. mapping
	mapping := `{
		"settings" : {
			"number_of_shards": 1,
			"number_of_replicas": 0
		},

		"mappings": {
			"properties": {
				"Id":{
					"type": "integer"
				},
				"WebId":{
					"type": "integer"
				},
				"GoodsSn":{
						"type": "keyword"
				},
				"SeeNum":{
					"type": "integer"
				},
				"CartNum":{
					"type": "integer"
				},
				"PayNum":{
					"type": "integer"
				},
				"GoodsAmount":{
						"type": "double"
				},
				"AddDate":{
						"type": "date",
						"format": "yyyy-MM-dd||yyyy/MM/dd||yyyy-MM-dd HH:mm:ss"
				},
				"StatDate":{
						"type": "date"
				}
			}
		}
	  }`

	// 2. 创建索引 与 mapping
	ctx := context.Background()
	createIndex, err := client.CreateIndex("goods-ga").BodyString(mapping).Do(ctx)

	if err != nil {
		fmt.Println("返回错误err", err.Error())
		return false, err
	} else if !createIndex.Acknowledged {
		fmt.Println("返回错误", createIndex.Acknowledged)
		return false, nil
	}
	return true, nil
}

type GoodsGa struct {
	Id          int
	WebId       int
	SeeNum      int
	PayNum      int
	CartNum     int
	GoodsAmount float64
	GoodsSn     string
	StatDate    string
	AddDate     string
}

// BulkAll 批量增加数据
func BulkAll() {
	goodsGa := []GoodsGa{}

	limits := 100000          // 每次查询的数量，
	for i := 1; i <= 3; i++ { // 表里总共234838 循环3次 每次取出100000条

		// 1. 获取mysql数据
		gorm.DB.Table("goods_ga_sku").
			Offset((i - 1) * limits).
			Limit(limits).
			Find(&goodsGa)
		//fmt.Printf("mysql数据 %+v\n", goodsGa)

		fmt.Printf("%d查询数据数据\n", i)

		// 2. 批量追加数据 每次3万条
		bulkRequest := client.Bulk()
		for _, v := range goodsGa {

			req := elastic.NewBulkIndexRequest().
				Index("goods-ga").      // 索引名
				Id(strconv.Itoa(v.Id)). // id 不增加会自动生成uuid
				Doc(v)                  // 数据
			bulkRequest = bulkRequest.Add(req) // 追加数据
		}

		// 3. 增加到es数据
		bulkResponse, err := bulkRequest.Do(context.Background())
		if err != nil {
			fmt.Printf("%d 返回错误:%v\n", i, err.Error())
		}

		// 4. 返回状态吗  正常返回201
		indexed := bulkResponse.Indexed()
		fmt.Printf("%d 返回数据:%#v\n", i, indexed[0].Status)
	}

	fmt.Println("操作完成")
}
func main() {
	// 1. 创建索引 与mapping
	//CreateIndex()

	// 2. 批量加入索引
	BulkAll()
}
