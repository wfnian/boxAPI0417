package getSyncTask

import (
	"../StdJsonrpc"
	"../StdMsgForm"
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func SetUsers(response StdMsgForm.Response) error {
	sess, err := mgo.Dial("localhost:27017")
	HandleErr(err, 1, "连接数据库失败")
	defer sess.Close()
	dbName := "users"
	err = Creat_db(dbName)
	//创建一个名字为users的db在识别系统里面

	HandleErr(err, 1, "create db err in facebmapi-server")

	c := sess.DB("wfnian").C(dbName)

	for i := 0; i < len(response.Results.Users); i++ {
		switch response.Results.Users[i].Action {
		case "USERADD":
			{
				if response.Results.Users[i].FeatureSource == "img" {
					for j := 0; j < len(response.Results.Users[i].UserImgs); j++ {
						if response.Results.Users[i].UserImgs[j].Action == "ADD" {
							imageBase64 := convertPic2Base64(response.Results.Users[i].UserImgs[j].Url)
							var images []string
							images = append(images, imageBase64)
							res := GetFeatures(images)
							CreatePersons(dbName,response.Results.Users[i].UserId,res[0])
						} else if response.Results.Users[i].UserImgs[j].Action == "DELETE" {

						}
					}
				}
			}
		case "USEREDIT":
			{

			}
		case "USERDELETE":
			{

			}

		}

	}

	return nil
}

//创建一个名字为users的db在识别系统里面
func Creat_db(db_name string) error {
	postData := StdJsonrpc.JsonrpcPost{
		Jsonrpc: "2.0",
		Method:  "create_db",
		Params: Create_db{
			Id:     db_name,
			Volume: 20000,
		},
	}
	url := "http://localhost:7002"

	send, _ := json.Marshal(postData)
	resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
	HandleErr(err, 1, "http Post ERR")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	HandleErr(err, 0, "")
	var res Create_db_returns
	err = json.Unmarshal([]byte(body), &res)
	if res.Result.Msg != "SUCC" {
		return errors.New("create_db error")
	}

	return nil
}

func convertPic2Base64(url string) string {
	resp, err := http.Get(url)
	HandleErr(err, 1, "")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	HandleErr(err, 1, "")
	imgBase64 := base64.StdEncoding.EncodeToString(body)
	imgBase64 = "data:image/jpeg;base64," + imgBase64

	return imgBase64
}

func GetFeatures(imgs []string) []string {
	param := Get_features{
		Images:  imgs,
		Retattr: false,
	}
	postDate := StdJsonrpc.JsonrpcPost{
		Jsonrpc: "2.0",
		Params:  param,
		Method:  "get_features",
	}
	url := "http://localhost:7002"
	send, _ := json.Marshal(postDate)
	resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
	HandleErr(err, 1, "http post error")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	HandleErr(err, 0, "")

	var res Get_features_returns
	err = json.Unmarshal([]byte(body), &res)
	HandleErr(err, 1, "解析json出错")
	if res.Result.Code == 0 && res.Result.Msg == "SUCC" {
		return res.Result.Features
	}
	return nil

}

func CreatePersons(dbName string, featureid string, feature string) {
	var fID []string
	fID = append(fID, featureid)
	var f [][]Features
	f = append(f, nil)
	f[0] = append(f[0], Features{feature, 0.8})

	param := Create_persons{
		Db:       dbName,
		Ids:      fID,
		Features: f,
	}
	postDate:=StdJsonrpc.JsonrpcPost{
		Jsonrpc:"2.0",
		Method:"create_persons",
		Params:param,
	}
	url := "http://localhost:7002"
	send, _ := json.Marshal(postDate)
	resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
	HandleErr(err, 1, "http post error")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	HandleErr(err, 0, "")

	var res Create_source_returns
	err = json.Unmarshal([]byte(body), &res)
	HandleErr(err, 1, "解析json出错")
	if res.Result.Code == 0 && res.Result.Msg == "SUCC" {
		log.Println("Create persons ok!")
	}


}

func Main() {
	_ = Creat_db("a")
}
