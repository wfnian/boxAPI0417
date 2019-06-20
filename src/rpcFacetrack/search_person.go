package rpcFacetrack

import (
	"../StdJsonrpc"
	"../StdMsgForm"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Search_person(tracked JsonrpcPost) {
	/*
		" params ": {
			"id": "",
			"source ": "",
			"faces ": [{}] ,
			"props ": {},
			" background ": "",
			" features ": [""]
		},
	*/
	// TODO
	var _feature [][]Feature

	var temp []Feature
	for i:=0;i<len(tracked.Params.Features);i++{
		temp = append(temp,Feature{tracked.Params.Features[i],0.8})
	}


	_feature=append(_feature, temp)
	param := Search{
		Top:       [2]int{1000, 10},
		Threshold: [2]int{70, 20},
		Db:        "wfnian",
		Features:  _feature,
	}

	postData:=StdJsonrpc.JsonrpcPost{
		Jsonrpc:"2.0",
		Params:param,
		Method:"search",
	}
	url:="http://localhost:7002"
	send,_:=json.Marshal(postData)
	resp,err:=http.Post(url,"application/json",strings.NewReader(string(send)))
	HandleErr(err,1,"http post error")
	defer resp.Body.Close()
	body,err:=ioutil.ReadAll(resp.Body)
	HandleErr(err,0,"")
	var res Search_returns
	err = json.Unmarshal([]byte(body),&res)
	HandleErr(err,0,"解析JSON出错")


	sess, err := mgo.Dial("localhost:27017")
	HandleErr(err, 1, "连接数据库失败")
	defer sess.Close()
	c := sess.DB("wfnian").C("boxConfig")
	var pass StdMsgForm.BoxConfigInfo
	err = c.Find(bson.M{}).One(&pass)
	HandleErr(err,1,"box配置信息为空")

	var visit StdMsgForm.VisitInfo
	var match StdMsgForm.MatchResult
	if float64(res.Result.Persons[0][0].Score/100)>pass.SecondPercent{
		//TODO 通过
		visit=StdMsgForm.VisitInfo{
			UserId:res.Result.Persons[0][0].Id,
			IsPermitted:0,
			FacetrackCreateTime:time.Now().Format("2006-01-02 15:04:05"),
		}
		match=StdMsgForm.MatchResult{
			IsMatched:0,
			UserId:res.Result.Persons[0][0].Id,
		}
	}else {
		visit=StdMsgForm.VisitInfo{
			UserId:res.Result.Persons[0][0].Id,
			IsPermitted:1,
			FacetrackCreateTime:time.Now().Format("2006-01-02 15:04:05"),
		}
		match=StdMsgForm.MatchResult{
			IsMatched:1,
			UserId:res.Result.Persons[0][0].Id,
		}
	}
	v :=sess.DB("wfnian").C("visitInfo")
	err = v.Insert(visit)
	HandleErr(err,0,"insert warning")
	m :=sess.DB("wfnian").C("matchResult")
	err = m.Insert(match)
	HandleErr(err,0,"insert warning")


}
