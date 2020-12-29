package shandle

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"isface"
	"snet"
	"time"
)

type FindUserName struct{
	snet.BaseRouter
}

//管理员查找用户
//页码，页面大小，i是查找方法，sx是关键词
//目前只有i=1的一种查找方法，后面可能要做热度查找
type FindUserNameJson struct{
	PageNo int `json:PageNo`
	PageSize int `json:"Pagesize"`
	Sx string `json:"SX"`
}

type JSONTime struct {
	time.Time
}
// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
type FindUserNameGorm struct{
	Name string `gorm:"column:name"`
	Stats int `gorm:"column:stats"`
	ModifiedTime JSONTime `gorm:"column:modified_time"`
}


func(T FindUserName)Handle(request isface.IRequest){
	conn := request.GetConnection()
	recvData := FindUserNameJson{}

	json.Unmarshal(request.GetData(),&recvData)

	fmt.Println("findusername Handle 从客户端接收到消息：",recvData)

	Data := make([]FindUserNameGorm,0)

	snet.SDB.Debug().Raw("call find_username(?,?,?)",recvData.PageNo,recvData.PageSize,recvData.Sx).Scan(&Data)

	SendData,_ := json.Marshal(Data)

	conn.SendMesg([]byte(""),SendData)
}