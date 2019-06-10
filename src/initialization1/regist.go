//package initialization1

package main

import (
	"encoding/json"
	"fmt"
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

	for {
		url := "http://" + IP + "/box/regist?identifierId="
		url += UUID
		//url = "http://localhost:3000/object" //临时测试
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("url err", err)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		var ack StdMsgForm.Response
		err = json.Unmarshal([]byte(body), &ack)
		if err != nil {
			//panic(err)
			fmt.Println(err)
			continue
		}
		if ack.Status == 200 {
			fmt.Println("ok")
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
