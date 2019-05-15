// 主程序入口
package main

import (
	"fmt"

	"./initialization1"
	// "./utils"
)

// 初始化，包括本机唯一的ID

func main() {
	fmt.Println("main")
	//heartbeat2.Init()
	//utils.Reboot()
	initialization1.InitSetting()

}
