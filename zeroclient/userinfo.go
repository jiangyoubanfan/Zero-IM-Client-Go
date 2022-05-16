package zeroclient

import (
	"github.com/gorilla/websocket"
	"github.com/showurl/Zero-IM-Client-Go/pb"
	"sync"
	"time"
)

type Client struct {
	UserID     string
	Nickname   string
	Avatar     string
	Token      string
	PlatformID pb.PlatformID

	wsAddr       string
	wsClient     *websocket.Conn
	callback     func(count int, msg *pb.MsgData)
	shutdownChan chan interface{}
	lock         sync.RWMutex
}

func NewClient(
	userID string,
	nickname string,
	avatar string,
	token string,
	platformID pb.PlatformID,
	wsAddr string,
	callback func(count int, msg *pb.MsgData),
) *Client {
	c := &Client{
		UserID:       userID,
		Nickname:     nickname,
		Avatar:       avatar,
		Token:        token,
		PlatformID:   platformID,
		wsAddr:       wsAddr,
		callback:     callback,
		shutdownChan: make(chan interface{}),
	}
	go c.heartbeat()
	go c.receiveHandler()
	return c
}

func (user *Client) GetSingleChatData(
	toUserId string,
	msg Msg,
) *pb.MsgData {
	return &pb.MsgData{
		SendID:           user.UserID,
		RecvID:           toUserId,
		ClientMsgID:      time.Now().Format("2006/01/02/15/04/05/.000"),
		SenderPlatformID: user.PlatformID.Int32(),
		SenderNickname:   user.Nickname,
		SenderFaceURL:    user.Avatar,
		SessionType:      pb.SingleChatType.Int32(),
		//MsgFrom:          0, 好像没啥用
		ContentType: msg.GetContentType().Int32(),
		Content:     msg.GetContent().Bytes(),
		CreateTime:  time.Now().UnixMilli(),
		Options: map[string]bool{
			"offlinePush": msg.IsOfflinePush(),
		},
		OfflinePushInfo: &pb.OfflinePushInfo{
			Title: msg.OfflinePushTitle(), // 离线标题
		},
	}
}

func (user *Client) GetSuperGroupChatData(
	groupId string,
	msg Msg,
	atUserIds ...string,
) *pb.MsgData {
	return &pb.MsgData{
		SendID:           user.UserID,
		GroupID:          groupId,
		ClientMsgID:      time.Now().Format("2006/01/02/15/04/05/.000"),
		SenderPlatformID: user.PlatformID.Int32(),
		SenderNickname:   user.Nickname,
		SenderFaceURL:    user.Avatar,
		SessionType:      pb.SuperGroupChatType.Int32(),
		//MsgFrom:          0, 好像没啥用
		ContentType: msg.GetContentType().Int32(),
		Content:     msg.GetContent().Bytes(),
		CreateTime:  time.Now().UnixMilli(),
		Options: map[string]bool{
			"offlinePush": msg.IsOfflinePush(),
		},
		OfflinePushInfo: &pb.OfflinePushInfo{
			Title: msg.OfflinePushTitle(), // 离线标题
		},
		AtUserIDList: atUserIds,
	}
}
