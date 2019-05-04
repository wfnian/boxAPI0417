package initialization1

import (
	"fmt"
	"net"
	"strings"
)

func init() {
	//https://blog.csdn.net/linuxweiyh/article/details/78413275 Ubuntu 16.04下设置静态IP
	//获取本机IP地址
	conn, err := net.Dial("udp4", "baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])
}
