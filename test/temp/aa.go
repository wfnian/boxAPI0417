package main

import (
	"fmt"
	"log"
)

type User struct {
	Name   string //这引号里面的就是tag
	Passwd string
}

func HandleErr(err error, level int, msg string) {
	/*
		level 0 :警告
		level 1 :终止
	*/
	if err != nil {
		if level == 0 {
			log.Println(err, msg)
		} else {
			log.Panicln(err, msg)
		}

	}
}

func _main() {
	//var a []User
	//a = append(a,User{"w","ww"})
	//a = append(a,User{"f","ff"})
	//fmt.Println(a)
	//j,err:=json.Marshal("_fda")
	//err=errors.New("啊啊啊")
	//HandleErr(err,1,"")
	//fmt.Println(string(j))
	s := "fdsafdsfafdsafadsfsad"
	fmt.Println(s[:10])
}
