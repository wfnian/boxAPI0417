package armTerminal

import (
	"../StdMsgForm"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"strings"
)

func matchUpload() {

	sess, _ := mgo.Dial("localhost:27017")

	defer sess.Close()
	c := sess.DB("wfnian").C("boxConfig")
	var box []StdMsgForm.BoxConfigInfo
	err := c.Find(bson.M{}).All(&box)
	HandleErr(err, 1, "Box 配置信息为空")

	//TODO srcId???
	verify := Str2Md5("srcId" + box[0].Verify)

	postData := UploadPostBody{
		CollectorId: "", //TODO ???
		VerifyCode:  verify,
		//TODO ???
	}
	url := "http://pass.deepdot.cn/deeppassEserver/" + "arm/match/upload"
	send, _ := json.Marshal(postData)
	resp, err := http.Post(url, "application/json", strings.NewReader(string(send)))
	HandleErr(err, 1, "Post Err")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	HandleErr(err, 0, "")
	var res StdMsgForm.Response
	err = json.Unmarshal([]byte(body), &res)

}

func Str2Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
