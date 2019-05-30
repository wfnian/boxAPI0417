package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	rpcdemo "../../rpctest"
)

func main() {
	err := rpc.Register(rpcdemo.DemoService{})
	if err != nil {
		panic(err)
	}
	fmt.Println("OK")
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}

}

// {"jsonrpc":"2.0","method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
// {"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
