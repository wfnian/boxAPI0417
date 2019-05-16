// 主程序入口
package main

import (
	"fmt"

	"./initialization1"
	"./utils"
)

// 初始化，包括本机唯一的ID

func main() {
	fmt.Println("main start")
	//heartbeat2.Init()
	//utils.Reboot()
	//initialization1.InitSetting()
	fmt.Println(utils.GetIP())
	initialization1.InitSetting()
	fmt.Println("main end")

}
