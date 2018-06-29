package main

import (
	"time"
	"fmt"
	"sync"
)

type human struct {
	name  string //姓名
	power   int    //力量 每次能砍的數目
	wait  int    //每次工作的時間
	color int    //訊息顯示的顏色
}

func main() {
	//兩個伐木工
	archer := human{
		name: "Archer",
		power:  10,
		wait: 3,
		color: 31,
	}
	wesley := human{
		name: "Wesley",
		power:  3,
		wait: 1,
		color: 32,
	}

	//建立WaitGroup
	var wg sync.WaitGroup
	//建立互斥鎖
	var mutex sync.Mutex
	//總共100跟樹木
	tree := 100

	//WaitGroup +2
	wg.Add(2)
	//使用goroutine讓兩位伐木工開始伐木
	go cutTree(archer, &tree, &wg, mutex)
	go cutTree(wesley, &tree, &wg, mutex)

	//等待所有工作完成
	wg.Wait()

	//所有工作都完成後才會繼續執行程式碼
	fmt.Printf("%c[0;40;36m已砍完所有樹木%c[0m\n", 0x1B, 0x1B)


}

func cutTree(x human, tree *int, wg *sync.WaitGroup, mutex sync.Mutex) {
	var cut int

	//在砍完所有樹木前不會停止工作
	for i := *tree; i>0; i = *tree {
		//樹木數量上鎖
		mutex.Lock()

		if *tree >= x.power {
			*tree = *tree-x.power
			cut = x.power
		} else {
			cut = *tree
			*tree = 0
		}
		fmt.Printf("%c[0;40;%dm%s 正在砍 %d 棵樹，剩下 %d 顆樹！%c[0m\n", 0x1B, x.color, x.name, cut, *tree, 0x1B)

		//樹木數量解鎖
		mutex.Unlock()

		time_s := time.Duration(x.wait) * time.Second
		time.Sleep(time_s)
	}

	//工作完成 回報WaitGroup -1
	wg.Done()
}
