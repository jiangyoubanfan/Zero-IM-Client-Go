package main

import (
	"github.com/showurl/Zero-IM-Client-Go/pb"
	"github.com/showurl/Zero-IM-Client-Go/utils"
	"github.com/showurl/Zero-IM-Client-Go/zeroclient"
)

func recvmsg() *zeroclient.Client {
	return zeroclient.NewClient(
		"1",
		"user1",
		"avatar1",
		"token1",
		pb.IOSPlatformID,
		"ws://localhost:17778",
		func(count int, msg *pb.MsgData) {
			utils.PrintMsg(msg, "1")
		},
	)
}
