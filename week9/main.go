package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

type browser struct {
	client *http.Client
}

type Fg_debug struct {
	ID int `gorm:"column:id;type:int(11);NOT NULL;"`
	Content string `gorm:"column:content;type:longtext;"`
}

func (Fg_debug) TableName() string {
	return "fg_debug"
}

/*
	初始化url連線
 */
func newBrowser() *browser {
	hc := &browser{}
	hc.client = &http.Client{}

	return hc;
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
	取得網路天氣資料
 */
func getWeather() string {
	data := ""
	url := newBrowser()

	dataByte, _ := url.Get("http://weather.json.tw/api?region=taichung_city")
	data = string(dataByte)
	//fmt.Println(data, statusCode)

	return data
}

/*
	初始redis連線
 */
func initRedis(url string) redis.Conn {
	c, err := redis.Dial("tcp", url)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil
	}

	return c
}

/*
	寫入redis
 */
func setRedis(c redis.Conn, key string, value string) {
	_, err := c.Do("SET", key, value)
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

/*
	讀取redis
 */
func getRedis(c redis.Conn, key string) string {
	res, err := redis.String(c.Do("GET", key))
	if err != nil {
		fmt.Println("redis get failed:", err)
		return ""
	}

	return res
}

/*
	初始化DB連線
 */
func initDB() *gorm.DB {
	db, err := gorm.Open("mysql", "promise_dev:S7xepPk9yVXH@tcp(192.168.140.14:3306)/SLOT_TOOL")
	if err != nil {
		fmt.Println("failed to connect database:", err)
	}

	return db
}

/*
	DB寫入
 */
func setDB(db *gorm.DB, key string, value string) {
	var fg_debug Fg_debug
	db.Model(&fg_debug).Where("id = ?", key).Update("content", value)
}

/*
	DB讀取
 */
func getDB(db *gorm.DB, key string) string {
	var fg_debug Fg_debug
	db.First(&fg_debug, "id = ?", key)
	//fmt.Printf("%+v\n", fg_debug)

	return fg_debug.Content
}

func weather(c echo.Context) error {
	//Get data form redis
	//建立redis連線
	redis_c := initRedis("192.168.140.16:8188")
	res := getRedis(redis_c, "archerWeather")
	if res != "" {
		return c.String(http.StatusOK, res)
	}

	//If redis not exist, get data form DB
	//建立DB連線
	db_c := initDB()
	res = getDB(db_c, "1")
	if res != "" {
		return c.String(http.StatusOK, res)
	}

	//If DB not exist, call api form http://weather.json.tw/api and set to db and redis
	weatherJson := getWeather()
	setRedis(redis_c, "archerWeather", weatherJson)
	setDB(db_c, "1", weatherJson)

	return c.String(http.StatusOK, weatherJson)

	defer func(){
		redis_c.Close()
		db_c.Close()
	}()
	return nil
}

func main()  {
	e := echo.New()
	e.GET("/weather", weather)

	e.Logger.Fatal(e.Start(":1111"))
}
