package main

import (
	"github.com/showurl/Zero-IM-Client-Go/pb"
	"github.com/showurl/Zero-IM-Client-Go/utils"
	"github.com/showurl/Zero-IM-Client-Go/zeroclient"
)

func recvmsg() *zeroclient.Client {
	return zeroclient.NewClient(
		"3",
		"user3",
		"avatar3",
		"token3",
		pb.OSXPlatformID,
		"ws://localhost:17778",
		func(count int, msg *pb.MsgData) {
			utils.PrintMsg(msg, "3")
		},
	)
}
