package armTerminal

import "log"

type UploadPostBody struct {
	CollectorId string   `json:"collectorId"`
	VerifyCode  string   `json:"verifyCode"`
	FacetrackId string   `json:"facetrackId"`
	IsNew       int      `json:"isNew"`
	Imgs        []string `json:"imgs"`
	Age         int      `json:"age"`
	Sex         int      `json:"sex"`
}
type FinishPostBody struct {
	CollectorId   string `json:"collectorId"`
	VerifyCode    string `json:"verifyCode"`
	FacetrackId   string `json:"facetrackId"`
	MatchState    int    `json:"matchState"`
	OpenGateState int    `json:"openGateState"`

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

