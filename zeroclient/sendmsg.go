package zeroclient

import (
	"github.com/gorilla/websocket"
	"github.com/showurl/Zero-IM-Client-Go/pb"
	"github.com/showurl/Zero-IM-Client-Go/utils"
	"time"
)

func (user *Client) sendMsg(req *Req) error {
	if user.wsClient == nil {
		time.Sleep(time.Second)
		return user.sendMsg(req)
	}
	user.lock.Lock()
	err := user.wsClient.WriteMessage(websocket.BinaryMessage, req.Gob())
	user.lock.Unlock()
	return err
}

func (user *Client) SendMsg(data *pb.MsgData) error {
	utils.PrintMsg(data, user.UserID)
	return user.sendMsg(&Req{
		ReqIdentifier: pb.WSSendMsg.Int32(),
		Token:         user.Token,
		SendID:        user.UserID,
		MsgIncr:       "unknown",
		Data:          data.Bytes(),
	})
}
