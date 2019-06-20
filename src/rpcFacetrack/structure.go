package rpcFacetrack

import (
	"../StdJsonrpc"
	"log"
)

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
type Params struct {
	Id         string             `json:"id"`
	Source     string             `json:"source"`
	Faces      []StdJsonrpc.Faces `json:"faces"`
	Props      StdJsonrpc.Props   `json:"props"`
	Background string             `json:"background"`
	Features   []string           `json:"features"`
}

type Feature struct {
	Feature string  `json:"feature"`
	Quality float64 `json:"quality"`
}

type Search struct {
	Features [][]Feature `json:"features"`
	Top [2]int `json:"top"`
	Threshold [2]int `json:"threshold"`
	Db string `json:"db"`
}
type Search_returns struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Persons [][]man `json:"persons"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	Id int `json:"id"`
}

type man struct {
	Id string `json:"id"`
	Score int `json:"score"`
}

func HandleErr(err error, level int, msg string) {
	/*
		level 0 :警告
		level 1 :终止
	*/
	if err != nil {
		if level == 0 {
			log.Println(err, msg)
		} else {
			log.Panicln(err, msg)
		}
	}
}
