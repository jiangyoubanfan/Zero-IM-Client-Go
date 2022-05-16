package main

import (
	"github.com/showurl/Zero-IM-Client-Go/pb"
	"github.com/showurl/Zero-IM-Client-Go/utils"
	"github.com/showurl/Zero-IM-Client-Go/zeroclient"
)

func recvmsg() *zeroclient.Client {
	return zeroclient.NewClient(
		"2",
		"user2",
		"avatar2",
		"token2",
		pb.AndroidPlatformID,
		"ws://localhost:17778",
		func(count int, msg *pb.MsgData) {
			utils.PrintMsg(msg, "2")
		},
	)
}
