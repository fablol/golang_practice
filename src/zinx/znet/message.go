package znet

type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		ID:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMsgID() uint32 {
	return m.ID
}
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}
func (m *Message) GetMsgData() []byte {
	return m.Data
}
func (m *Message) SetMsgID(id uint32) {
	m.ID = id
}
func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}
func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}
