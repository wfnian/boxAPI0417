package main

import (
	"../../src/StdMsgForm"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type s struct {
	A int
	B string
}

func main() {
	//var a [][]s
	//
	//var c[]s
	//c = append(c,s{2,"4"})
	//c = append(c,s{3,"a"})
	//fmt.Println(c)
	//
	//a=append(a,c)
	//fmt.Println(a)
	sess, _ := mgo.Dial("localhost:27017")

	defer sess.Close()
	c := sess.DB("wfnian").C("boxConfig")

	//config:=StdMsgForm.BoxConfigInfo{
	//	TaskId:"2",
	//	Action:"IDENTIFIEREDIT",
	//	Verify:"2",
	//	FirstPercent:0.8,
	//	SecondPercent:0.7,
	//	ImgQuality:0.9,
	//}
	//_ = c.Insert(config)
	var box []StdMsgForm.BoxConfigInfo
	err := c.Find(bson.M{}).All(&box)
	fmt.Println(err)
	fmt.Println(box)
	fmt.Println(time.RFC3339)
	now := time.Now().Format("2006-01-02 15:04:05")

}
