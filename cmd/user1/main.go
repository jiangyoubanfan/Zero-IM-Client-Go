package main

import (
	"fmt"
	"github.com/showurl/Zero-IM-Client-Go/pb"
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
	return &pb.TextContent{
		Text: m.Text,
	}
}

func (m *msg) IsOfflinePush() bool {
	return true
}

func (m *msg) OfflinePushTitle() string {
	return ""
}

var (
	Second = time.Duration(6)
)

func main() {
	runChan := make(chan interface{})
	client := recvmsg()
	j := 0
	if false {
		go func() {
			for {
				j++
				timer := time.NewTimer(time.Second * Second)
				select {
				case <-timer.C:
					for i := 1; i <= 10; i++ {
						client.SendMsg(client.GetSingleChatData(
							fmt.Sprintf("%d", i),
							NewText(fmt.Sprintf("每隔%d秒发一次单聊消息， 这是第%d次", Second, j)),
						))
						time.Sleep(time.Second * Second)

					}
				}
			}
		}()
	}
	gj := 0
	if true {
		go func() {
			for {
				gj++
				timer := time.NewTimer(time.Second * Second)
				select {
				case <-timer.C:
					text := fmt.Sprintf("每隔%d秒发一次大群消息， 这是第%d次", Second, gj)
					data := client.GetSuperGroupChatData(
						"supergroup_0",
						NewText(text),
					)
					fmt.Printf(`
                                                                                        [%s]请求发送消息:
                                                                                %s`,
						data.ClientMsgID,
						text,
					)
					client.SendMsg(data)
				}
			}
		}()
	}
	<-runChan
}
