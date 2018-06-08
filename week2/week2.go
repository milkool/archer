package main

import "fmt"

func main() {
	//練習1
	var a [5]int
	var t int

	a[0] = 87
	a[1] = 50
	a[2] = 88
	a[3] = 44
	a[4] = 100

	for _, v := range a {
		t = t + v
	}

	r := float32(t) / float32(len(a))

	fmt.Println("成績表：", a)
	fmt.Println("平均：", r)


	//練習2
	var min int
	x := [] int {
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	min = x[0]
	for _, y := range x {
		if y < min {
			min = y
		}
	}

	fmt.Println("最小數字：", min)
}