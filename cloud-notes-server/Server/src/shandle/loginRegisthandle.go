package shandle

import (
	Data "data"
	"fmt"
	"time"

	"isface"
	"snet"

	//"../isface"
	//"../snet"

	//"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/proto"
)

//"fmt"

//"github.com/jinzhu/gorm"

//var Uid int = 1

//LandingRegister 注册业务
type LandingRegister struct {
	snet.BaseRouter
}

//Handle 注册业务处理方法
func (lr *LandingRegister) Handle(request isface.IRequest) {
	fmt.Println("[langding_refister Test]msgID = ,", request.GetMsgId(), "序列化data = ", request.GetData())
	cus := &Data.Customer{}
	proto.Unmarshal(request.GetData(), cus)
	//LastTmp := &User_info{}
	//snet.SDB.AutoMigrate(&User_info{})
	//snet.SDB.LogMode(true).First(&User_info{ /*Account_name: "root", Account_password: "123", User_name: "root"*/ }).Scan(&LastTmp)

	//fmt.Println("----", LastTmp.UID, "----")
	timeUnix := int32(time.Now().Unix())
	u := &Data.User_info{
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
		//snet.SDB.Exec("INSERT into user_info(u,user_password) values(?,?)", Register.Username, Register.Password)
		cre := snet.SDB.LogMode(true).Create(&u)
		//var id int32
		snet.SDB.Raw("Select uid as UID from user_info where account_name = ? and account_password = ? ", cus.User, cus.Pass).Scan(&now)
		cre.Where(User_info{User_name: "test"}).First(&u2)
		fmt.Println("u2: ", u2)
	*/
	fmt.Println("用户", u.UID, "注册成功 信息为:", u)
	if err := snet.SDB.Create(u).Error; err != nil {
		//snet.SDB.Create()
		request.GetConnection().SendMesg(LAND_REGISTER_ACK, []byte("0"))
		fmt.Println("failed ", err)
	} else {
		fmt.Println("---------------插入成功: ")
		h1 := &Data.Hero{
			UID:        timeUnix,
			Hero_id:    1,
			Hero_lv:    1,
			Hero_exp:   0,
			Hero_skill: 1,
		}
		h2 := &Data.Hero{
			UID:        timeUnix,
			Hero_id:    2,
			Hero_lv:    1,
			Hero_exp:   0,
			Hero_skill: 1,
		}
		h3 := &Data.Hero{
			UID:        timeUnix,
			Hero_id:    3,
			Hero_lv:    1,
			Hero_exp:   0,
			Hero_skill: 1,
		}
		ac := &Data.Account_info{
			UID:             timeUnix,
			Account_lv:      1,
			Account_exp:     0,
			Weapon_shard:    0,
			Star_stage:      0,
			Choosen_hero_id: 1,
			Weapon_id:       1,
		}

		snet.SDB.Debug().Create(h1)
		snet.SDB.Debug().Create(h2)
		snet.SDB.Debug().Create(h3)
		snet.SDB.Debug().Create(ac)
		request.GetConnection().SendMesg(LAND_REGISTER_ACK, []byte("1"))
		//fmt.Println("success")
	}
}

//Landing 登录业务
type Landing struct {
	snet.BaseRouter
}

//Handle 登录业务处理方法
func (l *Landing) Handle(request isface.IRequest) {
	//fmt.Println("[langding Test]msgID = ,", request.GetMsgId(), "序列化data = ", request.GetData())
	cus := &Data.Customer{}
	//cus1 := &Customer{}
	proto.Unmarshal(request.GetData(), cus)
	//fmt.Println("cus : ", cus)
	//use := &User_info{}
	//use.User_name = cus.User
	//use.User_password = cus.Pass
	//fmt.Println("use = ", use)
	user := Data.User_info{}
	var count int = 0
	snet.SDB.Table("user_info").Where("Account_name=? AND Account_password=?", cus.User, cus.Pass).Scan(&user).Count(&count)
	fmt.Printf("[Login info ] user : ", user)
	//snet.SDB.Where("User_name=? AND User_password=?", cus.User, cus.Pass).Find(&users).Count(&count)
	//fmt.Println("users : \n", users)
	//u := &User_info{User_name: cus.User, User_password: cus.Pass}
	//snet.SDB.Model(u).Count(&count)
	//fmt.Println("count", count)
	if count == 1 {
		fmt.Println("用户", cus.User, "登陆了")
		msgTemp := &Data.PlayerInfo{
			Roomid:   1,
			Uid:      user.UID,
			Username: user.Account_name,
		}
		msgTempbuf, _ := proto.Marshal(msgTemp)
		request.GetConnection().SendMesg(LAND_ACK, msgTempbuf)
		//request.GetConnection().SetConnID(uint32(user.UID))
		request.GetConnection().SetConnID(uint32(user.UID))
		//添加到链接管理模块中去
		/*
			snet.ConnMap[uint32(user.UID)] = &snet.ConnState{
				Conn:   request.GetConnection(),
				Count:  -1,
				State:  snet.NORMAL,
				Roomid: 0,
			}
		*/
		snet.AddConn(uint32(user.UID), request.GetConnection())

	} else {
		fmt.Println("登录失败！！！")
		msgTemp := &Data.PlayerInfo{
			Roomid:   0,
			Uid:      999,
			Username: "error",
		}
		msgTempbuf, _ := proto.Marshal(msgTemp)

		//fmt.Println("2222")
		request.GetConnection().SendMesg(LAND_ACK, msgTempbuf)
	}

}

//Pingtest 测试业务
