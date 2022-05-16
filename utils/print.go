package utils

import (
	"fmt"
	"github.com/showurl/Zero-IM-Client-Go/pb"
)

func PrintMsg(msg *pb.MsgData, selfID string) {
	switch msg.ContentType {
	case pb.Text.Int32():
		printText(msg, selfID)
	}
}

func printText(msg *pb.MsgData, selfID string) {
	if selfID != msg.SendID {
		fmt.Printf(
			`
[%d]%s: %s
`,
			msg.Seq,
			msg.SenderNickname,
			pb.NewTextContent(msg.Content).String(),
		)
	} else {
		fmt.Printf(
			`
											[%s]请求发送消息: 
										%s
`,
			msg.ClientMsgID,
			pb.NewTextContent(msg.Content).String(),
		)
	}
}
