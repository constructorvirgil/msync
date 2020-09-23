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
func (this *LoginRouter) LoginHandle(request ziface.IRequest) {
	msgLogin := common.MsgLogin{}
	data := request.GetData()
	err := msgLogin.UnPack(data)
	if err != nil{
		fmt.Println("unpack failure: ", string(data))
		return
	}

	err = request.GetConnection().SendBuffMsg(0, []byte("login ok"))
	if err != nil {
		fmt.Println(err)
	}
}
