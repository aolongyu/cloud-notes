package isface

//一个请求的消息封装到Message中
//格式
// Len  消息ID  {消息Json格式}
//4字节  4字节   Len长度字节
type IMessage interface {
	//获取消息数据段长度
	GetDataLen() uint32

	//获取消息ID
	GetMsgId() string

	//获取消息内容
	GetData() []byte

	//设置消息ID
	SetMsgId(string)

	//设置消息内容
	SetData([]byte)

	//设置消息数据段长度
	SetDataLen(uint32)
}
