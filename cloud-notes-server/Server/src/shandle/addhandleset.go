package shandle

import (
	"isface"
)

//在这里添加Handle集合
const (
	//_ = iota
	LAND               = 1   //登录业务（收）
	LAND_ACK           = 2   //登录业务应答（发）
	LAND_REGISTER      = 3   //注册业务（收）
	LAND_REGISTER_ACK  = 4   //注册业务应答（发）
	CREAT_ROOM         = 5   //创建房间业务（收）
	CREAT_ROOM_ACK     = 6   //创建房间应答（发）
	JOIN_ROOM          = 7   //加入房间业务（收）
	JOIN_ROOM_ACK      = 8   //加入房间业务应答  通知用户是否成功进入房间（发）
	FLASH_ROOM         = 9   //刷新房间业务（收）
	FLASH_ROOM_ACK     = 10  //刷新房间业务应答（发）
	START_GAME         = 11  //房主开始游戏（收）
	START_GAME_ACK     = 12  //开始游戏的回应（发）
	GAME_START_PRO     = 13  //房间预处理的信息，由房主发给我（收）
	GAME_START_ACK     = 14  //房间预处理结束，广播给所有人，房间Id，让他们自己读取资源信息和玩家信息（发）
	UPDATA_PLAYER_PRO  = 15  //玩家更新所有自己的数据到服务器上
	UPDATA_PLAYER_ACK  = 16  //全部人更新完后，把数据下发给所有人，并且需要等待3-4秒
	UPDATA_MYSELF      = 17  //玩家更新自己的信息（收）
	UPDATA_RESOUCE     = 18  //玩家更新地图资源
	UPDATA_RESOUCE_ACK = 180 //玩家更新地图资源应答
	Game_OVERDEAL      = 19  //玩家结束游戏后提交给数据库的信息
	EXIT_ROOM          = 20  //退出房间的方法
	EXIT_ROOM_ACK      = 21  //退出房间后给所有人的消息

	INIT_PLAYERINFO     = 23 //给玩家登陆时候发送他的信息（收）
	INIT_PLAYERINFO_ACK = 24 //给玩家发送信息（发）

	PLAYER_READY          = 25 //玩家发送请求准备游戏
	PLAYER_READY_ACK      = 26 //给所有人广播该玩家切换准备状态。
	UPDATA_ACCOUNTMSG     = 27 //更新账号信息收不知道怎么写
	UPDATA_ACCOUNTMSG_ACK = 28 //回应（发）不知道怎么写
	UPDATA_HEROSKILL      = 29 //更新英雄技能
	UPDATA_HEROSKILL_ACK  = 30 //更新英雄技能回应
	UPDATA_WPAPON         = 31 //更新武器
	UPDATA_WPAPON_ACK     = 32 //更新武器回应

	UPDATA_FEN           = 33 //玩家更改积分情况（收）
	UPDATA_FEN_ACK       = 34 //当前所有玩家排好序的积分。（发）
	UPDATA_NICK_NAME     = 35 //更改玩家昵称
	UPDATA_NICK_NAME_ACK = 36 //更改玩家昵称回复
	UPDATA_PASS          = 37 //更改密码
	UPDATA_PASS_ACK      = 38 //更改密码回复

	JOIN_ROOM_BORAD = 100 //给所有人广播加入的人的信息

	BORAD_ROOM_ALLPLAYER = 101 //给所有人广播所有资源和人的信息

	CHAT_ROOM_BORAD      = 102 //聊天广播（收）
	TIME_ROOM_BORAD      = 104 //告诉房间内玩家时间广播（发）
	GAME_OVER_BORAD      = 105 //告诉房间内所有玩家游戏结束了（发）
	JOIN_ROOM_BORAD_HERO = 106 //给所有人广播加入人的具体信息

)

func AddHandleInit(s isface.IServer) {
	//不是房间方法，添加0，是房间方法添加1
	s.AddHandle(CREAT_ROOM, &CreateRoom{}, "创建房间的方法", 0)
	s.AddHandle(JOIN_ROOM, &JoinRoom{}, "加入房间的方法", 0)
	s.AddHandle(LAND, &Landing{}, "登陆的方法", 0)
	s.AddHandle(LAND_REGISTER, &LandingRegister{}, "注册的方法", 0)
	s.AddHandle(FLASH_ROOM, &FlashRoom{}, "刷新房间的方法", 0)
	s.AddHandle(START_GAME, &StartGame{}, "开始游戏的方法", 1)
	s.AddHandle(GAME_START_PRO, &StartGamePro{}, "开始房间预处理的处理方法", 1)
	s.AddHandle(UPDATA_MYSELF, &PlayerUpdataSelfInfo{}, "玩家刷新自己的信息的方法", 1)
	s.AddHandle(Game_OVERDEAL, &GameStopHandle{}, "游戏结束的方法", 1)
	s.AddHandle(EXIT_ROOM, &ExitRoom{}, "退出房间的方法", 0)
	s.AddHandle(PLAYER_READY, &ReadyGame{}, "准备游戏的方法", 1)
	s.AddHandle(405, &TestRun{}, "测试一下能不能玩的方法", 1)
	s.AddHandle(INIT_PLAYERINFO, &InitUiScene{}, "初始化玩家信息的方法，登陆UI界面时候什么英雄碎片发送", 0)
	s.AddHandle(CHAT_ROOM_BORAD, &Chatting{}, "聊天的方法", 1)
	s.AddHandle(UPDATA_FEN, &UpdataRank{}, "对玩家进行积分排序", 1)
	s.AddHandle(UPDATA_RESOUCE, &UpdateResource{}, "吃掉资源的方法", 1)
	s.AddHandle(UPDATA_NICK_NAME, &UpdataNickName{}, "更改玩家昵称", 0)
	s.AddHandle(UPDATA_WPAPON,&UpdataWpapon{},"更新绑定武器",0)
	s.AddHandle(UPDATA_HEROSKILL,&UpdataHeroSkill{},"更新玩家技能",0)
	s.AddHandle(UPDATA_PASS, &UpdataPass{}, "更改密码", 0)
}
