package StdJsonrpc

//JsonrpcPost std
type JsonrpcPost struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
}

//JsonrpcResponse std
type JsonrpcResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
	Id int `json:"id"`
}

type Faces struct {
	Aligned string  `json:"aligned"`
	Display string  `json:"display"`
	Quality float64 `json:"quality"`
}

type Props struct {
	Gender  int `json:"gender"`
	Age     int `json:"age"`
	Glasses int `json:"glasses"`
}

type DetectFaces struct {
	Aligned string `json:"aligned"`
	Box     struct {
		X int `json:"x"`
		Y int `json:"y"`
		W int `json:"w"`
		H int `json:"h"`
	} `json:"box"`
	Pts     []Point `json:"pts"`
	Score   float64 `json:"score"`
	Feature string  `json:"feature"`
	Attr    Props   `json:"attr"`
}
type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Config struct {
	Player          int     `json:"player"`
	Detect_interval int     `json:"detect_interval"`
	Track_interval  int     `json:"track_interval"`
	Sample_interval int     `json:"sample_interval"`
	Max_face        int     `json:"max_face"`
	Min_face        int     `json:"min_face"`
	Display_width   int     `json:"display_width"`
	Area            Area    `json:"area"`
	Upload_url      string  `json:"upload_url"`
	Merge_threshold float64 `json:"merge_threshold"`
	Min_face_count  int     `json:"min_face_count"`
	Max_tracker     int     `json:"max_tracker"`
	Max_feature     int     `json:"max_feature"`
	Upload_display  bool    `json:"upload_display"`
}

type Area struct {
	Top    int `json:"top"`
	Left   int `json:"left"`
	Width  int `json:"width"`
	Height int `json:"height"`
}
