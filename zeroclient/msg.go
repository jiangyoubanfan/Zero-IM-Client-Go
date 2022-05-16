package zeroclient

import "github.com/showurl/Zero-IM-Client-Go/pb"

type Msg interface {
	GetContentType() pb.ContentType
	GetContent() pb.Content
	IsOfflinePush() bool
	OfflinePushTitle() string
}
