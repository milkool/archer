package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"strconv"
	"time"
)


type browser struct {
	client *http.Client
}

/*
	get url data
 */
func (self *browser) Get(requestUrl string) ([]byte, int) {
	request,_ := http.NewRequest("GET", requestUrl, nil)
	response,_ := self.client.Do(request)
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)

	return data, response.StatusCode
}

/*
	初始化url連線
 */
func newBrowser() *browser {
	hc := &browser{}
	hc.client = &http.Client{}

	return hc;
}

func getUrl(url string, c chan time.Duration) {
	//data := ""
	browser := newBrowser()

	//記錄開始時間
	start_s := time.Now()
	//dataByte, _ := browser.Get(url)
	_ , _ = browser.Get(url)
	//紀錄全域結束時間
	end := time.Since(start_s)

	//data = string(dataByte)
	//fmt.Println(data)

	c <- end
}

func main()  {
	var url string
	var input string
	var go_r int
	var time_recoder []time.Duration
	var max float64 = 0
	var min float64 = 0
	var total float64 = 0

	//輸入遊戲ID
	fmt.Printf("請輸入目標API: ")
	fmt.Scanln(&url)

	//輸入併發數量
	fmt.Printf("請輸入併發數量: ")
	fmt.Scanln(&input)
	go_r, _ = strconv.Atoi(input)

	//開始連線
	//記錄全域開始時間
	start := time.Now()
	//建立通道與buff
	c := make(chan time.Duration, go_r)
	for i := 0; i < go_r; i++ {
		go getUrl(url, c)
	}

	//接收結果
	for i := 0; i < go_r; i++ {
		res := <- c
		time_recoder = append(time_recoder, res)
	}

	//紀錄全域結束時間
	end := time.Since(start)

	//單一連線花費時間
	for k, v := range time_recoder {
		sec := v.Seconds()
		if sec > max {
			max = sec
		} else if sec < min {
			min = sec
		}
		if min == 0 {
			min = sec
		}
		fmt.Printf("第 %3d 次花費時間: %f 秒\n", k+1, sec)
	}
	fmt.Printf("單一最大花費時間: %f 秒\n", max)
	fmt.Printf("單一最小花費時間: %f 秒\n", min)
	end_sec := end.Seconds()
	fmt.Printf("壓測總數花費時間: %f 秒\n", end_sec)

	//平均花費時間
	for _, v := range time_recoder {
		total += v.Seconds()
	}

	fmt.Printf("平均花費時間: %f 秒\n", total/float64(len(time_recoder)))

}
