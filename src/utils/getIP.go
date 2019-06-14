package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GetIP() string {
	file, ferr := os.OpenFile("PCIP.config", os.O_RDWR|os.O_CREATE, 0755)
	defer file.Close()
	if ferr != nil {
		log.Println("读取上位机文件错误", ferr)
		panic(ferr)
	}
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	IP := strings.Split(string(contents), "\n")

	if len(IP) == 0 { //说明配置文件为空的情况下
		return "192.168.1.250:8080" //设置为默认
	} else if IP[0] == "" {
		return "192.168.1.250:8080" //设置为默认
	} else {
		return IP[0]
	}
}
