package zeroclient

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func (user *Client) heartbeat() {
	socketUrl := fmt.Sprintf("%s?token=12345&sendID=%s&platformID=%s", user.wsAddr, user.UserID, user.PlatformID.String())
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Println("连接 server 失败:", err)
		time.Sleep(time.Second * 5)
		user.heartbeat()
		//log.Fatal("连接 server 失败:", err)
	}
	user.wsClient = conn
	defer conn.Close()
	go user.maxGroupSeq()
	user.maxSeq()
	for {
		timer := time.NewTimer(time.Second * 30)
		select {
		case <-user.shutdownChan:
			return
		case <-timer.C:
			go user.maxGroupSeq()
			user.maxSeq()
		}
	}

}
