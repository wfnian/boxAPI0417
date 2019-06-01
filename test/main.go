package main

import (
	"../src/StdMsgForm"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	var res StdMsgForm.Response
	fmt.Println(res.Message)

	data, err := ioutil.ReadFile("../config/collectorConfig.json")
	if err != nil {
		fmt.Println(err)
	}
	var collectorConfig StdMsgForm.CollectorConfig
	err = json.Unmarshal([]byte(data), &collectorConfig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(collectorConfig.Url)
}
