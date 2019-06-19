package getSyncTask

import "log"
import "../StdJsonrpc"

type Create_source struct {
	Url    string            `json:"url"`
	Id     string            `json:"id"`
	Config StdJsonrpc.Config `json:"config"`
}

//Update_source update
type Update_source struct {
	Url    string            `json:"url"`
	Id     string            `json:"id"`
	Config StdJsonrpc.Config `json:"config"`
}

//Delete_source delete
type Delete_source struct {
	Id string `json:"id"`
}

type JsonrpcResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Id   string `json:"id"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	Id int `json:"id"`
}

type LockConfig struct {
	Gate Gate `json:"gate"`
}

type Gate struct {
	Extension int `json:"extension"`
	Cmd struct{
		Type      int    `json:"type"`
		Interval  int    `json:"interval"`
		Delay     int    `json:"delay"`
		Host      string `json:"host"`
		Port      int    `json:"port"`
		SuckCmd   string `json:"suckCmd"`
		SuckReply string `json:"suckReply"`
		ShutCmd   string `json:"shutCmd"`
		ShutReply string `json:"shutReply"`
	} `json:"cmd"`
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