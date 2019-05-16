package getconfig3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../utils"
)

//ResponseNetwork 包含在下面ResponseResults
type ResponseNetwork struct {
	Address string
	Netmask string
	Gateway string
}

//ResponseResults 包含在下面ResponseGetServerUrlAndIP
type ResponseResults struct {
	ServerUrl string
	Network   ResponseNetwork
}

//ResponseGetServerUrlAndIP 获取上位机地址和box网络配置的应答body回应
type ResponseGetServerUrlAndIP struct {
	Message  string
	Results  ResponseResults
	Status   int
	Timstamp string
}

func GetConfig() {
	UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	IP := utils.GetIP()

	for {
		url := "http://" + IP + "/box/getServerUrl?identifierId="
		url += UUID
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("url err", err)
			continue
		}
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		var ack ResponseGetServerUrlAndIP
		err = json.Unmarshal([]byte(body), &ack)
		if err != nil {
			//panic(err)
			fmt.Println(err)
			continue
		}
		if ack.Message == "成功!" {
			//setting network
			fmt.Println(ack.Results.Network)
			break
		} else {
			continue
		}
	}
}
