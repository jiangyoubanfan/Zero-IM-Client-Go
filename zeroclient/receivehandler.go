package zeroclient

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/showurl/Zero-IM-Client-Go/pb"
	"time"
)

func (user *Client) receiveHandler() {
	count := 0
	for {
		if user.wsClient == nil {
			time.Sleep(time.Second)
			continue
		}
		_, msg, err := user.wsClient.ReadMessage()
		if err != nil {
			time.Sleep(time.Second * 5)
			fmt.Println("Error in receive:", err)
			continue
		}
		count++
		resp := &Resp{}
		err = gob.NewDecoder(bytes.NewReader(msg)).Decode(resp)
		if err != nil {
			fmt.Println("Error in Decode gob:", err)
			continue
		}
		if resp.ReqIdentifier == pb.WSGetNewestSeq.Int32() {
			res := &pb.GetMaxAndMinSeqResp{}
			e := proto.Unmarshal(resp.Data, res)
			if e == nil {
				if _, max := user.getSeq(); max < res.MaxSeq {
					user.pullMsg(0, res.MaxSeq)
				}
			}
		} else if resp.ReqIdentifier == pb.WSGetNewestSuperGroupSeq.Int32() {
			res := &pb.GetMaxAndMinSuperGroupSeqResp{}
			e := proto.Unmarshal(resp.Data, res)
			if e == nil {
				for _, item := range res.SuperGroupSeqList {
					if _, max := user.getSuperGroupSeq(item.SuperGroupID); max < item.MaxSeq {
						user.pullSuperGroupMsg(0, item.MaxSeq, item.SuperGroupID)
					}
				}
			}
		} else if resp.ReqIdentifier == pb.WSSendMsg.Int32() {
			data := &pb.UserSendMsgResp{}
			err = proto.Unmarshal(resp.Data, data)
			if err != nil {
				continue
			}
			fmt.Printf(`
											[%s]发送成功
`, data.ClientMsgID)
		} else if resp.ReqIdentifier == pb.WSPushMsg.Int32() {
			msgData := &pb.MsgData{}
			err = proto.Unmarshal(resp.Data, msgData)
			if err != nil {
				continue
			}
			if msgData.Seq == 0 {
				continue
			}
			if _, max := user.getSeq(); max < msgData.Seq {
				user.saveSeq(0, msgData.Seq)
			}
			user.callback(count, msgData)
		} else if resp.ReqIdentifier == pb.WSSuperGroupPushMsg.Int32() {
			msgData := &pb.MsgData{}
			err = proto.Unmarshal(resp.Data, msgData)
			if err != nil {
				continue
			}
			if msgData.Seq == 0 {
				continue
			}
			if _, max := user.getSeq(); max < msgData.Seq {
				user.saveSuperGroupSeq(0, msgData.Seq, msgData.GroupID)
			}
			user.callback(count, msgData)
		} else if resp.ReqIdentifier == pb.WSPullMsgBySuperGroupSeqList.Int32() {
			seqListResp := &pb.PullMessageBySeqListResp{}
			err = proto.Unmarshal(resp.Data, seqListResp)
			if err != nil {
				continue
			}
			for _, msgData := range seqListResp.List {
				if msgData.Seq == 0 {
					continue
				}
				if _, max := user.getSuperGroupSeq(msgData.GroupID); max < msgData.Seq {
					user.saveSuperGroupSeq(0, msgData.Seq, msgData.GroupID)
				}
				user.callback(count, msgData)
			}
		} else if resp.ReqIdentifier == pb.WSPullMsgBySeqList.Int32() {
			seqListResp := &pb.PullMessageBySeqListResp{}
			err = proto.Unmarshal(resp.Data, seqListResp)
			if err != nil {
				continue
			}
			for _, msgData := range seqListResp.List {
				if msgData.Seq == 0 {
					continue
				}
				if _, max := user.getSuperGroupSeq(msgData.GroupID); max < msgData.Seq {
					user.saveSeq(0, msgData.Seq)
				}
				user.callback(count, msgData)
			}
		}
	}
}
