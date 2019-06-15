//package main

package heartbeat2

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"../utils"
)

type BoxRunningInfo struct {
	CPU             string `json:"cpu"`
	TPU             string `json:"tpu"`
	CoreTemperature string `json:"coreTemperature"`
}

type Response struct {
	Message  string  `json:"message"`
	Results  Results `json:"results"`
	Status   int     `json:"status"`
	Timstamp string  `json:"timstamp"`
}

type Results struct {
	SyncStates int `json:"syncStates"`
}

// 获取CPU和VPU的利用率，通过计算底层文件
func cpuVpuUsage() (string, string) {
	cpuFile := "/proc/stat"
	vpuFile := "/proc/vpuinfo"

	var cpuPercentage, vpuPercentage float64

	// start calculate cpu usage

	all2 := 0
	all1 := 0
	idle1 := 0
	contents, err := ioutil.ReadFile(cpuFile)
	if err != nil {
		log.Println(err)
	} else {

		firstline := strings.Fields(strings.Split(string(contents), "\n")[0])

		for i := 1; i < 8; i++ {
			temp, _ := strconv.Atoi(firstline[i])
			all1 += temp
		}
		idle1, _ = strconv.Atoi(firstline[4])

		time.Sleep(time.Duration(2) * time.Second)
	}

	contents, err = ioutil.ReadFile(cpuFile)
	if err != nil {
		log.Println(err)
	} else {

		firstline := strings.Fields(strings.Split(string(contents), "\n")[0])

		for i := 1; i < 8; i++ {
			temp, _ := strconv.Atoi(firstline[i])
			all2 += temp
		}
		idle2, _ := strconv.Atoi(firstline[4])

		cpuPercentage = float64(all2-all1-(idle2-idle1)) / float64(all2-all1) * 100
	}
	// end calculate cpu usage

	// start calculate vpu usage
	contents, err = ioutil.ReadFile(vpuFile)
	if err != nil {
		log.Println(err)

	} else {
		totalMemSize, _ := strconv.Atoi(strings.Fields(strings.Split(string(contents), ",")[0])[2])
		usedMemSize, _ := strconv.Atoi(strings.Fields(strings.Split(string(contents), ",")[1])[2])

		vpuPercentage = (float64(float64(usedMemSize) / float64(totalMemSize))) * 100
	}

	return strconv.FormatFloat(cpuPercentage, 'f', 6, 64), strconv.FormatFloat(vpuPercentage, 'f', 6, 64)
}

//HeartBeat 心跳
func HeartBeat() {

	UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	IP := utils.GetIP()

	//url := "http://" + IP + "/box/heartBeat?identifierId="
	url := "http://" + IP + "/box/heartBeat?identifierId="
	url = "http://pass.deepdot.cn/deeppassEserver" + "/box/heartBeat?identifierId=" + UUID
	log.Println(url)
	cpu, vpu := "12", "23" //cpuVpuUsage()
	// log.Println(cpu, vpu)
	post := BoxRunningInfo{
		CPU:             cpu,
		TPU:             vpu,
		CoreTemperature: "pass",
	}
	send, err := json.Marshal(post)
	log.Println(string(send))
	//向上位机报告自己的状态
	resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
	//resp, err := http.Get(url)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}
	log.Println(string(body))
	var response Response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		log.Panicln(err)
	}
	if response.Message != "成功!" {
		log.Panicln(err)
	} else if response.Results.SyncStates == 1 {
		//调用获取同步任务
		log.Println("调用获取同步任务")
	}

}

func _main() {
	HeartBeat()

}
