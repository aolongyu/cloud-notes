package snet

import (
	"errors"
	"fmt"
	"io"
	"isface"
	"net"
)

//icconect的实现层

type Connection struct {

	//当前连接的socket
	Conn *net.TCPConn

	//当前连接的ID，之后可以当成客户端的ID
	ConnId uint32

	//当前连接的状态是否为关闭
	isClosed bool

	//当前连接管理消息如何处理的模块
	MsgHandle isface.IMsgHandle

	//告知其他函数该链接已经断开的管道
	ExitBuffChan chan bool

	//无缓冲管道，用于读写分离
	msgChan chan []byte

	//RoomManger的信息
	rom RoomMannger

	//TODO 日后扩展缓冲队列把数据发送给channel下，交给工作池工作
	//SendBufchan chan []byte

}

//创建一个链接
func NewConntion(conn *net.TCPConn, connId uint32, msgHandler isface.IMsgHandle) *Connection {

	c := &Connection{
		Conn:         conn,
		ConnId:       connId,
		isClosed:     false,
		MsgHandle:    msgHandler,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
	}
	return c
}

//写端，用户可以将消息发送给该链接的客户端
func (c *Connection) StartWriter() {
	defer fmt.Println(c.GetAddr().String(), "写端链接断开")
	//defer Logs.Warn("写端链接断开")
	defer c.Stop() //释放该链接的资源
	//阻塞该链接的写端
	fmt.Println("Id为：", c.ConnId, "IP地址为", c.GetAddr(), "的地址写端开启")
	//Logs.Debug("Id为：", c.ConnId, "IP地址为", c.GetAddr(), "的地址写端开启")

	for {
		if c.isClosed == true {
			return
		}
		select {
		case data := <-c.msgChan:
			//当有数据写入的时候select 就执行
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("发送数据失败 错误为:", err)
				//Logs.Error("发送数据失败 错误为:", err)
				return
			}
		case <-c.ExitBuffChan:
			//coon关闭了，由一个bool管道通知所有阻塞协程关闭链接
			return
		}
	}
}

func IsClosedch(ch <-chan interface{}) bool {
	select {
	case <-ch:
		return true
	default:

	}
	return false
}

//读端，从客户端中读取消息
func (c *Connection) StartReader() {
	fmt.Println("id为", c.ConnId, "Ip：", c.GetAddr(), "的地址读端开启")
	//Logs.Debug("id为", c.ConnId, "Ip：", c.GetAddr(), "的地址读端开启")
	//defer Logs.Warn(c.GetAddr().String(), "的地址读端退出")
	defer fmt.Println(c.GetAddr().String(), "的地址读端退出")
	defer c.Stop()

	for {
		//读对象阻塞函数，创建一个拆包实例
		pack := NewDataPack()

		headData := make([]byte, pack.GetHendLen())

		if _, err := io.ReadFull(c.GetTcpConnection(), headData); err != nil {
			fmt.Println("读取数据时候失败", err)
			//Logs.Error("读取数据时候失败", err)
			if c.isClosed == true {
				break
			}
			c.isClosed = true
			c.ExitBuffChan <- true
			break
		}
		msg, err := pack.Unpack(headData)
		if err != nil {
			fmt.Println("解包失败 error", err)
			//Logs.Error("解包失败 error", err)
			c.isClosed = true
			c.ExitBuffChan <- true
			break
		}

		//根据Len读取data，放在msg.Data中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTcpConnection(), data); err != nil {
				fmt.Println("读取数据错误 :", err)
				//Logs.Error("读取数据错误 :", err)
				c.isClosed = true
				c.ExitBuffChan <- true
				break
			}
		}

		msg.SetData(data)
		//fmt.Println("获得消息",msg)
		//得到当前客户端请求的Request数据
		req := Request{
			conn: c,
			msg:  msg,
		}

		c.MsgHandle.SendMsgToMesQueue(&req)
		//这个链接来一个消息，单位时间消息数+1

		//以下是单纯开辟的时候用的方法
		//go c.MsgHandle.DoMsgHandler(&req)

	}
}

//启动连接，开始工作。
func (c *Connection) Start() {
	//开启读端写端
	go c.StartReader()
	go c.StartWriter()

	//得到退出消息就不再阻塞
	for {
		select {
		case <-c.ExitBuffChan:
			return
		}
	}
}

//停止连接，结束当前连接状态
func (c *Connection) Stop() {
	//如果连接已经关闭，那么就退出了
	//回收资源

	//如果有房间的话则执行该函数，对房间进行相应的操作
	connOne := ConnMap[c.GetConnID()]
	if connOne != nil {
		connOne.DisConn()
	}

	//c.loginFromDisconn()

	if c.isClosed == true {
		return
	}

	//TODO 如果用户关闭链接时候有什么回调函数，则在此处进行调用

	//关闭socket
	c.Conn.Close()

	//通知其他阻塞函数关闭业务
	c.ExitBuffChan <- true

	c.isClosed = true

	close(c.ExitBuffChan)
	//delete(ConnMap, c.GetConnID())

}

//获取当前链接的conn
func (c *Connection) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

//获取当前连接的ID,也算是后来的游戏玩家ID（唯一)
func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

//设置当前连接的ID
func (c *Connection) SetConnID(id uint32) {
	c.ConnId = id
}

//当前连接的的IP地址
func (c *Connection) GetAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//服务器发送数据给客户端
func (c *Connection) SendMesg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		//Logs.Error("链接关闭，无法发送信息")

		return errors.New("链接关闭，无法发送信息")
	}

	//进行封包发送
	pack := NewDataPack()
	msg, err := pack.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("封包失败 ，消息ID为", msgId)
		//Logs.Error("封包失败，消息id为 "msgId)
		return errors.New("[error]封包失败 ")
	}
	c.msgChan <- msg
	return nil
}

//
//func (c *Connection) GetRoomManager() RoomMannger {
//	return c.rom
//}
