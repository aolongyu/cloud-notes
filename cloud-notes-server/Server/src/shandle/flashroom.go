package shandle

import (
	Data "data"
	"fmt"
	"isface"
	"snet"

	"google.golang.org/protobuf/proto"
)

/*
import (
	Data "data"
	"fmt"
	"github.com/golang/protobuf/proto"
	"isface"
	"snet"
)
*/
type FlashRoom struct {
	snet.BaseRouter
}

//刷新房间的hand
func (F *FlashRoom) Handle(request isface.IRequest) {
	//allroomnow := &Data.Rooms{}
	//proto.Unmarshal(request.GetDat(),allroomnow)
	//RoomClient := request.GetConnection().GetRoomManager()

	allroom := snet.RoomMgr.GetAllRoom()

	rooms := &Data.Rooms{
		Readys:  0,
		Battles: 0,
		Overs:   0,
		Room:    nil,
	}
	readyCount := 0
	battleCount := 0
	overCount := 0
	for _, roomOne := range allroom {
		//if allroomnow.Room
		roomPro := &Data.RoomPro{
			Uid:    roomOne.GetRoomID(),
			Status: roomOne.GetRoomStatus(),
			Nums:   roomOne.GetLen(),
		}
		if roomPro.Status == snet.GAME_STATUS_READY {
			readyCount++
			rooms.Room = append(rooms.Room, roomPro)
		} else if roomPro.Status == snet.GAME_STATUS_BATTLE {
			battleCount++
		} else if roomPro.Status == snet.GAME_STATUS_OVER {
			overCount++
		} else {
			fmt.Println("flashroom err!")
		}
	}
	rooms.Overs = int32(overCount)
	rooms.Readys = int32(readyCount)
	rooms.Battles = int32(battleCount)
	if rooms.Readys == 0{
		rooms.Battles = -1
	}
	data, _ := proto.Marshal(rooms)
	request.GetConnection().SendMesg(FLASH_ROOM_ACK, data)
	fmt.Println("收到刷新的请求,并且也发送回去了")
	fmt.Println("收到--", request.GetConnection().GetConnID(), "--刷新的请求,并且也发送回去了10--Rooms")
	/*
		rooms := &Data.Rooms
			Data.Room{

			}
		}
		Playinfo := &Data.PlayerInfo{}
		proto.Unmarshal(request.GetData(), Playinfo)
	*/
}
