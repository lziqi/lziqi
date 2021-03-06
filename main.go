package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
)

type Weather struct {
	userId int
	id     int
	title  string
	body   string
}

func writeStringToFile(text string) {
	err := ioutil.WriteFile("README.md", []byte(text), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("start")

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	response, err := http.Get("https://restapi.amap.com/v3/weather/weatherInfo?city=420100&extensions=all&key=a2df40ded51af6989ac101865d0c516b")

	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		println(err)
	}

	city := gjson.Get(string(body), "forecasts.#.city").Array()[0]
	weather := gjson.Get(string(body), "forecasts.#.casts").Array()[0]

	txt := "# " + city.String() + "\n"
	txt += "|日期|白天天气|夜晚天气|白天温度|夜晚温度|" + "\n"
	txt += "|:--:|:--:|:--:|:--:|:--:|" + "\n"
	for _, day := range weather.Array() {
		date := gjson.Get(day.String(), "date").String()
		dayweather := gjson.Get(day.String(), "dayweather").String()
		nightweather := gjson.Get(day.String(), "nightweather").String()
		daytemp := gjson.Get(day.String(), "daytemp").String()
		nighttemp := gjson.Get(day.String(), "nighttemp").String()
		txt += "|" + date + "|" + dayweather + "|" + nightweather + "|" + daytemp + "℃|" + nighttemp + "℃|\n"
	}
	txt += " \n"
	txt += "[物联网学习指南](http://doc.lziqi.top/IoT)\n"
	writeStringToFile(txt)

	//https://jsonplaceholder.typicode.com/posts/1
	//https://restapi.amap.com/v3/weather/weatherInfo?city=420100&extensions=all&key=a2df40ded51af6989ac101865d0c516b
}
