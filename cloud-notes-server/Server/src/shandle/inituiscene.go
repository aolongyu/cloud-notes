package shandle

import (
	Data "data"
	"fmt"
	"isface"
	"snet"

	"github.com/golang/protobuf/proto"
)

type InitUiScene struct {
	snet.BaseRouter
}

func (I *InitUiScene) Handle(request isface.IRequest) {
	data := request.GetData()
	PlayerInfoData := &Data.PlayerInfo{
		Roomid:   0,
		Uid:      0,
		Username: "",
	}
	proto.Unmarshal(data, PlayerInfoData)
	fmt.Println("UId------", PlayerInfoData.Uid)

	SendDataUser := &Data.User_info{}
	SendDataAccount := &Data.Account_info{}
	SendDataHero := &Data.Hero{}

	//snet.SDB.Debug().Raw("select a.HeroId,a.HeroLv,a.HeroExp,a.HeroSkill,user_name as UserName,account_lv as AccountLv ,account_exp as AccountExp, choosen_hero_id as ChooseHeroId,weapon_id as WeaponId FROM account_info , user_info ,( select a.hero_id as HeroId , a.hero_lv as HeroLv , hero_exp as HeroExp,hero_skill as HeroSkill from hero as a where UID = ? and hero_id = (SELECT choosen_hero_id from account_info where Uid = ?)  ) as a where account_info.uid = user_info.uid and user_info.uid = ?", PlayerInfoData.Uid, PlayerInfoData.Uid, PlayerInfoData.Uid).Scan(SendData)
	//snet.SDB.Debug().Raw("select a.HeroId,a.HeroLv,a.HeroExp,a.HeroSkill,user_name as UserName,account_lv as AccountLv ,account_exp as AccountExp, choosen_hero_id as ChooseHeroId,weapon_id as WeaponId FROM account_info , user_info ,( select a.hero_id as HeroId , a.hero_lv as HeroLv , hero_exp as HeroExp,hero_skill as HeroSkill from hero as a where UID = ? and hero_id = (SELECT choosen_hero_id from account_info where Uid = ?)  ) as a where account_info.uid = user_info.uid and user_info.uid = ?", PlayerInfoData.Uid, PlayerInfoData.Uid, PlayerInfoData.Uid).FirstOrInit(SendData)
	//snet.SDB.Debug().Raw("select a.HeroId,a.HeroLv,a.HeroExp,a.HeroSkill,user_name as UserName,account_lv as AccountLv ,account_exp as AccountExp, choosen_hero_id as ChooseHeroId,weapon_id as WeaponId FROM account_info , user_info ,( select a.hero_id as HeroId , a.hero_lv as HeroLv , hero_exp as HeroExp,hero_skill as HeroSkill from hero as a where UID = ? and hero_id = (SELECT choosen_hero_id from account_info where Uid = ?)  ) as a where account_info.uid = user_info.uid and user_info.uid = ?", PlayerInfoData.Uid, PlayerInfoData.Uid, PlayerInfoData.Uid).Find(SendData)
	snet.SDB.Debug().Table("user_info").Where("uid=?", PlayerInfoData.Uid).Scan(&SendDataUser)
	snet.SDB.Debug().Table("account_info").Where("uid=?", PlayerInfoData.Uid).Scan(&SendDataAccount)
	snet.SDB.Debug().Table("hero").Where("hero_id=?", SendDataAccount.Choosen_hero_id).Scan(&SendDataHero)
	//fmt.Println("hero_exp:", SendDataHero.Hero_exp, "skill", SendDataHero.Choosen_skill, "weapon", SendDataAccount.Hero_weapon, "accexp", SendDataAccount.Account_exp)
	SendData := &Data.InitPlayerInfo{
		HeroId:       SendDataHero.Hero_id,
		HeroLv:       SendDataHero.Hero_lv,
		HeroExp:      SendDataHero.Hero_exp,     //0
		HeroSkill:    SendDataHero.Hero_skill,   //0
		WeaponId:     SendDataAccount.Weapon_id, //0
		UserName:     SendDataUser.User_name,
		AccountLv:    SendDataAccount.Account_lv,
		AccountExp:   SendDataAccount.Account_exp, //0
		ChooseHeroId: SendDataAccount.Choosen_hero_id,
	}
	//fmt.Println("给玩家发送了初始化登陆界面的信息------------,SendData:", SendData)

	buf, err := proto.Marshal(SendData)
	if err != nil {
		fmt.Println("---------------------------------------")
	}
	fmt.Println("给玩家发送了初始化登陆界面的信息,SendData:", SendData)
	request.GetConnection().SendMesg(INIT_PLAYERINFO_ACK, buf)
}
