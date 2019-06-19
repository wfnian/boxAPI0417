package getSyncTask

import (
	"../StdMsgForm"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func SetBox(response StdMsgForm.Response) error {

	sess, err := mgo.Dial("localhost:27017")
	HandleErr(err, 1, "连接数据库失败")
	defer sess.Close()
	c := sess.DB("wfnian").C("boxConfig")

	for i := 0; i < len(response.Results.BoxConfigs); i++ {
		if response.Results.BoxConfigs[i].Action == "IDENTIFIEREDIT" {
			//修改
			err = c.Update(bson.M{"verify": response.Results.BoxConfigs[i].Verify}, response.Results.BoxConfigs[i])
			if err != nil {
				//说明数据库中没有这个东西，改为插入
				err = c.Insert(response.Results.BoxConfigs[i])
				HandleErr(err, 1, "插入失败")
			}
		} else {
			err = c.Insert(response.Results.BoxConfigs[i])
			HandleErr(err, 1, "插入失败")
		}
	}


	return nil
}
