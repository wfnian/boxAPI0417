package rpcFacetrack

import (
	"../StdJsonrpc"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

type Params struct {
	Id         string             `json:"id"`
	Source     string             `json:"source"`
	Faces      []StdJsonrpc.Faces `json:"faces"`
	Props      StdJsonrpc.Props   `json:"props"`
	Background string             `json:"background"`
	Features   []string           `json:"features"`
}

type Returns struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type JsonrpcResponseSuccess struct {
	Jsonrpc string  `json:"jsonrpc"`
	Result  Returns `json:"result"`
	Id      int     `json:"id"`
}

type JsonrpcPost struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
	Id      int    `json:"id"`
}

func facetrack(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		log.Fatalln(err)
	}
	var tracked JsonrpcPost
	err = json.Unmarshal([]byte(data), &tracked)

	//log.Println(tracked)
	log.Println(tracked.Id)
	log.Println(tracked.Params.Source)
	log.Println(tracked.Params.Props)
	log.Println(tracked.Params.Faces)
	//log.Println(tracked.Params)

	c.JSON(200, JsonrpcResponseSuccess{
		Jsonrpc: "2.0",
		Id:      2,
		Result:  Returns{Code: 0, Msg: "SUCC"},
	})
}

func Track() error {
	router := gin.Default()

	router.POST("", facetrack)

	err := router.Run(":9821")
	if err != nil {
		log.Panicln(err)
	}
	return nil
}

//func (t *Track) Multiply(r *http.Request, args *Params, result *int) error {
//	log.Println("Multiply %d with %d\n")
//	*result = 2
//	return nil
//}
