package pb

type (
	SessionType int32
)

const (
	SingleChatType       SessionType = 1
	GroupChatType        SessionType = 2
	SuperGroupChatType   SessionType = 3
	NotificationChatType SessionType = 4
)

func (s SessionType) Int32() int32 {
	return int32(s)
}
