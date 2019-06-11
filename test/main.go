package main

import (
	"../src/StdJsonrpc"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type source struct {
	Id string `json:"id"`
}

type JsonrpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Code    int      `json:"code"`
		Msg     string   `json:"msg"`
		Sources []source `json:"sources"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	Id int `json:"id"`
}

func main() {
	url := "http://222.199.233.18:7001"
	postData := StdJsonrpc.JsonrpcPost{
		Method:  "get_sources",
		Jsonrpc: "2.0",
		Params:  nil,
	}
	send, _ := json.Marshal(postData)
	resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(string(body))
	var res JsonrpcResponse

	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		log.Println(err)
	}
	log.Println(res.Result)

}
