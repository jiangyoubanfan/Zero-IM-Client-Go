package main

import (
	"fmt"
	"github.com/showurl/Zero-IM-Client-Go/pb"
	"strconv"
	"time"
)

type msg struct {
	Text string
}

func NewText(text string) *msg {
	return &msg{text}
}
func (m *msg) GetContentType() pb.ContentType {
	return pb.Text
}

func (m *msg) GetContent() pb.Content {
	return &pb.TextContent{Text: m.Text}
}

func (m *msg) IsOfflinePush() bool {
	return true
}

func (m *msg) OfflinePushTitle() string {
	return ""
}

var (
	Second = time.Duration(20)
)

func main() {
	runChan := make(chan interface{})
	client := recvmsg()
	j := 0
	f := func() {
		for i := 1; i <= 10; i++ {
			toUserId := strconv.Itoa(i)
			msgData := client.GetSingleChatData(
				toUserId,
				NewText(fmt.Sprintf("给 user"+toUserId+" 发一条消息")),
			)
			err := client.SendMsg(msgData)
			if err != nil {
				fmt.Printf(`
                                                                                        [%s]发送失败
`, msgData.ClientMsgID)
			}
			time.Sleep(time.Second * Second / 10)
		}
	}
	if false {
		go func() {
			f()
			for {
				j++
				timer := time.NewTimer(time.Second * Second)
				select {
				case <-timer.C:
					f()
				}
			}
		}()
	}

	<-runChan
}
