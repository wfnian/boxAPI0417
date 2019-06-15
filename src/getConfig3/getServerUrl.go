package getconfig3

//package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../StdMsgForm"
	"../utils"
)

//GetConfig 从上位机获取网络配置
func GetServerUrl() {
	UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	IP := utils.GetIP()
	log.Println(IP)
	for {
		//url := "http://" + IP + "/box/getServerUrl?identifierId="
		url := "http://pass.deepdot.cn/deeppassEserver" + "/box/getServerUrl?identifierId="
		url += UUID
		log.Println(url)
		// 临时测试用
		// url = "http://localhost:3000/object"

		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		var ack StdMsgForm.Response
		err = json.Unmarshal([]byte(body), &ack)
		if err != nil {
			log.Println(err)
			continue
		}
		if ack.Status == 200 && ack.Message == "成功！" {
			//setting network 配置网络情况
			fmt.Println(ack.Results.Network)
			log.Println(ack.Results.ServerUrl)
			//networkSetting(ack.Results.Network)
			ConfigNetwork(ack.Results.Network.Address, ack.Results.Network.Netmask, ack.Results.Network.Gateway)
			break
		} else {
			continue
		}
	}
}

// func main() {
// 	GetConfig()
// }
