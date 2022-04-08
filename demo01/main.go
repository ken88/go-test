package main

import "fmt"

func main() {
	var game = 0
	var dianji = 0
	for i := 0; i < 4; i++ {
		fmt.Println("请输入第", i+1, "个游戏的点击率:")

		fmt.Scan(&game)

		if dianji > 100 {
			game++
		}
	}
	fmt.Println("点击率大于100的游戏数是: ", game)
	fmt.Println("点击率大于100的游戏所占的比例为: ", float64(game/4.0*100), "%")
}
