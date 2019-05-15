package utils

import (
	"fmt"
	"io/ioutil"
)

//GetUUID 浏览/dev/disk/by-uuid/下的设备文件信息.
//以磁盘设备文件标识符相加得到机器标识符
func GetUUID() string {
	var UUID string
	files := "/dev/disk/by-uuid"
	dirlist, e := ioutil.ReadDir(files)
	if e != nil {
		fmt.Println("read dir error")
	}

	for _, v := range dirlist {
		UUID += v.Name()
	}
	return UUID
}
