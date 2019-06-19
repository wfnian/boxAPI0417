package main

import (
	"encoding/json"
	"fmt"
	"log"

	"../../src/StdMsgForm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Returns struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func main() {
	sess, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Panicln(err)
	}
	defer sess.Close()
	c := sess.DB("wfnian").C("a")
	res := Returns{3, "哈哈"}

	err = c.Insert(res)
	if err != nil {
		log.Println(err)
	}
	var v []Returns
	err = c.Find(bson.M{"code": 3}).All(&v)
	fmt.Println(v)
	a := StdMsgForm.BoxRunningInfo{"1", "2", "3"}
	s, _ := json.Marshal(a)
	fmt.Println(string(s))
}
