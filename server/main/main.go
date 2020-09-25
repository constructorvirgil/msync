package main

import (
	"github.com/aceld/zinx/znet"
	"github.com/constructorvirgil/msync/main/filemanage"
	"github.com/constructorvirgil/msync/main/router"
)

func main() {
	//创建服务器句柄
	s := znet.NewServer()

	//配置路由
	s.AddRouter(router.MsgIdLogin, &router.LoginRouter{})
	s.AddRouter(router.MsgIdTransFile, &router.TransFileRouter{})

	//启动文件管理routine
	filemanage.Init()

	//开启服务
	s.Serve()
}