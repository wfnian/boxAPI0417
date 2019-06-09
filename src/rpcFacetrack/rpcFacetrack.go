package rpcFacetrack

import (
	"../StdJsonrpc"
	"errors"
	"fmt"
	"net/http"
)

type Track struct {
}

type Params struct {
	Id         string             `json:"id"`
	Source     string             `json:"source"`
	Faces      []StdJsonrpc.Faces `json:"faces"`
	Props      StdJsonrpc.Props   `json:"props"`
	Background string             `json:"background"`
	Features   []string           `json:"features"`
}

type Returns struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (Track) Facetrack(r *http.Request, params *Params, result *Returns) error {
	if params.Id == "1" {
		var err error = errors.New("ID ERR")
		return err
	}
	*result = Returns{0, "SUCC"}
	fmt.Println(params)
	return nil
}

//func (t *Track) Multiply(r *http.Request, args *Params, result *int) error {
//	log.Println("Multiply %d with %d\n")
//	*result = 2
//	return nil
//}
