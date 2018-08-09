package main

import (
	"net/http"
	"time"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type movieList []string
type Param map[string][]string

func main()  {
	//--------------查上映電影
	var list movieList
	body := getMovieList()
	_ = json.Unmarshal(body, &list)

	fmt.Printf("現在上映的電影：\n")
	for _, movie := range list{
		fmt.Printf("%s\n", movie)
	}
}

func getMovieList () []byte{
	client := http.Client{
		Timeout: time.Second * 30,
	}

	req, err := http.NewRequest("GET", "http://127.0.0.1:8888/get/movie", nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Host = "127.0.0.1:8888"


	res, getErr := client.Do(req)
	if getErr != nil {
		fmt.Println(err)
	}
	//fmt.Printf("%+v\n", res)
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(err)
	}

	return body
}
