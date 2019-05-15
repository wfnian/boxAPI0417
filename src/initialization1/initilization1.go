package initialization1
//package main

import (
	"fmt"

	"../utils"
)

func InitSetting() {
	fmt.Println()
	UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	fmt.Println(UUID)

}

func main() {
	InitSetting()
}
