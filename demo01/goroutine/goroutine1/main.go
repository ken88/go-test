package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup // 协程

func main() {
	for i := 0; i <= 20; i++ {
		time.Sleep(time.Millisecond * 5000)
		wg.Add(1)
		go getUrl(i)
	}
	wg.Wait()
	fmt.Println("hello word")
}

func getUrl(i int) {
	fmt.Printf("携程%d开始\n", i)
	url := "http://oms.lwhs.me/apicli.php?s=/api/TrackingLogistics/getInfo"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("携程%d结束\n", i)
	fmt.Println(string(body))
	wg.Done() // 协程计数器-1
}
