package shandle

import (
	"isface"
)

//在这里添加Handle集合
const (
	LOGIN_HANDLE = 1
	TEST_HANDLE = 404
)

func AddHandleInit(s isface.IServer) {
	//s.AddHandle(TEST_HANDLE,&Nofound{},"测试能否连接",0)
	s.AddHandle("login",&Login{},"登陆",0)
	s.AddHandle("regist",&Register{},"注册",0)
	s.AddHandle("finduser",&FindUser{},"管理员查找用户",0)
	s.AddHandle("closeuser",&CloseUserByid{},"封锁用户，根据ID",0)
	s.AddHandle("deluser",&DeleteUserById{},"根据用户id删除用户",0)
	s.AddHandle("fiusn",&FindUserName{},"模糊查找用户",0)

	s.AddHandle("getNList",&GetNlist{},"获取笔记本列表",0)
	s.AddHandle("CrNoBook",&CreateNoteBook{},"创建笔记本",0)
	s.AddHandle("CrNote",&CreateNote{},"创建笔记",0)
	s.AddHandle("FindNote",&FindNoteByUserName{},"根据id查找笔记",0)
	s.AddHandle("AddToBook",&AddNoteToNoteBook{},"把笔记加入笔记本",0)
	s.AddHandle("ViewNote",&ViewNote{},"根据笔记Id查看笔记",0)
	s.AddHandle("CollnoBook",&CollectNoteBook{},"收藏笔记",0)
	s.AddHandle("FiNoByid",&GetNoteListById{},"根据笔记本ID获得笔记",0)
	s.AddHandle("NoBByName",&FindNoteBookByUName{},"根据用户名查找笔记本",0)
	s.AddHandle("delnote",&DeleteNote{},"根据笔记id删除笔记",0)
	s.AddHandle("editnote",&EditNote{},"修改笔记",0)
	s.AddHandle("edNobook",&EditNoteBook{},"修改笔记本",0)
	s.AddHandle("finbyuid",&FindByidDetail{},"获得详细笔记本得id",0)
	s.AddHandle("getallnot",&GetAllNote{},"获得所有笔记信息",0)
	s.AddHandle("thumnote",&ThumbsUpNote{},"给笔记点赞",0)
	s.AddHandle("thumnobo",&ThumbsUpNoteBook{},"给笔记本点赞",0)
	s.AddHandle("frcover",&FindRountOver{},"找被举报10次的笔记",0)
	s.AddHandle("renote",&ReportNote{},"举报笔记",0)
	s.AddHandle("renotebo",&ReportNoteNoteBook{},"举报笔记本",0)
	s.AddHandle("conoadd",&CollectNoteBookAdd{},"收藏笔记进笔记本",0)
	s.AddHandle("collnrev",&CollectNoteBookRemove{},"移除笔记本",0)
	s.AddHandle("crecbo",&CreateCBook{},"创建收藏笔记本",0)
	s.AddHandle("delcb",&DelCBook{},"删除删除的笔记",0)
	//s.AddHandle("finalcb")
	//s.AddHandle("fnicb")

}
