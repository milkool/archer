package main

import (
	"fmt"
	"strconv"
	"./archerpackage"
)

/*
	程式入口
 */
func main() {
	//宣告變數a, b, c
	var a int = 1
	var b int32 = 2
	var c int64 = 3
	var d string = "999"
	var e float32 = 88.8
	var f float64 = 99.9

	//使用 := 可不用先宣告變數型態
	g := int32(a) + b
	//印出結果 %d 為印出int型態
	fmt.Printf("g = %d\n", g)
	h := int64(g) + c
	//印出結果 %d 為印出int型態
	fmt.Printf("h = %d\n", h)

	i := f/float64(e)
	//印出結果 %f 為印出float型態
	fmt.Printf("i = %f\n", i)


	//呼叫function
	o := a + strToInt(d)
	fmt.Printf("o = %d\n", o)


	//字串相接
	var x string = "Archer_"
	//型態轉換 int to string
	x += archerpackage.IntToStr(a)
	//印出結果 %s 為印出string型態
	fmt.Printf("d = %s\n", x)
}

/*
	小寫開頭的function為私有權限
	其他package的程式無法使用
 */
func strToInt(str string) int {
	//多個回傳值用, 分隔接收的變數
	//用_ 底線符號代表略過這個回傳值
	res, _ := strconv.Atoi(str)

	return res
}