package pb

type (
	Identifier int32
)

const (
	WSGetNewestSeq               Identifier = 1001
	WSPullMsgBySeqList           Identifier = 1002
	WSSendMsg                    Identifier = 1003
	WSSendSignalMsg              Identifier = 1004
	WSGetNewestSuperGroupSeq     Identifier = 1005
	WSPullMsgBySuperGroupSeqList Identifier = 1006

	WSPushMsg           Identifier = 2001
	WSKickOnlineMsg     Identifier = 2002
	WsLogoutMsg         Identifier = 2003
	WSSuperGroupPushMsg Identifier = 2004
)

func (i Identifier) Int32() int32 {
	return int32(i)
}
