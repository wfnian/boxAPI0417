package main

import (
	"../../rpctest"
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn,err:=net.Dial("tcp",":1234")
	if err!=nil{
		panic(err)
	}
	client:=jsonrpc.NewClient(conn)
	var res float64
	err=client.Call("DemoService.Div",rpcdemo.Args{10,4},&res)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(res,err)

}
