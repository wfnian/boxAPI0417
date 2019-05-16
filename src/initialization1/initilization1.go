package initialization1

//package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../utils"
)

//ResponseRegist 注册时候的应答body回应
type ResponseRegist struct {
	Message  string
	Status   int
	Timstamp string
}

//InitSetting 第一个需求：初始化
func InitSetting() {

	UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	IP := utils.GetIP()

	for {
		url := "http://" + IP + "/box/regist?identifierId="
		url += UUID
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("url err", err)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		var ack ResponseRegist
		err = json.Unmarshal([]byte(body), &ack)
		if err != nil {
			//panic(err)
			fmt.Println(err)
			continue
		}
		if ack.Message == "成功!" {
			fmt.Println("ok")
			break
		} else {
			fmt.Println("不成功")
			continue
		}
	}


}

func main() {
	InitSetting()
}
