package isface

/*
	封包拆包

 */
type IDataPack interface {
	//获取包头长度
	GetHendLen() uint32

	//进行封包
	Pack(msg IMessage)([]byte , error)

	//进行拆包,拆出数据段
	Unpack([]byte)(IMessage ,error)

}
