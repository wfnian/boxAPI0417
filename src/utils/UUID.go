package utils

import (
	"io/ioutil"
	"log"
)

//GetUUID 浏览/dev/disk/by-uuid/下的设备文件信息.
//以磁盘设备文件标识符相加得到机器标识符
func GetUUID() string {
	var UUID string
	files := "/dev/disk/by-uuid"
	dirlist, err := ioutil.ReadDir(files)
	if err != nil {
		log.Println(err)
		log.Println("In get UUID:read dir error")
	}

	for _, v := range dirlist {
		UUID += v.Name()
	}
	return UUID[:20]
}
