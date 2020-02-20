package ziface

type IMessage interface {
	GetMsgID() uint32
	GetMsgLen() uint32
	GetMsgData() []byte

	SetMsgID(uint32)
	SetMsgLen(uint32)
	SetMsgData([]byte)
}
