package zeroclient

import (
	"bytes"
	"encoding/gob"
)

type Resp struct {
	ReqIdentifier int32  `json:"reqIdentifier"`
	MsgIncr       string `json:"msgIncr"`
	ErrCode       int32  `json:"errCode"`
	ErrMsg        string `json:"errMsg"`
	Data          []byte `json:"data"`
}

type Req struct {
	ReqIdentifier int32  `json:"reqIdentifier" validate:"required"`
	Token         string `json:"token" `
	SendID        string `json:"sendID" validate:"required"`
	MsgIncr       string `json:"msgIncr" validate:"required"`
	Data          []byte `json:"data"`
}

func (r *Req) Gob() []byte {
	var buf = bytes.NewBuffer([]byte{})
	_ = gob.NewEncoder(buf).Encode(r)
	return buf.Bytes()
}
