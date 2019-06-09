package main

import (
	"../../json-rpc"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

func main() {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	arith := new(json_rpc.Arith)
	_ = server.RegisterService(arith, "")
	r := mux.NewRouter()
	r.Handle("/", server)
	log.Println("JSON RPC service listen and serving on port 1234")
	if err := http.ListenAndServe(":1234", r); err != nil {
		log.Fatalf("Error serving: %s", err)
	}
}