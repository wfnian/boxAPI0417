package heartbeat2

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)


func Init() {
	for {
		cpuFile := "/proc/stat"
		memFile := "/proc/meminfo"
		vpuFile := "/proc/vpuinfo"

		var cpuPercentage,vpuPercentage,memPercentage float64

		// start calculate cpu usage

		all2 := 0
		all1 := 0
		idle1:=0
		contents, err := ioutil.ReadFile(cpuFile)
		if err != nil {
			fmt.Println(err)
		}else {

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
		}else {

			firstline := strings.Fields(strings.Split(string(contents), "\n")[0])

			for i := 1; i < 8; i++ {
				temp, _ := strconv.Atoi(firstline[i])
				all2 += temp
			}
			idle2, _ := strconv.Atoi(firstline[4])

			cpuPercentage = float64(all2-all1-(idle2-idle1)) / float64(all2-all1) * 100
		}
		// end calculate cpu usage

		// start calculate memory usage
		contents, err = ioutil.ReadFile(memFile)
		if err != nil {
			fmt.Println(err)
		}else {
			total, _ := strconv.Atoi(strings.Fields(strings.Split(string(contents), "\n")[:3][0])[1])
			free, _ := strconv.Atoi(strings.Fields(strings.Split(string(contents), "\n")[:3][2])[1])
			memPercentage = float64(float64(total-free)/float64(total)) * 100
		}

		// end calculate memory usage

		// start calculate vpu usage
		contents, err = ioutil.ReadFile(vpuFile)
		if err != nil {
			fmt.Println(err)

		} else {
			totalMemSize, _ := strconv.Atoi(strings.Fields(strings.Split(string(contents), ",")[0])[2])
			usedMemSize, _ := strconv.Atoi(strings.Fields(strings.Split(string(contents), ",")[1])[2])

			vpuPercentage = (float64(float64(usedMemSize) / float64(totalMemSize))) * 100
		}
		fmt.Printf("Memory usage is %.3f%%\n", memPercentage)
		fmt.Printf("CPU usage is %.3f%%\n", cpuPercentage)
		fmt.Printf("VPU usage is %.3f%%\n", vpuPercentage)

	}
}
