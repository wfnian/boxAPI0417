// 主程序入口
package main

import (
	"./heartbeat2"
	_ "./initialization1"
	"fmt"
)

func main() {
	fmt.Println("main")
	heartbeat2.Init()
}
