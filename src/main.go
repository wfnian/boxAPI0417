// 主程序入口
package main

import (
	_ "./getConfig3"
	"./getSyncTask"
	_ "./initialization1"
	"./utils"
	"fmt"
)

// 初始化，包括本机唯一的ID

func main() {
	fmt.Println("main start")
	//heartbeat2.Init()
	//utils.Reboot()
	//initialization1.Regist()
	fmt.Println(utils.GetIP())
	//initialization1.Regist()
	//getconfig3.GetServerUrl()
	//heartbeat2.HeartBeat()
	fmt.Println("main end")
	getSyncTask.Main()
}
