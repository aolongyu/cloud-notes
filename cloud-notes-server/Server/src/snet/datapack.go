package snet

import (
	"Settings"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"isface"
	"bytes"
	"encoding/binary"
	"errors"
	_ "golang.org/x/net/websocket"
	"strings"
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

func (this *DataPack)Webconn(c *Connection) {
	strbuf := make([]byte,2048)
	//读取Conn里面的值，最大包长限定2048，出错关闭这个连接
	_,err := c.Conn.Read(strbuf)
	if err != nil{
		fmt.Println("Webfirstconn读取数据时候失败", err)
		//Logs.Error("读取数据时候失败", err)
		if c.isClosed == true {
			return
		}
		c.isClosed = true
		c.ExitBuffChan <- true
		return
	}
	isHttp := false
	if string(strbuf[0:3]) == "GET" {
		isHttp = true
	}
	fmt.Println("isHttp: ",isHttp)

	if isHttp{
		//表示头部信息的Map[string]string
		headers := parseHandshake(string(strbuf))
		fmt.Println("headers: ",headers)
		//获取secwebsocket的值
		secWebsocketKey := headers["Sec-WebSocket-Key"]

		guid := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

		// 计算Sec-WebSocket-Accept
		h := sha1.New()

		fmt.Println("accept 的值为：",secWebsocketKey + guid)

		io.WriteString(h,secWebsocketKey+guid)

		accept := make([]byte,28)

		base64.StdEncoding.Encode(accept,h.Sum(nil))

		fmt.Println(string(accept))

		response := "HTTP/1.1 101 Switching Protocols\r\n"
		response = response + "Sec-WebSocket-Accept: " + string(accept) + "\r\n"
		response = response + "Connection: Upgrade\r\n"
		response = response + "Upgrade: websocket\r\n\r\n"

		fmt.Println("response:",response)

		if length,err := c.Conn.Write([]byte(response));err != nil{
			fmt.Println(err)
		}else{
			fmt.Println("Websocket Send Len:",length)
		}

	}
}
func parseHandshake(content string)map[string]string{
	headers := make(map[string]string,10)
	lines := strings.Split(content,"\r\n")
	for _,line := range lines {
		if len(line) >= 0{
			words := strings.Split(line,":")
			if len(words) == 2{
				headers[strings.Trim(words[0],"")] = strings.Trim(words[1], " ")
			}
		}
	}
	return headers
}

