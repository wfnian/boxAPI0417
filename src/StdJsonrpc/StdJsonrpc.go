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
	Result  interface{} `json:"method"`
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
