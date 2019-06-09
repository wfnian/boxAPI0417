package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../utils"
)

type Response2 struct {
	Message  string   `json:"message"`
	Results  Results2 `json:"results"`
	Status   int      `json:"status"`
	Timstamp string   `json:"timstamp"`
}

type Results2 struct {
	BoxConfigs []BoxConfigInfo `json:"box_configs"`
	Collectors []Collectors    `json:"collectors"`
	Users      []UserInfo      `json:"users"`
}

type BoxConfigInfo struct {
	TaskId        string  `json:"taskId"`
	Action        string  `json:"action"`
	Verify        string  `json:"verify"`
	FirstPercent  float64 `json:"firstPercent"`
	SecondPercent float64 `json:"secondPercent"`
	ImgQuality    float64 `json:"imgQuality"`
}

type Collectors struct {
	TaskId          string          `json:"taskId"`
	Action          string          `json:"action"`
	CollectorId     string          `json:"collectorId"`   //采集端ID
	CollectorType   string          `json:"collectorType"` //采集端LEIXING
	SrcId           string          `json:"srcId"`
	CollectorName   string          `json:"collectorName"`
	LockConfig      LockConfig      `json:"lockConfig"`
	CollectorConfig CollectorConfig `json:"collectorConfig"`
}

type LockConfig struct {
	Gate Gate `json:"gate"`
}

type Gate struct {
	Extension int `json:"extension"`
	Cmd       Cmd `json:"cmd"`
}

type Cmd struct {
	Type      int    `json:"type"`
	Interval  int    `json:"interval"`
	Delay     int    `json:"delay"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	SuckCmd   string `json:"suckCmd"`
	SuckReply string `json:"suckReply"`
	ShutCmd   string `json:"shutCmd"`
	ShutReply string `json:"shutReply"`
}

type CollectorConfig struct {
	Max_face       int  `json:"max_face"`
	Min_face       int  `json:"min_face"`
	Upload_display bool `json:"upload_display"`
	Display_width  int  `json:"display_width"`
	Area           struct {
		Top    int `json:"top"`
		Left   int `json:"left"`
		Width  int `json:"width"`
		Height int `json:"height"`
	}
	Url string `json:"url"`
}

type UserInfo struct {
	TaskId                  string        `json:"taskId"`
	Action                  string        `json:"action"`
	UserId                  string        `json:"userId"`
	UserName                string        `json:"userName"`
	FeatureSource           string        `json:"featureSource"`
	UserImgs                []UserImgInfo `json:"userImages"`
	Feature                 string        `json:"feature"`
	CollectorIds            []string      `json:"collectorIds"`
	PermissionCollectorType string        `json:"permissionCollectorType"`
	PermissionStartTime     string        `json:"permissionStartTime"`
	PermissionEndTime       string        `json:"permissionEndTime"`
	PermissionTimeType      string        `json:"permissionTimeType"`
	Message                 string        `json:"message"`
	CardId                  string        `json:"cardId"`
	IsNeedCard              int           `json:"isNeedCard"`
}

type UserImgInfo struct {
	ImgId   string  `json:"imgId"`
	Action  string  `json:"action"`
	Quality float64 `json:"quality"`
	Url     string  `json:"url"`
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
	fmt.Println(string(body))

	var response Response2
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Panicln(err)
	}
	if response.Message != "成功!" {
		log.Panicln(err)
	} else {
		//TODO 执行同步任务

		//上报同步任务结果

	}
}

func main() {
	GetSyncTask()
}
