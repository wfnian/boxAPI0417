package initialization1

//package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../StdMsgForm"
	"../utils"
)

//Regist 第一个需求：初始化
func Regist() {

	UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	IP := utils.GetIP()
	log.Println(IP)

	for {
		//url := "http://" + IP + "/box/regist?identifierId="
		url := "http://pass.deepdot.cn/deeppassEserver" + "/box/regist?identifierId="
		url += UUID
		log.Println(url)
		//url = "http://localhost:3000/object" //临时测试
		resp, err := http.Get(url)
		if err != nil {
			log.Println("url err", err)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		var ack StdMsgForm.Response
		err = json.Unmarshal([]byte(body), &ack)
		if err != nil {
			//panic(err)
			log.Println(err)
			continue
		}
		if ack.Status == 200 && ack.Message == "成功！" {
			log.Println("ok")
			break
		} else {
			log.Println("Error:", ack.Status, "retry")
			continue
		}
	}

}

func main() {
	Regist()
}
