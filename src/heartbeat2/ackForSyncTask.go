package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../utils"
)

type SyncResult struct {
	Action  string `json:"action"`
	TaskId  string `json:"taskId"`
	Message string `json:"message"`
	Stat    int    `json:"stat"`
}

func AckForSyncTask() error {
	UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	IP := utils.GetIP()

	url := "http://" + IP + "/box/ackForSyncTask?identifierId="
	url += UUID
	url = "http://localhost:3000/object"

	//resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
	resp, err := http.Get(url)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	var response Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Panicln(err)
	}
	if response.Message != "成功!" {
		log.Panicln(err)
	} else if response.Results.SyncStates == 1 {
		//调用获取同步任务
		fmt.Println("调用获取同步任务")
		
		//GetSyncTask()
		return nil
	}

	return nil
}
