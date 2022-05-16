package pb

type (
	ContentType int32
	Content     interface {
		Bytes() []byte
		String() string
	}
)

func (c ContentType) Int32() int32 {
	return int32(c)
}

const ( ///消息类型
	Text           ContentType = 101
	Picture        ContentType = 102
	Voice          ContentType = 103
	Video          ContentType = 104
	File           ContentType = 105
	AtText         ContentType = 106
	Merger         ContentType = 107
	Card           ContentType = 108
	Location       ContentType = 109
	Custom         ContentType = 110
	Revoke         ContentType = 111
	HasReadReceipt ContentType = 112
	Typing         ContentType = 113
	Quote          ContentType = 114
	Common         ContentType = 200
	GroupMsg       ContentType = 201
)
