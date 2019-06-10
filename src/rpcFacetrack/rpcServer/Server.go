package main

import (
	"../../rpcFacetrack"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

func main() {
	//启动本地9821端口的服务
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	track := new(rpcFacetrack.Track)
	err:=server.RegisterService(track, "")
	if err != nil {
		log.Println(err)
	}
	r := mux.NewRouter()
	r.Handle("/", server)
	log.Println("JSON RPC service listen and serving on port 9821")
	if err := http.ListenAndServe(":9821", r); err != nil {
		log.Fatalf("Error serving: %s", err)
	}


}
