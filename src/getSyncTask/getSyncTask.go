package getSyncTask

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/pkg/errors"

	"../StdJsonrpc"
	"../StdMsgForm"
	"../utils"
)

//GetSyncTask 获取同步任务，box配置，视频源的配置，用户的配置。
func GetSyncTask() {
	UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	IP := utils.GetIP()

	url := "http://" + IP + "/box/getSyncTask?identifierId="
	url += UUID

	resp, err := http.Post(url, "application/json", strings.NewReader(""))

	HandleErr(err, 1, "post请求 获取上位机信息失败")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	HandleErr(err, 0, "")

	var response StdMsgForm.Response
	err = json.Unmarshal([]byte(body), &response)
	HandleErr(err, 1, "Json 解析失败")

	if response.Message != "成功!" {
		HandleErr(nil, 1, "失败")
	} else {
		// TODO 执行同步任务，配置box信息，配置collector信息
		err = setCollectors(response)
		HandleErr(err, 1, "err in set Collectors.")

		err = setBox(response)
		HandleErr(err, 1, "err in set Box.")

		err = setUsers(response)
		HandleErr(err, 1, "err in set Users.")
		// 上报同步任务结果

	}
}

func setBox(response StdMsgForm.Response) error {

	return nil
}

func setCollectors(response StdMsgForm.Response) error {
	var config StdJsonrpc.Config
	for i := 0; i < len(response.Results.Collectors); i++ {

		setConfig(&config, response.Results.Collectors[i])

		var postData StdJsonrpc.JsonrpcPost
		url := "http://localhost:7001"

		if response.Results.Collectors[i].CollectorType == "panel" {
			//若是面板机的情况
			var lock LockConfig
			if response.Results.Collectors[i].Action == "COLLECTOREDITLOCKCONFIG" {
				setlockConfig(&lock, response.Results.Collectors[i])
			}

		} else {

			switch response.Results.Collectors[i].Action {
			case "COLLECTORINSUSE": //采集启用
				{
					sess, err := mgo.Dial("localhost:27017")
					HandleErr(err, 1, "连接数据库出错")
					defer sess.Close()
					c := sess.DB("box").C("cameras")
					var conf StdMsgForm.CollectorInfo
					err = c.Find(bson.M{"collectorId": response.Results.Collectors[i].CollectorId}).One(&conf)
					HandleErr(err, 1, "数据库中没有 collectorId")

					postData.Method = "create_source"
					postData.Params = Create_source{
						Url:    conf.CollectorConfig.Url,
						Config: config,
						Id:     conf.CollectorId,
					}
					send, _ := json.Marshal(postData)
					resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
					HandleErr(err, 0, "http Post ERR")
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)

					HandleErr(err, 0, "")
					fmt.Println(string(body))
					var res JsonrpcResponse

					err = json.Unmarshal([]byte(body), &res)
					HandleErr(err, 0, "解析json出错")
					if res.Result.Code != 0 && res.Result.Msg != "SUCC" {
						HandleErr(nil, 0, "启用采集 ERROR")
						return errors.New("启用采集 ERROR")
					}
					return nil

				}
			case "COLLECTORDISUSE": //采集停用
				{
					/* 通过将视频源地址设置为空达到效果*/
					sess, err := mgo.Dial("localhost:27017")
					HandleErr(err, 1, "连接数据库出错")
					defer sess.Close()
					c := sess.DB("box").C("cameras")
					var conf StdMsgForm.CollectorInfo
					err = c.Find(bson.M{"collectorId": response.Results.Collectors[i].CollectorId}).One(&conf)
					HandleErr(err, 1, "数据库中没有 collectorId")

					postData.Method = "create_source"
					postData.Params = Create_source{
						Url:    "", //**************
						Config: config,
						Id:     conf.CollectorId,
					}
					send, _ := json.Marshal(postData)
					resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
					HandleErr(err, 0, "http Post ERR")
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)

					HandleErr(err, 0, "")
					fmt.Println(string(body))
					var res JsonrpcResponse

					err = json.Unmarshal([]byte(body), &res)
					HandleErr(err, 0, "解析json出错")
					if res.Result.Code != 0 && res.Result.Msg != "SUCC" {
						HandleErr(nil, 0, "启用采集 ERROR")
						return errors.New("启用采集 ERROR")
					}
					return nil
				}
			case "COLLECTORADD": //采集添加
				{
					postData.Method = "create_source"
					postData.Params = Create_source{
						Url:    response.Results.Collectors[i].CollectorConfig.Url,
						Config: config,
						Id:     response.Results.Collectors[i].CollectorId,
					}
					send, _ := json.Marshal(postData)
					resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
					HandleErr(err, 0, "http Post ERR")
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)

					HandleErr(err, 0, "")
					fmt.Println(string(body))
					var res JsonrpcResponse

					err = json.Unmarshal([]byte(body), &res)
					HandleErr(err, 0, "解析json出错")
					if res.Result.Code != 0 && res.Result.Msg != "SUCC" {
						HandleErr(nil, 0, "Creat_source ERROR")
						return errors.New("Creat_source ERROR")
					}
					return nil

				}
			case "COLLECTORREMOVE": //采集删除
				{
					postData.Method = "delete_source"
					postData.Params = Delete_source{
						Id: response.Results.Collectors[i].CollectorId,
					}
					send, _ := json.Marshal(postData)
					resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
					HandleErr(err, 0, "http Post ERR")
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)

					HandleErr(err, 0, "")
					fmt.Println(string(body))
					var res JsonrpcResponse

					err = json.Unmarshal([]byte(body), &res)
					HandleErr(err, 0, "解析json出错")
					if res.Result.Code != 0 && res.Result.Msg != "SUCC" {
						HandleErr(nil, 0, "Creat_source ERROR")
						return errors.New("Creat_source ERROR")
					}
					return nil

				}
			case "COLLECTOREDITCONFIG": //采集修改
				{
					postData.Method = "update_source"
					postData.Params = Update_source{
						Url:    response.Results.Collectors[i].CollectorConfig.Url,
						Config: config,
						Id:     response.Results.Collectors[i].CollectorId,
					}
					send, _ := json.Marshal(postData)
					resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
					HandleErr(err, 0, "http Post ERR")
					defer resp.Body.Close()
					body, err := ioutil.ReadAll(resp.Body)

					HandleErr(err, 0, "")
					fmt.Println(string(body))
					var res JsonrpcResponse

					err = json.Unmarshal([]byte(body), &res)
					HandleErr(err, 0, "解析json出错")
					if res.Result.Code != 0 && res.Result.Msg != "SUCC" {
						HandleErr(nil, 0, "Creat_source ERROR")
						return errors.New("Creat_source ERROR")
					}
					return nil

				}

			}

		}
	}

	return nil
}

func setUsers(response StdMsgForm.Response) error {

	return nil
}

func setConfig(config *StdJsonrpc.Config, response StdMsgForm.CollectorInfo) {
	config.Detect_interval = 5
	config.Track_interval = 1
	config.Sample_interval = 4
	config.Merge_threshold = 0.8
	config.Min_face_count = 2
	config.Max_tracker = 12
	config.Max_feature = 3
	config.Max_face = response.CollectorConfig.Max_face
	config.Min_face = response.CollectorConfig.Min_face
	config.Upload_display = response.CollectorConfig.Upload_display
	config.Display_width = response.CollectorConfig.Display_width
	config.Area = response.CollectorConfig.Area
}

func setlockConfig(lock *LockConfig, response StdMsgForm.CollectorInfo) {
	lock.Gate = response.LockConfig.Gate
}

func main() {
	GetSyncTask()
}
