package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang_practice/zinx/utils"
	"golang_practice/zinx/ziface"
)

// unpack
type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (dp *DataPack) GetHeadLen() uint32 {
	// len + msgid
	return 8
}

func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	data_buff := bytes.NewBuffer([]byte{})
	// write len
	if err := binary.Write(data_buff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	// write id
	if err := binary.Write(data_buff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	// write data
	if err := binary.Write(data_buff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}
	return data_buff.Bytes(), nil
}

func (dp *DataPack) Unpack(binary_data []byte) (ziface.IMessage, error) {
	// io reader
	data_buff := bytes.NewReader(binary_data)
	// read head
	msg := &Message{}
	if err := binary.Read(data_buff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(data_buff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}

	// edge
	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("msg data too large!")
	}

	return msg, nil
}
