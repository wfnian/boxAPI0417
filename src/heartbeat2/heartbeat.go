package main

//package heartbeat2

import (
	"fmt"
	"io/ioutil"
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
		fmt.Println(err)
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
		fmt.Println(err)
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
		fmt.Println(err)

	} else {
		totalMemSize, _ := strconv.Atoi(strings.Fields(strings.Split(string(contents), ",")[0])[2])
		usedMemSize, _ := strconv.Atoi(strings.Fields(strings.Split(string(contents), ",")[1])[2])

		vpuPercentage = (float64(float64(usedMemSize) / float64(totalMemSize))) * 100
	}

	// fmt.Printf("CPU usage is %.3f%%\n", cpuPercentage)
	// fmt.Printf("VPU usage is %.3f%%\n", vpuPercentage)

	return strconv.FormatFloat(cpuPercentage, 'f', 6, 64), strconv.FormatFloat(vpuPercentage, 'f', 6, 64)
}

//HeartBeat 心跳
func HeartBeat() {

	//UUID := utils.GetUUID() //获取本机唯一标识符，本机唯一标识符设置详见函数
	IP := utils.GetIP()
	fmt.Println(IP)

}

func main() {
	HeartBeat()
	
}
