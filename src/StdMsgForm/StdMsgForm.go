package StdMsgForm

// Response 通用消息格式Response
type Response struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	Reference string `json:"reference"`
	Timstamp  string `json:"timstamp"`
	Results   Result `json:"results"`
}

type Result struct {
	ServerUrl   string          `json:"serverUrl"`
	SyncStates  int             `json:"syncStates"`
	BoxConfigs  []BoxConfigInfo `json:"boxConfigs"`
	Collectors  []CollectorInfo `json:"collectors"`
	Users       []UserInfo      `json:"users"`
	MatchResult MatchResult     `json:"matchResult"`
	Network     Network         `json:"network"`
}

//SyncResult 用户信息同步结果SyncResult
type SyncResult struct {
	Action  string `json:"action"`
	TaskId  string `json:"taskId"`
	Message string `json:"message"`
	Stat    int    `json:"stat"`
}

//用户信息UserInfo
type UserInfo struct {
	TaskId                  string        `json:"taskId"`
	Action                  string        `json:"action"`
	UserId                  string        `json:"userId"`
	UserName                string        `json:"userName"`
	FeatureSource           string        `json:"featureSource"`
	UserImgs                []UserImgInfo `json:"userImages"`
	Feature                 string        `json:"feature"`
	CollectorIds            []string      `json:"collectorIds"`
	PermissionCollectorType string        `json:"permissionCollectorType"`
	PermissionStartTime     string        `json:"permissionStartTime"`
	PermissionEndTime       string        `json:"permissionEndTime"`
	PermissionTimeType      string        `json:"permissionTimeType"`
	Message                 string        `json:"message"`
	CardId                  string        `json:"cardId"`
	IsNeedCard              int           `json:"isNeedCard"`
}

//用户图片信息UserImgInfo
type UserImgInfo struct {
	ImgId   string  `json:"imgId"`
	Action  string  `json:"action"`
	Quality float64 `json:"quality"`
	Url     string  `json:"url"`
}

//采集CollectorInfo
type CollectorInfo struct {
	TaskId          string          `json:"taskId"`
	Action          string          `json:"action"`
	CollectorId     string          `json:"collectorId"`
	CollectorType   string          `json:"collectorType"`
	SrcId           string          `json:"srcId"`
	CollectorName   string          `json:"collectorName"`
	LockConfig      LockConfig      `json:"lockConfig"`
	CollectorConfig CollectorConfig `json:"collectorConfig"`
}

//box配置信息BoxConfigInfo
type BoxConfigInfo struct {
	TaskId        string  `json:"taskId"`
	Action        string  `json:"action"`
	Verify        string  `json:"verify"`
	FirstPercent  float64 `json:"firstPercent"`
	SecondPercent float64 `json:"secondPercent"`
	ImgQuality    float64 `json:"imgQuality"`
}

type VisitInfo struct {
	Action              string  `json:"action"`
	FacetrakcId         string  `json:"facetrakcId"`
	UserId              string  `json:"userId"`
	Percent             float64 `json:"percent"`
	CollectorId         string  `json:"collectorId"`
	FacetrackCreateTime string  `json:"facetrackCreateTime"`
	Sex                 int     `json:"sex"`
	Age                 int     `json:"age"`
	Glasses             int     `json:"glasses"`
	IsPermitted         int     `json:"isPermitted"`
}

//box运行信息BoxRunningInfo
type BoxRunningInfo struct {
	Cpu             string `json:"cpu"`
	Tpu             string `json:"tpu"`
	CoreTemperature string `json:"coreTemperature"`
}

//ImgInfo
type ImgInfo struct {
	ImgName string  `json:"imgName"`
	Quality float64 `json:"quality"`
}

//MatchResult 匹配结果信息
type MatchResult struct {
	IsMatched   int     `json:"isMatched"`
	UserId      string  `json:"userId"`
	UserName    string  `json:"userName"`
	CardId      string  `json:"cardId"`
	Note        string  `json:"note"`
	Percent     float64 `json:"percent"`
	UserHeadImg string  `json:"userHeadImg"`
}

type LockConfig struct {
	Gate struct {
		Extension int `json:"extension"`
		Cmd       struct {
			Type      int    `json:"type"`
			Interval  int    `json:"interval"`
			Delay     int    `json:"delay"`
			Host      string `json:"host"`
			Port      int    `json:"port"`
			SuckCmd   string `json:"suckCmd"`
			SuckReply string `json:"suckReply"`
			ShutCmd   string `json:"shutCmd"`
			ShutReply string `json:"shutReply"`
		} `json:"cmd"`
	} `json:"gate"`
}

type CollectorConfig struct {
	Max_face       int  `json:"max_face"`
	Min_face       int  `json:"min_face"`
	Upload_display bool `json:"upload_display"`
	Display_width  int  `json:"display_width"`
	Area           struct {
		Top    int `json:"top"`
		Left   int `json:"left"`
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"area"`
	Url string `json:"url"`
}

type Network struct {
	Address string `json:"address"`
	Netmask string `json:"netmask"`
	Gateway string `json:"gateway"`
}
