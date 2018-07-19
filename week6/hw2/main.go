package main

import (
	"archer"
	"fmt"
	"time"
	"strconv"
)

type Rocket struct {
	Name string
	Fuel int
}

func (r Rocket) Launch() {
	fmt.Println(r.Name + " 啟動發射程序")
	time.Sleep(1 * time.Second)

	for i := 10; i>=0; i-- {
		fmt.Println(r.Name + " 發射倒數 " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}

	fmt.Println(r.Name + " 發射!!!")
	time.Sleep(2 * time.Second)
}


func main()  {
	rocket1 := Rocket{Name:"洽洽號", Fuel:3}
	rocket2 := Rocket{Name:"DDD號", Fuel:3}
	archer.Start(rocket1)
	archer.Start(rocket2)
}
