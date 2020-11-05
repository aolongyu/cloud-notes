package snet

import (
	"Settings"
	"isface"
	"bytes"
	"encoding/binary"
	"errors"
)

//拆包的实例对象，拥有拆包封包功能
type DataPack struct{

}

//返回拆包封包实例初始化
func NewDataPack() *DataPack  {
	return &DataPack{

	}
}

//获取包头长度
// Len  消息ID  {消息Json格式}
//4字节  4字节   Len长度字节
func (Msg *DataPack) GetHendLen() uint32{
	return 8
}



//进行封包 ,以大端序封装
func (Msg *DataPack) Pack(msg isface.IMessage)([]byte , error){
	//创建一个存放bytes的缓冲
	dataBuf := bytes.NewBuffer([]byte{})

	//以大端序存储信息，将ID封入网络协议中
	if err := binary.Write(dataBuf,binary.LittleEndian,msg.GetMsgId()); err != nil{
		return nil,err
	}

	//写dataLen
	if err := binary.Write(dataBuf,binary.LittleEndian,msg.GetDataLen());err != nil {
		return nil,err
	}

	//写data数据
	if err := binary.Write(dataBuf,binary.LittleEndian,msg.GetData()); err != nil{
		return nil,err
	}
	return dataBuf.Bytes(),nil
}

//进行拆包,拆出数据段
func (Msg *DataPack) Unpack(Data []byte)(isface.IMessage ,error) {
	//创建一个ioReader
	dataBuf := bytes.NewReader(Data)

	//创建解压信息的实例
	msg := &Message{}

	//解压出ID号长度4字节
	if err:= binary.Read(dataBuf,binary.LittleEndian,&msg.Id); err != nil{
		return nil,err
	}
	//解压DataLen，长度4字节
	if err:= binary.Read(dataBuf,binary.LittleEndian,&msg.DataLen); err != nil{
		return nil,err
	}

	if msg.DataLen > Settings.GlobalObject.MaxPacketSize {
		return nil,errors.New("数据包太大，无法接受")
	}

	//可以直接用链接ID，读取conn 通过包头长度获得该包的数据。
	return msg,nil
}
