package main

import (
	"time"
	"fmt"
	"strconv"
)

type human struct {
	name  string //姓名
	eat   int    //一次吃下雞腿的數量
	wait  int    //每次吃完後要休息的時間
	color int    //訊息顯示的顏色
}

func main() {
	//比賽結束訊息
	var done string

	//兩個大胃王
	archer := human{
		name:  "Archer",
		eat:   9,
		wait:  2,
		color: 31,
	}
	wesley := human{
		name:  "Wesley",
		eat:   4,
		wait:  1,
		color: 32,
	}

	//建立兩個通道
	channel_1 := make(chan string)
	channel_2 := make(chan string)

	//總共100隻雞腿
	food := 100

	//使用goroutine讓兩位大胃王開始吃雞腿
	go eat(archer, food, channel_1)
	go eat(wesley, food, channel_2)

	//使用select	等待通道傳回吃完的訊息
	select {
	case done = <-channel_1:
		fmt.Printf("\n\n%c[1;40;36m%s %c[0m\n", 0x1B, done, 0x1B)
	case done = <-channel_2:
		fmt.Printf("\n\n%c[1;40;36m%s %c[0m\n", 0x1B, done, 0x1B)
	}
}

func eat(x human, food int, c chan string) {
	//總共food隻雞腿 x每次只能吃eat隻
	for i := food; i > 0; i = i - x.eat {
		fmt.Printf("%c[0;40;%dm%s 剩下 %d 隻炸雞！%c[0m\n", 0x1B, x.color, x.name, i, 0x1B)
		//休息
		time_s := time.Duration(x.wait) * time.Second
		time.Sleep(time_s)
	}

	//將訊息傳入channel 通道
	c <- x.name + "吃完了 " + strconv.Itoa(food) + " 隻炸雞！"
}
