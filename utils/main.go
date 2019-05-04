package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("udp4", "baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	fmt.Println(strings.Split(conn.LocalAddr().String(), ":")[0])

	addrs, err := net.InterfaceAddrs()
    if err != nil{
        fmt.Println(err)
        return
    }
    for _, value := range addrs{
        if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback(){
            if ipnet.IP.To4() != nil{
				fmt.Println(ipnet.IP.String())
				break
            }
        }
    }
}
