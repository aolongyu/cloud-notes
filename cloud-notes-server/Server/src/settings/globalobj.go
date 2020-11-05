package Settings

import (
	"isface"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//全局变量
type GlobalObj struct{

	/*
	Server list
	 */
	//当前全局参数
	TCPServer isface.IServer

	//主机IP
	Host string

	//服务器主机监听端口号
	TcpPort int

	//服务器版本号
	Version string

	//最大包长
	MaxPacketSize uint32

	//配置文件路径
	ConFilePath string

	//当前服务器主机允许的最大链接个数
	MaxConn       int

	MaxHeath 	  int32

	//当前业务工作池的数量
	WorkerPoolSize int32

	//最多一个work能多少个消息
	MaxWorkerSize int32

	//游戏玩家数量
	PlayerNum int32

	//定时器时间（ms级别）
	Time int32
}

//全局变量实例
var GlobalObject *GlobalObj

//判断一个文件是否存在
func PathExists(path string) (bool,error) {

	_,err := os.Stat(path)

	if err == nil{

		return true,nil
	}
	if os.IsNotExist(err){
		fmt.Println("服务器打开路径失败")
		return false,nil
	}
	return false ,err
}

//读取用户配置文件

func (g *GlobalObj) Reload(){

		if confFileExists,_ := PathExists(g.ConFilePath) ; confFileExists != true{

			return
		}
		//从路径中读取文件
		data ,err := ioutil.ReadFile(GlobalObject.ConFilePath)
		if err!= nil {
			panic(err)
		}

		//把json数据解析到struct中
		err = json.Unmarshal(data,&GlobalObject)
		if err != nil{
			panic(err)
		}

}

func init(){

	//初始化GlobalObject变量
	GlobalObject = &GlobalObj{
		ConFilePath:   "../settings/Server.json",
	}

	//从文件中读取路径并且把值读入
	GlobalObject.Reload()
}
