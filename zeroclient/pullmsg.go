package zeroclient

func (user *Client) pullMsg(min uint32, max uint32) {
	_, originMax := user.getSeq()
	if originMax < max {
		// 拉取缺失的消息
		var needSeqs []uint32
		for i := originMax + 1; i <= max; i++ {
			needSeqs = append(needSeqs, i)
		}
		user.getSeqMsgs(needSeqs)
	}
}

func (user *Client) pullSuperGroupMsg(min uint32, max uint32, groupId string) {
	_, originMax := user.getSuperGroupSeq(groupId)
	if originMax < max {
		// 拉取缺失的消息
		var needSeqs []uint32
		for i := originMax + 1; i <= max; i++ {
			needSeqs = append(needSeqs, i)
		}
		user.getSuperGroupSeqMsgs(needSeqs, groupId)
	}
}
