package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/constructorvirgil/msync/common"
)

const (
	MsgIdLogin = iota
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
		return
	}

	err = request.GetConnection().SendBuffMsg(0, []byte("login ok"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//创建服务器句柄
	s := znet.NewServer()

	//配置路由
	s.AddRouter(MsgIdLogin, &LoginRouter{})

	//开启服务
	s.Serve()
}