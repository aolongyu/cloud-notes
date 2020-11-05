package isface


/*
负责执行handler中的处理业务的方法
 */

type IRouter interface {
	//处理conn业务的方法

	Handle(request IRequest)
}