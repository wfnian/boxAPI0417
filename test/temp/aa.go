package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	url:="https://pic3.zhimg.com/v2-6391ef8cc345fbd7d7d30a6e3d22aa21_xl.jpg"
	resp,err:=http.Get(url)
	if err!=nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T",(body))
	f,err:=os.Create("a.jpg")
	_, _ = f.Write(body)
	imgBase64 := base64.StdEncoding.EncodeToString(body)
	fmt.Println(imgBase64)
}