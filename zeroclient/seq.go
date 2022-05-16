package zeroclient

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/showurl/Zero-IM-Client-Go/pb"
	"github.com/showurl/Zero-IM-Client-Go/utils"
	"log"
	"os"
	"strconv"
	"strings"
)

func (user *Client) maxSeq() {
	rpcReq := &pb.GetMaxAndMinSeqReq{
		UserID: user.UserID,
	}
	buf, _ := proto.Marshal(rpcReq)
	// 获取服务端最大seq
	req := &Req{
		ReqIdentifier: pb.WSGetNewestSeq.Int32(),
		Token:         user.Token,
		SendID:        user.UserID,
		MsgIncr:       "unknown",
		Data:          buf,
	}
	err := user.sendMsg(req)
	if err != nil {
		log.Println("获取服务端最大seq:", err.Error())
	}
}

func (user *Client) maxGroupSeq() {
	// TODO 获取当前用户所有的群id
	var groupIds = []string{
		"supergroup_0",
	}
	rpcReq := &pb.GetMaxAndMinSuperGroupSeqReq{
		SuperGroupIDList: groupIds,
	}
	buf, _ := proto.Marshal(rpcReq)
	// 获取服务端最大seq
	req := &Req{
		ReqIdentifier: pb.WSGetNewestSuperGroupSeq.Int32(),
		Token:         user.Token,
		SendID:        user.UserID,
		MsgIncr:       "unknown",
		Data:          buf,
	}
	err := user.sendMsg(req)
	if err != nil {
		log.Println("获取服务端最大seq:", err.Error())
	}
}

func (user *Client) saveSeq(min uint32, max uint32) {
	if min > 0 {
		minFile := fmt.Sprintf("%s.%s.min.seq", user.UserID, user.PlatformID.String())
		utils.ExecCommand("bash", "-c", fmt.Sprintf(
			`echo "%d" > %s`,
			min, minFile,
		))
	}
	if max > 0 {
		maxFile := fmt.Sprintf("%s.%s.max.seq", user.UserID, user.PlatformID.String())
		utils.ExecCommand("bash", "-c", fmt.Sprintf(
			`echo "%d" > %s`,
			max, maxFile,
		))
	}
}

func (user *Client) saveSuperGroupSeq(min uint32, max uint32, groupId string) {
	if min > 0 {
		minFile := fmt.Sprintf("%s.%s.%s.min.seq", user.UserID, groupId, user.PlatformID.String())
		utils.ExecCommand("bash", "-c", fmt.Sprintf(
			`echo "%d" > %s`,
			min, minFile,
		))
	}
	if max > 0 {
		maxFile := fmt.Sprintf("%s.%s.%s.max.seq", user.UserID, groupId, user.PlatformID.String())
		utils.ExecCommand("bash", "-c", fmt.Sprintf(
			`echo "%d" > %s`,
			max, maxFile,
		))
	}
}

func (user *Client) getSeq() (min uint32, max uint32) {
	minFile := fmt.Sprintf("%s.%s.min.seq", user.UserID, user.PlatformID.String())
	maxFile := fmt.Sprintf("%s.%s.max.seq", user.UserID, user.PlatformID.String())
	minf, _ := os.ReadFile(minFile)
	maxf, err := os.ReadFile(maxFile)
	if err != nil {
		return 0, 0
	}
	minI, _ := strconv.Atoi(strings.TrimSpace(string(minf)))
	maxI, _ := strconv.Atoi(strings.TrimSpace(string(maxf)))
	return uint32(minI), uint32(maxI)
}

func (user *Client) getSuperGroupSeq(groupId string) (min uint32, max uint32) {
	minFile := fmt.Sprintf("%s.%s.%s.min.seq", user.UserID, groupId, user.PlatformID.String())
	maxFile := fmt.Sprintf("%s.%s.%s.max.seq", user.UserID, groupId, user.PlatformID.String())
	minf, _ := os.ReadFile(minFile)
	maxf, err := os.ReadFile(maxFile)
	if err != nil {
		return 0, 0
	}
	minI, _ := strconv.Atoi(strings.TrimSpace(string(minf)))
	maxI, _ := strconv.Atoi(strings.TrimSpace(string(maxf)))
	return uint32(minI), uint32(maxI)
}

func (user *Client) getSeqMsgs(seqs []uint32) {
	rpcReq := &pb.PullMessageBySeqListReq{
		UserID:  user.UserID,
		SeqList: seqs,
	}
	buf, _ := proto.Marshal(rpcReq)
	// 获取服务端最大seq
	req := &Req{
		ReqIdentifier: pb.WSPullMsgBySeqList.Int32(),
		Token:         user.Token,
		SendID:        user.UserID,
		MsgIncr:       "unknown",
		Data:          buf,
	}
	err := user.sendMsg(req)
	if err != nil {
		log.Println("获取服务端最大seq:", err.Error())
	}
}

func (user *Client) getSuperGroupSeqMsgs(seqs []uint32, groupId string) {
	rpcReq := &pb.PullMessageBySuperGroupSeqListReq{
		GroupID: groupId,
		SeqList: seqs,
	}
	buf, _ := proto.Marshal(rpcReq)
	// 获取服务端最大seq
	req := &Req{
		ReqIdentifier: pb.WSPullMsgBySuperGroupSeqList.Int32(),
		Token:         user.Token,
		SendID:        user.UserID,
		MsgIncr:       "unknown",
		Data:          buf,
	}
	err := user.sendMsg(req)
	if err != nil {
		log.Println("获取服务端最大seq:", err.Error())
	}
}
