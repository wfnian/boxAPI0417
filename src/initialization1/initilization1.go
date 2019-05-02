package initialization1

import (
	"fmt"
	"net"
)

func init() {
	//https://blog.csdn.net/linuxweiyh/article/details/78413275 Ubuntu 16.04下设置静态IP
	//获取本机IP地址
	var pcid string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				pcid = ipnet.IP.String()
			}
		}
	}

	fmt.Println(pcid)
	for { //

		break
	}

}
