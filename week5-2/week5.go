package main

import (
	"fmt"
	"time"
)

func main()  {
	//設定限制數值
	const limit int = 20
	//記錄開始時間
	start := time.Now()

	//建立通道與buff
	c := make(chan int64, limit)
	//用goroutine運算各階乘
	//for i := 1; i<=limit; i++ {
	//	go factorial(i, c)
	//}

	go factorial2(int64(20), c)

	var anser int64
	var res int64
	//接收階乘結果並相加
	for i:=0; i<limit; i++{
		res = <- c
		anser = anser + res
	}

	//紀錄消耗時間
	end := time.Since(start)
	fmt.Println(end)
	fmt.Println(anser)
}

func factorial(n int, c chan int64)  {
	var res int64 = 1
	for i := n; i > 0; i -- {
		res = res * int64(i)
	}
	//fmt.Printf("%d!= %d\n", n, res)
	c <- res
}


func factorial2(limit int64, c chan int64)  {
	var res int64 = 1
	var i int64
	for i = 1; i <= limit; i ++ {
		res = res * int64(i)
		c <- res
		//fmt.Printf("%d!= %d\n", i, res)
	}
	//c <- -1
}