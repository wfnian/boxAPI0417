package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"../StdJsonrpc"
	"../StdMsgForm"
	"../utils"
)

type Create_source struct {
	Url    string            `json:"url"`
	Id     string            `json:"id"`
	Config StdJsonrpc.Config `json:"config"`
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

func GetSyncTask() {
	UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	IP := utils.GetIP()

	url := "http://" + IP + "/box/getSyncTask?identifierId="
	url += UUID
	url = "http://localhost:3000/object"
	//resp, err := http.Post(url, "application/json", strings.NewReader(""))
	resp, err := http.Get(url)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	//fmt.Println(string(body))

	var response StdMsgForm.Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Panicln(err)
	}
	if response.Message != "成功!" {
		log.Panicln(err)
	} else {
		// TODO 执行同步任务，配置box信息，配置collector信息
		err = setCollectors(response)
		if err != nil {
			log.Panicln(err)
		}

		err = setBox(response)
		if err != nil {
			log.Panicln(err)
		}

		err = setUsers(response)
		if err != nil {
			log.Panicln(err)
		}
		// 上报同步任务结果

	}
}

func setBox(response StdMsgForm.Response) error {

	return nil
}

func setCollectors(response StdMsgForm.Response) error {
	var config StdJsonrpc.Config
	for i := 0; i < len(response.Results.Collectors); i++ {
		config.Detect_interval = 5
		config.Track_interval = 1
		config.Sample_interval = 4
		config.Merge_threshold = 0.8
		config.Min_face_count = 2
		config.Max_tracker = 12
		config.Max_feature = 3
		config.Max_face = response.Results.Collectors[i].CollectorConfig.Max_face
		config.Min_face = response.Results.Collectors[i].CollectorConfig.Min_face
		config.Upload_display = response.Results.Collectors[i].CollectorConfig.Upload_display
		config.Display_width = response.Results.Collectors[i].CollectorConfig.Display_width
		config.Area = response.Results.Collectors[i].CollectorConfig.Area

		var postData StdJsonrpc.JsonrpcPost
		url := "http://localhost:7001"
		postData.Method = "create_source"
		postData.Params = Create_source{
			Url:    response.Results.Collectors[i].CollectorConfig.Url,
			Config: config,
		}
		send, _ := json.Marshal(postData)
		resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
		if err != nil {
			log.Println(err)
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
			log.Panicln(err)
		}
		if res.Result.Code != 0 && res.Result.Msg != "SUCC" {
			log.Panicln("Creat_source ERROR")
			return errors.New("Creat_source ERROR")
		}

	}

	return nil
}

func setUsers(response StdMsgForm.Response) error {

	return nil
}

func main() {
	GetSyncTask()
}
