package main

import (
	"time"
	"fmt"
	"sync"
)

//伐木廠
type cutTree struct {
	cut int //每秒生產紙漿 公斤
}

//造紙廠
type makePaper struct {
	span int //每秒消耗紙漿 公斤
	make int //每秒生產紙張
}

//印刷廠
type printer struct {
	print int //每秒印出紙張
}

func main()  {
	//建立WaitGroup
	var wg sync.WaitGroup

	//伐木廠
	cutTree := cutTree{
		cut: 1000,
	}
	//造紙廠1
	makePaper1 := makePaper{
		make: 500,
	}
	//造紙廠2
	makePaper2 := makePaper{
		make: 300,
	}
	//印刷廠
	printer := printer{
		print: 6000,
	}

	//伐木廠 -> 造紙廠
	c1 := make(chan int, 100)
	//造紙廠 -> 印刷廠
	c2 := make(chan int, 100)

	go cutTreeFunc(cutTree, c1)
	go makePaperFunc(makePaper1, c1, c2)
	go makePaperFunc(makePaper2, c1, c2)
	wg.Add(1)
	go printerFunc(printer, c2, &wg)
	wg.Wait()
	fmt.Println("DONE!\n")
}

func cutTreeFunc(maker cutTree, c chan int)  {
	for ; ;  {
		//每秒放入紙漿
		c <- maker.cut
		fmt.Printf("生產%d公斤紙漿\n", maker.cut)
		time.Sleep(1 * time.Second)
	}
}

func makePaperFunc(maker makePaper, c1 chan int, c2 chan int) {
	//紙漿庫
	var tree int
	var get int

	for ; ; {
		//判斷紙漿庫存夠不夠
		if tree < maker.span {
			select {
			case get = <- c1 :
				tree += get
			}
		}
		tree = tree - maker.span
		c2 <- maker.make
		fmt.Printf("生產%d張紙張\n", maker.make)
		time.Sleep(1 * time.Second)
	}
}
func printerFunc(maker printer, c chan int, wg *sync.WaitGroup) {
	//紙張庫存
	var paper int
	var get int
	//已製造紙張
	var print_paper int

	for ; print_paper < 60000; {
		//判斷紙庫存夠不夠
		if paper < maker.print {
			select {
			case get = <- c:
				paper += get
			}
		}
		paper = paper - maker.print

		print_paper += maker.print
		fmt.Printf("印刷%d張紙\n", maker.print)
		fmt.Printf("----已印刷%d張紙\n", print_paper)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("----已印刷%d張紙\n", print_paper)

	wg.Done()
}