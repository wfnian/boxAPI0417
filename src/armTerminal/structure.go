package armTerminal

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


