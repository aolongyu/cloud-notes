package shandle

import (
	Data "data"
	"fmt"
	"isface"
	"snet"

	"github.com/golang/protobuf/proto"
)

type UpdateResource struct {
	snet.BaseRouter
}

func (ur *UpdateResource) Handle(request isface.IRequest) {

	var NowRoom *snet.Room
	var update *Data.PreloadPro         //房主更新资源数据
	var updateDate *Data.MapResourcePro //吃掉资源数据
	nowId := request.GetConnection().GetConnID()
	//修改的数据放在preloadpro中的allmapresource数组中
	update = &Data.PreloadPro{}
	//反序列化
	proto.Unmarshal(request.GetData(), update)
	fmt.Println("收到的消息------", update)
	if update.AllPlayer == nil {
		updateDate = &Data.MapResourcePro{}
		//反序列化
		proto.Unmarshal(request.GetData(), updateDate)
		NowRoom = snet.RoomMgr.GetRoom(int32(updateDate.Roomid))

	} else {
		fmt.Println(update.AllPlayer[0].PlayerMsg.Roomid)
		NowRoom = snet.RoomMgr.GetRoom(int32(update.AllPlayer[0].PlayerMsg.Roomid))
	}
	if NowRoom==nil{
		return
	}
	//先绑定该链接所在的房间
	//NowRoom := snet.RoomMgr.GetRoom(int32(request.GetConnection().GetConnID()))

	//若是房主则房主的链接id==房间id，对地图资源进行替换；
	//先删除原有的资源，再用所传数据更新该房间的资源
	if uint32(NowRoom.Id) == nowId {

		NowRoom.RWMutex.Lock()
		defer NowRoom.RWMutex.Unlock()

		fmt.Println("----------", update.AllPlayer, "-------------")
		if update.AllPlayer == nil {
			fmt.Println("这是房主吃掉资源的更新消息")
			updateDate := &Data.MapResourcePro{}
			//反序列化
			err := proto.Unmarshal(request.GetData(), updateDate)
			//fmt.Println("***************", updateDate, "***************")
			//fmt.Println("&&&&&&&&&&&&&", NowRoom.AllResource, "&&&&&&&&&&&&&&&&&&&")
			if err != nil {
				fmt.Println(nowId, "--updateresource房主数据解析崩了")
			} else {
				//将所碰到的资源信息转化为房间资源信息类型
				coo1 := snet.Coord{
					X: updateDate.MapResourceCoord.X,
					Y: updateDate.MapResourceCoord.Y,
				}
				//在房间资源map中删除该元素
				if _, ok := NowRoom.AllResource[coo1]; ok {
					//存在
					delete(NowRoom.AllResource, coo1)
					request.GetConnection().SendMesg(UPDATA_RESOUCE_ACK, []byte("1"))
				} else {
					request.GetConnection().SendMesg(UPDATA_RESOUCE_ACK, []byte("0"))
				}
			}
			snet.ConnMap[nowId].Count += 1
			return
		}
		//获得新资源个数
		fmt.Println("这是房主的更新地图资源消息")
		lens := len(update.AllMapResource)
		//开辟新的map用于替换旧的房间资源map
		newallresource := make(map[snet.Coord]snet.MapResource, lens)
		//遍历所有的新资源
		for _, v := range update.AllMapResource {
			tmp1 := snet.Coord{
				X: v.MapResourceCoord.X,
				Y: v.MapResourceCoord.Y,
			}
			tmp2 := snet.MapResource{
				Id:    v.Id,
				Coord: tmp1,
			}
			//将每个新资源以服务端的形势存到新map中去
			newallresource[tmp1] = tmp2
		}
		//用已填充的新map替换旧map
		NowRoom.AllResource = newallresource
		snet.ConnMap[nowId].Count += 1
		fmt.Println("房主更新资源成功！")
	} else {
		//若不是房主，则只对所碰到的资源进行得到判定
		//只发送所碰到的资源信息
		NowRoom.RWMutex.Lock()
		defer NowRoom.RWMutex.Unlock()
		//update := &Data.MapResourcePro{}
		//反序列化
		//proto.Unmarshal(request.GetData(), update)
		//将所碰到的资源信息转化为房间资源信息类型
		coo := snet.Coord{
			X: updateDate.MapResourceCoord.X,
			Y: updateDate.MapResourceCoord.Y,
		}
		//在房间资源map中删除该元素
		if _, ok := NowRoom.AllResource[coo]; ok {
			//存在
			delete(NowRoom.AllResource, coo)
			request.GetConnection().SendMesg(UPDATA_RESOUCE_ACK, []byte("1"))
		} else {
			request.GetConnection().SendMesg(UPDATA_RESOUCE_ACK, []byte("0"))
		}
		snet.ConnMap[nowId].Count += 1
		fmt.Println("这是非房主的玩家吃掉资源的更新消息")
	}

}
