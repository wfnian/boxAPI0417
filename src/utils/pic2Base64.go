package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Pic2Base64() string {
	file, err := os.Open("2/home/wfnian/Pictures/2.png")
	if err != nil {
		log.Println("image to base64 err")
		return ""
	}
	defer file.Close()
	image, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("image to base64 err")
		return ""
	}
	imgBase64 := base64.StdEncoding.EncodeToString(image)
	fmt.Println(imgBase64)
	return imgBase64
}
