package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"../utils"
)

//ConfigNetwork 配置网络参数
func ConfigNetwork(address, netmask, gateway string) {

	//指定配置文件的路径 一般在 /etc/network/interfaces.d 下
	configFile := "/home/wfnian/project/boxAPI0417/networkConfig"
	file, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	//fmt.Println(string(content))
	raw := strings.Split(string(content), "\n")
	var change []string
	var isChanged bool
	for _, c := range raw {
		isIn := strings.Index(c, "address")
		if isIn != -1 { //这一行里是address 配置相关
			old := c[8:] //address 后面的具体地址
			old = strings.Trim(old, " ")
			if old != address {
				c = strings.Replace(c, old, address, -1)
				isChanged = true
			}

		}
		isIn = strings.Index(c, "netmask")
		if isIn != -1 {
			old := c[8:]
			old = strings.Trim(old, " ")
			if old != netmask {
				c = strings.Replace(c, old, netmask, -1)
				isChanged = true
			}
		}
		isIn = strings.Index(c, "gateway")
		if isIn != -1 {
			old := c[8:]
			old = strings.Trim(old, " ")
			if old != gateway {
				c = strings.Replace(c, old, gateway, -1)
				isChanged = true
			}
		}
		change = append(change, c)
	}
	if isChanged {
		fmt.Println(strings.Join(change, "\n"))
		err = ioutil.WriteFile(configFile, []byte(strings.Join(change, "\n")), 0666)
		// 修改后需要重启
		utils.Reboot()
	}

}

func main() {
	address, netmask, gateway := "192.168.1.109", "255.255.255.0", "192.168.1.1"
	ConfigNetwork(address, netmask, gateway)
}
