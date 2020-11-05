package myhandler

import (
	"fmt"
	"time"

	"../myconcrete"
	"../myinterface"
	"google.golang.org/protobuf/proto"
)

//var Uid int = 1

const (
	LAND                   = 1   //登录处理
	LAND_ACK_SUC           = 10  //返回登录成功
	LAND_ACK_FAIL          = 11  //返回登录失败
	LAND_REGISTER          = 100 //注册处理
	LAND_REGISTER_ACK_SUC  = 101 //注册成功
	LAND_REGISTER_ACK_FAIL = 102 //注册失败
)

//LandingRegister 注册业务
type LandingRegister struct {
	myconcrete.BaseRouter
}

//Handle 注册业务处理方法
func (lr *LandingRegister) Handle(request myinterface.IRequest) {
	fmt.Println("[langding_refister Test]msgID = ,", request.GetMsgID(), "序列化data = ", request.GetMsgData())
	cus := &Customer{}
	proto.Unmarshal(request.GetMsgData(), cus)
	//LastTmp := &User_info{}
	//myconcrete.SDB.AutoMigrate(&User_info{})
	//myconcrete.SDB.LogMode(true).First(&User_info{ /*Account_name: "root", Account_password: "123", User_name: "root"*/ }).Scan(&LastTmp)

	//fmt.Println("----", LastTmp.UID, "----")
	timeUnix := int32(time.Now().Unix())
	u := &User_info{
		UID:              timeUnix,
		Account_name:     cus.User,
		Account_password: cus.Pass,
		User_name:        cus.Nickname,
	}
	//h := &Hero{
	/*
		type ans struct {
			UID int32
		}
		now := &ans{}
		u1 := User_info{}
		u2 := User_info{}
		fmt.Println("u.uid = ", u.UID)
		//myconcrete.SDB.Exec("INSERT into user_info(u,user_password) values(?,?)", Register.Username, Register.Password)
		cre := myconcrete.SDB.LogMode(true).Create(&u)
		//var id int32
		myconcrete.SDB.Raw("Select uid as UID from user_info where account_name = ? and account_password = ? ", cus.User, cus.Pass).Scan(&now)
		cre.Where(User_info{User_name: "test"}).First(&u2)
		fmt.Println("u2: ", u2)
	*/
	if err := myconcrete.SDB.Create(u).Error; err != nil {
		//myconcrete.SDB.Create()
		request.GetConnection().SendMsg(LAND_REGISTER_ACK_FAIL, []byte(""))
		fmt.Println("failed ", err)
	} else {
		//fmt.Println("-----=u1: ", u1)
		h1 := &Hero{
			UID:         timeUnix,
			Hero_id:     1,
			Hero_lv:     1,
			Hero_exp:    0,
			Hero_weapon: 1,
		}
		h2 := &Hero{
			UID:         timeUnix,
			Hero_id:     2,
			Hero_lv:     1,
			Hero_exp:    0,
			Hero_weapon: 1,
		}
		h3 := &Hero{
			UID:         timeUnix,
			Hero_id:     3,
			Hero_lv:     1,
			Hero_exp:    0,
			Hero_weapon: 1,
		}
		ac := &Account_info{
			UID:             timeUnix,
			Account_lv:      1,
			Account_exp:     0,
			Weapon_shard:    0,
			Star_stage:      0,
			Choosen_hero_id: 1,
			Choosen_skill:   1,
		}

		myconcrete.SDB.Create(h1)
		myconcrete.SDB.Create(h2)
		myconcrete.SDB.Create(h3)
		myconcrete.SDB.Create(ac)
		request.GetConnection().SendMsg(LAND_REGISTER_ACK_SUC, []byte(""))
		fmt.Println("success")
	}
}

//Landing 登录业务
type Landing struct {
	myconcrete.BaseRouter
}

//Handle 登录业务处理方法
func (l *Landing) Handle(request myinterface.IRequest) {
	fmt.Println("[langding Test]msgID = ,", request.GetMsgID(), "序列化data = ", request.GetMsgData())
	cus := &Customer{}
	//cus1 := &Customer{}
	proto.Unmarshal(request.GetMsgData(), cus)
	//fmt.Println("cus : ", cus)
	//use := &User_info{}
	//use.User_name = cus.User
	//use.User_password = cus.Pass
	//fmt.Println("use = ", use)
	var users []User_info
	var count int = 0
	myconcrete.SDB.Where("Account_name=? AND Account_password=?", cus.User, cus.Pass).Find(&users).Count(&count)
	//myconcrete.SDB.Where("User_name=? AND User_password=?", cus.User, cus.Pass).Find(&users).Count(&count)
	//fmt.Println("users : \n", users)
	//u := &User_info{User_name: cus.User, User_password: cus.Pass}
	//myconcrete.SDB.Model(u).Count(&count)
	//fmt.Println("count", count)
	if count == 1 {
		//fmt.Println("1111")
		request.GetConnection().SendMsg(LAND_ACK_SUC, []byte("success"))
	} else {
		//fmt.Println("2222")
		request.GetConnection().SendMsg(LAND_ACK_FAIL, []byte("failed"))
	}

}

//Pingtest 测试业务
type Pingtest struct {
	myconcrete.BaseRouter
}

//Handle 测试业务处理方法
func (p *Pingtest) Handle(request myinterface.IRequest) {

	//fmt.Println("this is a Ping Handle")

	fmt.Println("[Ping Test]msgID = ,", request.GetMsgID(), "序列化data = ", request.GetMsgData())
	cus := &Customer{}
	cus1 := &Customer{}
	proto.Unmarshal(request.GetMsgData(), cus)

	fmt.Println("反序列化data = ", cus)
	if cus.User == "abc" {
		cus1.User = "000"
		fmt.Println("111 = ", cus1.User)
		if cus.Pass == "123" {
			cus1.Pass = "000"
			fmt.Println("222 = ", cus1.Pass)
		}
	}

	data, _ := proto.Marshal(cus1)
	err := request.GetConnection().SendMsg(5, data) // []byte("This is"))

	if err != nil {
		fmt.Println("ping err : ", err)
	}
}
