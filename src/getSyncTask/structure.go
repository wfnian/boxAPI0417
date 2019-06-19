package getSyncTask

import "log"
import "../StdJsonrpc"

type Create_source struct {
	Url    string            `json:"url"`
	Id     string            `json:"id"`
	Config StdJsonrpc.Config `json:"config"`
}
type Create_source_returns struct {
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

//Update_source update
type Update_source struct {
	Url    string            `json:"url"`
	Id     string            `json:"id"`
	Config StdJsonrpc.Config `json:"config"`
}
type Update_source_returns struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	Id int `json:"id"`
}

//Delete_source delete
type Delete_source struct {
	Id string `json:"id"`
}
type Delete_source_returns struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
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
	Cmd       struct {
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

//Create_db
type Create_db struct {
	Id     string `json:"id"`
	Volume int    `json:"volume"`
}
type Create_db_returns struct {
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

type Get_features struct {
	Images  []string `json:"images"`
	Retattr bool     `json:"retattr"`
}
type Get_features_returns struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Code     int                `json:"code"`
		Msg      string             `json:"msg"`
		Features []string           `json:"features"`
		Attrs    []StdJsonrpc.Faces `json:"attrs"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	Id int `json:"id"`
}

type Create_persons struct {
	Features [][]Features `json:"features"`
	Ids      []string     `json:"ids"`
	Db       string       `json:"db"`
}
type Create_persons_returns struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Code    int      `json:"code"`
		Msg     string   `json:"msg"`
		Persons []Person `json:"persons"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	Id int `json:"id"`
}

type Person struct {
	Id    int   `json:"id"`
	Faces []int `json:"faces"`
}
type Features struct {
	Feature string  `json:"feature"`
	Quality float64 `json:"quality"`
}

type Delete_person struct {
	Id string `json:"id"`
	Db string `json:"db"`
}
type Delete_person_returns struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"result"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	Id int `json:"id"`
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
