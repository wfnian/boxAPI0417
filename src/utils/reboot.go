package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

//Reboot 重启机器 修改了本机IP等需要重启
func Reboot() {
	fmt.Println("reboot")
	shell := exec.Command("ls") //修改为重启 reboot
	stdout, err := shell.StdoutPipe()
	defer stdout.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err := shell.Start(); err != nil {
		log.Fatal(err)
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(opBytes))

}
