package getSyncTask

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

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
		err = SetCollectors(response)
		HandleErr(err, 1, "err in set Collectors.")

		err = SetBox(response)
		HandleErr(err, 1, "err in set Box.")

		err = SetUsers(response)
		HandleErr(err, 1, "err in set Users.")
		// 上报同步任务结果

	}
}

func main() {
	GetSyncTask()
}
