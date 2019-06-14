package ackForSyncTask

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"../StdMsgForm"
	"../utils"
)

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

func AckForSyncTask(syncResult []StdMsgForm.SyncResult) error {
	UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	IP := utils.GetIP()

	url := "http://" + IP + "/box/ackForSyncTask?identifierId="
	url += UUID
	//url = "http://localhost:3000/object"
	send, err := json.Marshal(syncResult)
	//send = []SyncResult
	HandleErr(err, 0, "")

	resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
	//resp, err := http.Get(url)
	HandleErr(err, 1, "Post err")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	HandleErr(err, 1, "")
	var response StdMsgForm.Response
	err = json.Unmarshal([]byte(body), &response)
	HandleErr(err, 0, "Unmarshal err")
	if response.Message != "成功!" {
		HandleErr(errors.New(response.Message), 0, "上报同步任务结果失败")
	} else if response.Results.SyncStates == 1 {
		//调用获取同步任务
		fmt.Println("调用获取同步任务")

		//GetSyncTask()
		return nil
	}

	return nil
}
