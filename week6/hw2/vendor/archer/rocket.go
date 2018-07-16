package archer

import (
	"fmt"
	"time"
	"strconv"
)

type RocketInterface interface {
	Launch (rocket Rocket)
}

type Rocket struct {
	Name string
	Fuel int
}

func (i Rocket) Launch(rocket Rocket) {
	fmt.Println(rocket.Name + " 啟動發射程序")
	time.Sleep(1 * time.Second)

	for i := 10; i>=0; i-- {
		fmt.Println(rocket.Name + " 發射倒數 " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}

	fmt.Println(rocket.Name + " 發射!!!")
	time.Sleep(2 * time.Second)
}