package snet

import (
	"Settings"
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	_ "golang.org/x/net/websocket"
	"io"
	"isface"
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
func (Msg *DataPack) Unpack(data []byte)(isface.IMessage ,error) {
	en_bytes := []byte("")
	cn_bytes := make([]int,0)

	v := data[1]&0x7f
	p:=0
	res:=""
	switch v {
	case 0x7e:
		p =4
	case 0x7f:
		p = 10
	default:
		p = 2
	}
	mask := data[p:p+4]
	data_tmp := data[p+4:]
	nv :=""
	nv_bytes :=[]byte("")
	nv_len :=0

	for k,v := range(data_tmp){

		nv = string(int(v ^ mask[k%4]))
		nv_bytes = []byte(nv)
		nv_len = len(nv_bytes)
		if nv_len == 1 {
			en_bytes=BytesCombine(en_bytes,nv_bytes)
		}else{
			en_bytes=BytesCombine(en_bytes,[]byte("%s"))
			cn_bytes = append(cn_bytes,int(v ^ mask[k%4]))
		}
	}

	//处理中文
	cn_str := make([]interface{},0)
	if len(cn_bytes) >2 {

		clen := len(cn_bytes)
		count := int(clen/3)

		for i:=0;i<count;i++ {
			mm := i*3
			hh := make([]byte,3)
			h1  := IntToBytes(cn_bytes[mm])
			h2  := IntToBytes(cn_bytes[mm+1])
			h3  := IntToBytes(cn_bytes[mm+2])
			hh[0]=h1[0]
			hh[1]=h2[0]
			hh[2]=h3[0]
			cn_str = append(cn_str, string(hh))
		}
		new := string(bytes.Replace(en_bytes,[]byte("%s%s%s"),[]byte("%s"),-1))
		res = fmt.Sprintf(new,cn_str...)
	}else{
		res = string(en_bytes)
	}
	fmt.Println("拆包结果得到：",res)
	return NewMsgPackage(404,data),nil
}
var Tempbuf = make([]byte,Settings.GlobalObject.MaxPacketSize)

func (this *DataPack)Webconn(c *Connection) {

	//读取Conn里面的值，最大包长限定2048，出错关闭这个连接
	nums,err := c.Conn.Read(Tempbuf)
	strbuf := make([]byte,nums)
	for i:=0;i<nums;i++{
		strbuf[i] = Tempbuf[i]
	}
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

		fmt.Println("回应信息: ",response)

		if length,err := c.Conn.Write([]byte(response));err != nil{
			fmt.Println(err)
		}else{
			fmt.Println("Websocket 首次发送长度:",length)
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

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func IntToBytes(n int) ([]byte) {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}