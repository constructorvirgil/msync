package router

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/constructorvirgil/msync/common"
)

//Login 路由
type LoginRouter struct {
	znet.BaseRouter
}

//Login 回调
func (this *LoginRouter) Handle(request ziface.IRequest) {
	msgLogin := common.MsgLogin{}
	data := request.GetData()
	err := msgLogin.UnPack(data)
	if err != nil{
		fmt.Println("unpack failure: ", string(data))
		_ = Response(request, MsgIdLogin, RespCodeDecodeError, "decode message error")
		return
	}

	_ = Response(request, MsgIdLogin, RespCodeOK, "login ok")
}
