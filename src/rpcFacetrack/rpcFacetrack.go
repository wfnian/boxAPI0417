package rpcFacetrack

import (
	_ "../StdJsonrpc"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)



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
	log.Println(tracked.Params.Faces) //list [{}]
	log.Println(tracked.Params.Features)
	//log.Println(tracked.Params)
	go Search_person(tracked)

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
	HandleErr(err, 1, "端口9821启动失败")
	return nil
}
