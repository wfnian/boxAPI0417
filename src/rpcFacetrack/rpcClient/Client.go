package main

import (
	"../../rpcFacetrack"
	"bytes"
	"fmt"
	"github.com/gorilla/rpc/json"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:9821"
	params := rpcFacetrack.Params{
		Id:     "2",
		Source: "w",
	}
	message, err := json.EncodeClientRequest("Track.Facetrack", params)
	if err != nil {
		log.Println(err)
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(message))
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	var res rpcFacetrack.Returns
	err = json.DecodeClientResponse(resp.Body, &res)
	fmt.Println(res)

}
