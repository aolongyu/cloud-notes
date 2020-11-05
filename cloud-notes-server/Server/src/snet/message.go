package snet

type Message struct{
	//消息ID
	Id uint32

	//消息长度
	DataLen uint32

	//消息内容
	Data []byte
}

//创建一个Message实例
func NewMsgPackage(id uint32,data []byte) *Message{
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

//获取消息数据段长度
func (ms *Message) GetDataLen() uint32{
	return ms.DataLen
}

//获取消息ID
func (ms *Message) GetMsgId() uint32{
	return ms.Id
}

//获取消息内容
func (ms *Message) GetData() []byte{
	return ms.Data
}

//设置消息ID
func (ms *Message) SetMsgId(id uint32){
	ms.Id = id
}

//设置消息内容
func (ms *Message) SetData(data []byte){
	ms.Data = data
}

//设置消息数据段长度
func (ms *Message) SetDataLen(Len uint32){
	ms.DataLen = Len
}
